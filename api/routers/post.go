package routers

import (
	c "pogo/api/controllers"
	v "pogo/api/middleware"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func SetPostRoutes(router *mux.Router) *mux.Router {

	// All posts
	router.Handle("/api/posts", negroni.New(negroni.HandlerFunc(v.VerifyToken.HandlerWithNext), negroni.WrapFunc(c.GetAllPosts))).Methods("GET")
	// Only user's posts
	router.Handle("/api/{username}/posts", negroni.New(negroni.HandlerFunc(v.VerifyToken.HandlerWithNext), negroni.WrapFunc(c.GetOwnPosts))).Methods("GET")
	// Get post by id
	router.Handle("/api/posts/{id}", negroni.New(negroni.HandlerFunc(v.VerifyToken.HandlerWithNext), negroni.WrapFunc(c.GetPost))).Methods("GET")
	// Create Post
	router.Handle("/api/posts", negroni.New(negroni.HandlerFunc(v.VerifyToken.HandlerWithNext), negroni.WrapFunc(c.CreatePost))).Methods("POST")
	// Update Post
	router.Handle("/api/posts/{id}", negroni.New(negroni.HandlerFunc(v.VerifyToken.HandlerWithNext), negroni.WrapFunc(c.UpdatePost))).Methods("PUT")
	// Delete Post
	router.Handle("/api/posts/{id}", negroni.New(negroni.HandlerFunc(v.VerifyToken.HandlerWithNext), negroni.WrapFunc(c.DeletePost))).Methods("DELETE")

	return router
}
