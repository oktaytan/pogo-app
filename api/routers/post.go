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
	router.Handle("/api/posts/liked", negroni.New(negroni.WrapFunc(c.MostLikedPosts))).Methods("GET")
	// Get post by id
	router.Handle("/api/posts/{id}", negroni.New(negroni.HandlerFunc(v.VerifyToken.HandlerWithNext), negroni.WrapFunc(c.GetPost))).Methods("GET")
	// Create Post
	router.Handle("/api/posts", negroni.New(negroni.HandlerFunc(v.VerifyToken.HandlerWithNext), negroni.WrapFunc(c.CreatePost))).Methods("POST")
	// Search Post
	router.Handle("/api/posts/search", negroni.New(negroni.HandlerFunc(v.VerifyToken.HandlerWithNext), negroni.WrapFunc(c.SearchPosts))).Methods("POST")
	// Like Post
	router.Handle("/api/posts/{id}", negroni.New(negroni.HandlerFunc(v.VerifyToken.HandlerWithNext), negroni.WrapFunc(c.UpdatePost))).Methods("PUT")
	// Get likes by user id
	router.Handle("/api/likes/{id}", negroni.New(negroni.HandlerFunc(v.VerifyToken.HandlerWithNext), negroni.WrapFunc(c.GetLikes))).Methods("GET")
	// User Like
	router.Handle("/api/likes", negroni.New(negroni.HandlerFunc(v.VerifyToken.HandlerWithNext), negroni.WrapFunc(c.LikedByUser))).Methods("POST")
	// Delete Post
	router.Handle("/api/posts/{id}", negroni.New(negroni.HandlerFunc(v.VerifyToken.HandlerWithNext), negroni.WrapFunc(c.DeletePost))).Methods("DELETE")
	// Get All Comments
	router.Handle("/api/comments", negroni.New(negroni.HandlerFunc(v.VerifyToken.HandlerWithNext), negroni.WrapFunc(c.GetAllComments))).Methods("GET")
	// Add Comment
	router.Handle("/api/comments", negroni.New(negroni.HandlerFunc(v.VerifyToken.HandlerWithNext), negroni.WrapFunc(c.AddComment))).Methods("POST")
	// Delete Comment
	router.Handle("/api/comments/{id}", negroni.New(negroni.HandlerFunc(v.VerifyToken.HandlerWithNext), negroni.WrapFunc(c.DeleteComment))).Methods("DELETE")

	return router
}
