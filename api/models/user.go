package models

// User Struct
type User struct {
	ID       string `json:"id"`
	UserName string `json:"username"`
	Email    string `json:"email"`
}

type Users []User

// Register User Struct
type RegisterUser struct {
	UserName  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Password2 string `json:"password2"`
}

type RegisterOK struct {
	Error     bool `json:"error"`
	UserAdded bool `json:"user_added"`
	User      User `json:"user"`
}

// Login User Struct
type LoginUser struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

// Auth User Struct
type AuthUser struct {
	IsLogin bool   `json:"is_login"`
	Token   string `json:"token"`
}
