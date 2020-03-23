package controllers

import (
	"encoding/json"
	"net/http"
	h "pogo/api/helpers"
	m "pogo/api/models"

	"golang.org/x/crypto/bcrypt"

	"github.com/dgrijalva/jwt-go"
	_ "github.com/mattn/go-sqlite3"
)

// All users
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "aplication/json")

	// Query execute for all users
	rows, err := mainDB.Query("SELECT * FROM users")
	h.CheckErr(err)

	var users m.Users

	for rows.Next() {
		var user m.User
		var password string

		err = rows.Scan(&user.ID, &user.UserName, &user.Email, &password)
		h.CheckErr(err)

		users = append(users, user)
	}

	json.NewEncoder(w).Encode(users)
}

// User Regsiter
func Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "aplication/json")

	var registerUser m.RegisterUser

	// Get new user content from body of request
	_ = json.NewDecoder(r.Body).Decode(&registerUser)

	// Check to username exists
	stmtCheckUsername, errCheckUsername := mainDB.Prepare("SELECT COUNT(username) FROM users AS u WHERE u.username = ?")
	h.CheckErr(errCheckUsername)

	rowsCountUsername, errQueryUsername := stmtCheckUsername.Query(registerUser.UserName)
	h.CheckErr(errQueryUsername)

	var countUsername int

	for rowsCountUsername.Next() {
		errCountUsername := rowsCountUsername.Scan(&countUsername)
		h.CheckErr(errCountUsername)
	}

	// Check to email exists
	stmtCheckEmail, errCheckEmail := mainDB.Prepare("SELECT COUNT(email) FROM users AS u WHERE u.email = ?")
	h.CheckErr(errCheckEmail)

	rowsCountEmail, errQueryEmail := stmtCheckEmail.Query(registerUser.Email)
	h.CheckErr(errQueryEmail)

	var countEmail int

	for rowsCountEmail.Next() {
		errCountEmail := rowsCountEmail.Scan(&countEmail)
		h.CheckErr(errCountEmail)
	}

	if countUsername > 0 {
		json.NewEncoder(w).Encode(h.Error("Username already registered"))
	} else if countEmail > 0 {
		json.NewEncoder(w).Encode(h.Error("Email already registered"))
	} else {
		// Checking if passwords don't match
		if registerUser.Password != registerUser.Password2 {
			var error m.Error
			error.Error = true
			error.Message = "Passwords do not match"
			json.NewEncoder(w).Encode(error)
		} else {

			// Hashed password of user
			hashPassword, errPass := bcrypt.GenerateFromPassword([]byte(registerUser.Password), bcrypt.MinCost)
			h.CheckErr(errPass)

			// Changed password which send to database
			registerUser.Password = string(hashPassword)

			// Query prepare for register user
			stmt, err := mainDB.Prepare("INSERT INTO users( username, email, password ) VALUES (?, ?, ?)")
			h.CheckErr(err)

			// Query execute for register user
			result, errExec := stmt.Exec(&registerUser.UserName, &registerUser.Email, &registerUser.Password)
			h.CheckErr(errExec)

			rowAffected, errLast := result.RowsAffected()
			h.CheckErr(errLast)

			if rowAffected == 0 {
				json.NewEncoder(w).Encode(h.Error("Something went wrong"))
				return
			} else {
				var user m.User
				var registerOk m.RegisterOK

				// Query prepare for added user
				stmt, err := mainDB.Prepare("SELECT * FROM users AS u WHERE u.username = ?")
				h.CheckErr(err)

				// Query execute for added user
				rows, errExec := stmt.Query(registerUser.UserName)
				h.CheckErr(errExec)

				for rows.Next() {
					var password string

					err = rows.Scan(&user.ID, &user.UserName, &user.Email, &password)
					h.CheckErr(err)
				}

				registerOk.Error = false
				registerOk.UserAdded = true
				registerOk.User = user

				json.NewEncoder(w).Encode(registerOk)
			}
		}
	}
}

// Login
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "aplication/json")

	var loginUser m.LoginUser

	// Get user content from body of request
	_ = json.NewDecoder(r.Body).Decode(&loginUser)

	// Query prepare for login user
	stmt, err := mainDB.Prepare("SELECT *, count(id) FROM users AS u WHERE u.username = ?")
	h.CheckErr(err)

	// Query execute for login user
	rows, errQuery := stmt.Query(loginUser.UserName)
	h.CheckErr(errQuery)

	for rows.Next() {
		var user m.User
		var password string
		var count int

		err = rows.Scan(&user.ID, &user.UserName, &user.Email, &password, &count)

		// Checking user existing
		if count == 0 {
			var error m.Error
			error.Error = true
			error.Message = "User not found"
			json.NewEncoder(w).Encode(error)
		} else {

			errPass := bcrypt.CompareHashAndPassword([]byte(password), []byte(loginUser.Password))

			if errPass != nil {
				var error m.Error
				error.Error = true
				error.Message = "Password incorrect"
				json.NewEncoder(w).Encode(error)
			} else {
				mySigningKey := []byte("secretKey")

				type MyCustomClaims struct {
					Username string `json:"username"`
					Password string `json:"password"`
					jwt.StandardClaims
				}

				// Create the Claims
				claims := MyCustomClaims{
					loginUser.UserName,
					loginUser.Password,
					jwt.StandardClaims{},
				}

				token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
				ss, _ := token.SignedString(mySigningKey)

				var authUser m.AuthUser
				authUser.IsLogin = true
				authUser.Token = ss

				json.NewEncoder(w).Encode(authUser)
			}
		}
	}

}

// Logout
func Logout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "aplication/json")

	var authUser m.AuthUser
	authUser.IsLogin = false
	authUser.Token = ""

	json.NewEncoder(w).Encode(authUser)
}
