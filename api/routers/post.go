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
	// Get most liked posts
	router.Handle("/api/posts/liked", negroni.New(negroni.HandlerFunc(v.VerifyToken.HandlerWithNext), negroni.WrapFunc(c.MostLikedPosts))).Methods("GET")
	// Get post by id
	router.Handle("/api/posts/{id}", negroni.New(negroni.HandlerFunc(v.VerifyToken.HandlerWithNext), negroni.WrapFunc(c.GetPost))).Methods("GET")
	// Create Post
	router.Handle("/api/posts", negroni.New(negroni.HandlerFunc(v.VerifyToken.HandlerWithNext), negroni.WrapFunc(c.CreatePost))).Methods("POST")
	// Search Post
	router.Handle("/api/posts/search", negroni.New(negroni.HandlerFunc(v.VerifyToken.HandlerWithNext), negroni.WrapFunc(c.SearchPosts))).Methods("POST")
	// Update Post
	router.Handle("/api/posts/{id}", negroni.New(negroni.HandlerFunc(v.VerifyToken.HandlerWithNext), negroni.WrapFunc(c.UpdatePost))).Methods("PUT")
	// Delete Post
	router.Handle("/api/posts/{id}", negroni.New(negroni.HandlerFunc(v.VerifyToken.HandlerWithNext), negroni.WrapFunc(c.DeletePost))).Methods("DELETE")
	// Get All Comments
	router.Handle("/api/comments", negroni.New(negroni.HandlerFunc(v.VerifyToken.HandlerWithNext), negroni.WrapFunc(c.GetAllComments))).Methods("GET")
	// Add Comment
	router.Handle("/api/comments", negroni.New(negroni.HandlerFunc(v.VerifyToken.HandlerWithNext), negroni.WrapFunc(c.AddComment))).Methods("POST")
	// Add Comment
	router.Handle("/api/comments/{id}", negroni.New(negroni.HandlerFunc(v.VerifyToken.HandlerWithNext), negroni.WrapFunc(c.DeleteComment))).Methods("DELETE")

	return router
}
