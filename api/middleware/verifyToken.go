package middleware

import (
	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
)

var VerifyToken = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		mySigningKey := []byte("secretKey")
		return mySigningKey, nil
	},

	SigningMethod: jwt.SigningMethodHS256,
})
