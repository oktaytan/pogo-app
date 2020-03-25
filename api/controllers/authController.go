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

// Tüm kullanıcılar
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "aplication/json")

	// Tüm kullanıcılar için çalışacak sorgu
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

// Kullanıcı kaydı
func Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "aplication/json")

	var registerUser m.RegisterUser

	// Yeni kullanıcı bilgileri request gövedesinden alınıyor
	_ = json.NewDecoder(r.Body).Decode(&registerUser)

	// Kullanıcı adı veritabanında var mı diye kontrol ediliyor
	stmtCheckUsername, errCheckUsername := mainDB.Prepare("SELECT COUNT(username) FROM users AS u WHERE u.username = ?")
	h.CheckErr(errCheckUsername)

	rowsCountUsername, errQueryUsername := stmtCheckUsername.Query(registerUser.UserName)
	h.CheckErr(errQueryUsername)

	var countUsername int

	for rowsCountUsername.Next() {
		errCountUsername := rowsCountUsername.Scan(&countUsername)
		h.CheckErr(errCountUsername)
	}

	// Email veritabanında var mı diye kontrol ediliyor
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
		json.NewEncoder(w).Encode(h.Error("Bu kullanıcı adı sistemde kayıtlı!"))
	} else if countEmail > 0 {
		json.NewEncoder(w).Encode(h.Error("Bu email adresi sistemde kayıtlı!"))
	} else {

		// Şifrelerin eşleşip eşleşmediği kontrol ediliyor
		if registerUser.Password != registerUser.Password2 {
			var error m.Error
			error.Error = true
			error.Message = "Şifreler eşleşmiyor!"
			json.NewEncoder(w).Encode(error)
		} else {

			// Şifre hash formatına getiriliyor
			hashPassword, errPass := bcrypt.GenerateFromPassword([]byte(registerUser.Password), bcrypt.MinCost)
			h.CheckErr(errPass)

			// Veriatabanına kaydediecek şifre hash formatındaki şifreyle değiştiriliyor
			registerUser.Password = string(hashPassword)

			// Yeni kaydedilecek kullanıcı için sorgu oluşturuluyor
			stmt, err := mainDB.Prepare("INSERT INTO users( username, email, password ) VALUES (?, ?, ?)")
			h.CheckErr(err)

			// Yeni kaydedilecek kullanıcı için sorgu çalıştırılıyor
			result, errExec := stmt.Exec(&registerUser.UserName, &registerUser.Email, &registerUser.Password)
			h.CheckErr(errExec)

			rowAffected, errLast := result.RowsAffected()
			h.CheckErr(errLast)

			// Hata oluşması durumu kontrol ediliyor
			if rowAffected == 0 {
				json.NewEncoder(w).Encode(h.Error("Bi terslik oldu.Tekrar deneyiniz."))
				return
			} else {
				var user m.User
				var registerOk m.RegisterOK

				// Yeni eklenen kullanıcıyı veritabanından alacak sorgu hazırlanıyor
				stmt, err := mainDB.Prepare("SELECT * FROM users AS u WHERE u.username = ?")
				h.CheckErr(err)

				// Yeni eklenen kullanıcıyı veritabanından alacak sorgu çalıştırılıyor
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

	// Login olacak kullanıcı bilgileri request gövdesinden alınıyor
	_ = json.NewDecoder(r.Body).Decode(&loginUser)

	// Kullanıcı bilgileri veritabanından alacak sorgu hazırlanıyor
	stmt, err := mainDB.Prepare("SELECT *, count(id) FROM users AS u WHERE u.username = ?")
	h.CheckErr(err)

	// Kullanıcı bilgileri veritabanından alacak sorgu çalıştırılıyor
	rows, errQuery := stmt.Query(loginUser.UserName)
	h.CheckErr(errQuery)

	for rows.Next() {
		var user m.User
		var password string
		var count int

		err = rows.Scan(&user.ID, &user.UserName, &user.Email, &password, &count)

		// Kullanıcının veritabanında kayıtlı olup olmadığı kontrol ediliyor
		if count == 0 {
			json.NewEncoder(w).Encode(h.Error("Bu kullanıcı adı sistemde kayıtlı değil!"))
		} else {

			errPass := bcrypt.CompareHashAndPassword([]byte(password), []byte(loginUser.Password))

			// Kullancının şifresi kontrol ediliyor
			if errPass != nil {
				json.NewEncoder(w).Encode(h.Error("Şifre hatalı!"))
			} else {
				mySigningKey := []byte("secretKey")

				type MyCustomClaims struct {
					Username string `json:"username"`
					Password string `json:"password"`
					jwt.StandardClaims
				}

				// JWT claims hazırlığı
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
