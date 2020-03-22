package routers

import (
	c "pogo/api/controllers"
	v "pogo/api/middleware"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func SetAuthenticationRoutes(router *mux.Router) *mux.Router {

	// All users
	router.Handle("/api/users", negroni.New(negroni.HandlerFunc(v.VerifyToken.HandlerWithNext), negroni.WrapFunc(c.GetAllUsers))).Methods("GET")
	// User login
	router.Handle("/api/register", negroni.New(negroni.WrapFunc(c.Register))).Methods("POST")
	// User login
	router.Handle("/api/login", negroni.New(negroni.WrapFunc(c.Login))).Methods("POST")
	// User logout
	router.Handle("/api/logout", negroni.New(negroni.HandlerFunc(v.VerifyToken.HandlerWithNext), negroni.WrapFunc(c.Logout))).Methods("GET")

	return router
}
