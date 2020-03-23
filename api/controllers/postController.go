package controllers

import (
	"database/sql"
	"encoding/json"
	"math/rand"
	"net/http"
	"pogo/api/db"
	h "pogo/api/helpers"
	m "pogo/api/models"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

// DB initiali<e
var mainDB = db.InitDB()

// All posts
func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "aplication/json")

	// Query
	rows, err := mainDB.Query("SELECT * FROM posts AS p INNER JOIN (SELECT id, username, email FROM users) AS u ON p.user_id = u.id ORDER BY p.created_at DESC")

	h.CheckErr(err)

	var posts m.Posts

	for rows.Next() {
		var post m.Post
		var author m.Author
		var user_id string
		var post_id string
		var comment_post_id string
		var comments m.Comments

		err = rows.Scan(&post.ID, &post.Title, &post.Body, &post.CreatedAt, &post.UpdatedAt, &user_id, &post.Likes, &author.ID, &author.UserName, &author.Email)

		h.CheckErr(err)

		post.Author = author

		stmt, err := mainDB.Prepare("SELECT * FROM comments AS c INNER JOIN ( SELECT id, username, email FROM users) u ON c.user_id = u.id WHERE c.post_id = ?")

		var commentsRow *sql.Rows
		commentsRow, err = stmt.Query(post.ID)

		for commentsRow.Next() {
			var comment m.Comment
			var commentAuthor m.Author

			err = commentsRow.Scan(&comment.ID, &comment.Comment, &post_id, &comment_post_id, &commentAuthor.ID, &commentAuthor.UserName, &commentAuthor.Email)

			h.CheckErr(err)
			comment.CommentAuthor = commentAuthor
			comments = append(comments, comment)
		}
		post.Comments = comments
		posts = append(posts, post)
	}
	json.NewEncoder(w).Encode(posts)
}

// Only user's posts
func GetOwnPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "aplication/json")

	// Get post id from paramaters of request
	params := mux.Vars(r)

	// Query prepare
	stmt, err := mainDB.Prepare("SELECT posts.id, posts.title, posts.body, posts.created_at, posts.updated_at, posts.user_id, posts.likes, users.id FROM posts INNER JOIN users ON users.id = posts.user_id WHERE users.username = ? ORDER BY posts.created_at DESC")
	h.CheckErr(err)

	// Query execute
	rows, errQuery := stmt.Query(params["username"])
	h.CheckErr(errQuery)

	var ownPosts m.OwnPosts

	for rows.Next() {
		var post m.OwnPost
		var post_user_id string
		var comment_post_id string
		var post_id string
		var user_id string
		var comments m.Comments

		err = rows.Scan(&post.ID, &post.Title, &post.Body, &post.CreatedAt, &post.UpdatedAt, &post_user_id, &post.Likes, &user_id)
		h.CheckErr(err)

		// Query prepare for comments
		stmt, err := mainDB.Prepare("SELECT * FROM comments AS c INNER JOIN ( SELECT id, username, email FROM users) u ON c.user_id = u.id WHERE c.post_id = ?")
		h.CheckErr(err)

		var commentsRow *sql.Rows

		// Query execute for comments
		commentsRow, err = stmt.Query(post.ID)
		h.CheckErr(err)

		for commentsRow.Next() {
			var comment m.Comment
			var commentAuthor m.Author

			err = commentsRow.Scan(&comment.ID, &comment.Comment, &post_id, &comment_post_id, &commentAuthor.ID, &commentAuthor.UserName, &commentAuthor.Email)
			h.CheckErr(err)

			comment.CommentAuthor = commentAuthor
			comments = append(comments, comment)
		}

		post.Comments = comments
		ownPosts = append(ownPosts, post)
	}

	json.NewEncoder(w).Encode(ownPosts)
}

// Get post by id
func GetPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "aplication/json")

	// Get post id from paramaters of request
	params := mux.Vars(r)

	// Query prepare for posts
	stmt, err := mainDB.Prepare("SELECT * FROM posts AS p INNER JOIN (SELECT id, username, email from users) AS u ON p.user_id = u.id WHERE p.id = ?")
	h.CheckErr(err)

	// Query execute for posts
	rows, errQuery := stmt.Query(params["id"])
	h.CheckErr(errQuery)

	var post m.Post
	var author m.Author

	for rows.Next() {
		var user_id string
		var post_id string
		var comments m.Comments
		var comment_post_id string

		err = rows.Scan(&post.ID, &post.Title, &post.Body, &post.CreatedAt, &post.UpdatedAt, &user_id, &post.Likes, &author.ID, &author.UserName, &author.Email)
		h.CheckErr(err)

		post.Author = author

		var stmtNew *sql.Stmt

		// Query prepare for comments
		stmtNew, err := mainDB.Prepare("SELECT * FROM comments AS c INNER JOIN ( SELECT id, username, email FROM users) u ON c.user_id = u.id WHERE c.post_id = ?")

		var commentsRow *sql.Rows

		// Query execute for comments
		commentsRow, err = stmtNew.Query(post.ID)

		for commentsRow.Next() {
			var comment m.Comment
			var commentAuthor m.Author

			err = commentsRow.Scan(&comment.ID, &comment.Comment, &post_id, &comment_post_id, &commentAuthor.ID, &commentAuthor.UserName, &commentAuthor.Email)
			h.CheckErr(err)

			comment.CommentAuthor = commentAuthor
			comments = append(comments, comment)
			post.Comments = comments
		}
	}

	// Checking if post doesn't exist
	isID, _ := strconv.ParseInt(post.ID, 10, 0)

	if isID == 0 {
		json.NewEncoder(w).Encode(h.Error("Post not found"))
		return
	} else {
		json.NewEncoder(w).Encode(post)
	}
}

// Create Post
func CreatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "aplication/json")

	var post m.SendPost

	// Get post content from body of request
	_ = json.NewDecoder(r.Body).Decode(&post)

	// Generate random id for new post
	post.ID = strconv.Itoa(rand.Intn(100000000000))

	// Query prepare for new post
	stmt, err := mainDB.Prepare("INSERT INTO posts(id, title, body, user_id) VALUES (?, ?, ?, ?)")
	h.CheckErr(err)

	// Query execute for new post
	_, errExec := stmt.Exec(post.ID, post.Title, post.Body, post.UserID)
	h.CheckErr(errExec)

	json.NewEncoder(w).Encode(post)
}

// Update Post
func UpdatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "aplication/json")

	// Get post id from paramaters of request
	params := mux.Vars(r)

	var updatedPost m.UpdatedPost

	// Get post content from body of request
	_ = json.NewDecoder(r.Body).Decode(&updatedPost)

	// Query prepare for updated post
	stmt, err := mainDB.Prepare("UPDATE posts SET id = ?, title = ?, body = ?, created_at = ?, updated_at = ?, user_id = ?, likes = ? WHERE id = ?")
	h.CheckErr(err)

	// Query execute for updated post
	result, errExec := stmt.Exec(updatedPost.ID, updatedPost.Title, updatedPost.Body, updatedPost.CreatedAt, updatedPost.UpdatedAt, updatedPost.UserID, updatedPost.Likes, params["id"])
	h.CheckErr(errExec)

	rowAffected, errLast := result.RowsAffected()
	h.CheckErr(errLast)

	// Checking if post doesn't exist
	if rowAffected == 0 {
		json.NewEncoder(w).Encode(h.Error("Post not found"))
		return
	} else {
		json.NewEncoder(w).Encode(updatedPost)
	}
}

// Delete Post
func DeletePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "aplication/json")

	// Get post id from paramaters of request
	params := mux.Vars(r)

	// Query prepare for deleted post
	stmt, err := mainDB.Prepare("DELETE FROM posts WHERE id = ?")
	h.CheckErr(err)

	// Query execute for deleted post
	result, errExec := stmt.Exec(params["id"])
	h.CheckErr(errExec)

	rows, errRow := result.RowsAffected()
	h.CheckErr(errRow)

	// Checking if post doesn't exist
	if rows == 0 {
		json.NewEncoder(w).Encode(h.Error("Post not found"))
		return
	} else {
		json.NewEncoder(w).Encode("Post Deleted")
	}
}

// Get All Comments
func GetAllComments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "aplication/json")

	var comments m.GetComments

	// Query get all comments
	rows, err := mainDB.Query("SELECT * FROM comments AS c INNER JOIN ( SELECT id, username, email FROM users) u ON c.user_id = u.id")
	h.CheckErr(err)

	for rows.Next() {
		var comment m.GetComment
		var author m.Author

		err = rows.Scan(&comment.ID, &comment.Comment, &comment.PostID, &comment.UserID, &author.ID, &author.UserName, &author.Email)
		h.CheckErr(err)

		comment.CommentAuthor = author
		comments = append(comments, comment)
	}

	json.NewEncoder(w).Encode(comments)
}

// Added Comment
func AddComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "aplication/json")

	var addComment m.AddComment

	// Get post content from body of request
	_ = json.NewDecoder(r.Body).Decode(&addComment)

	// Query prepare for new comment
	stmt, err := mainDB.Prepare("INSERT INTO comments( comment, post_id, user_id) VALUES(?, ?, ?)")
	h.CheckErr(err)

	// Query execute for new comment
	result, errExec := stmt.Exec(&addComment.Comment, &addComment.PostID, &addComment.UserID)
	h.CheckErr(errExec)

	rows, errRow := result.RowsAffected()
	h.CheckErr(errRow)

	// Checking if post doesn't exist
	if rows == 0 {
		json.NewEncoder(w).Encode(h.Error("Something went wrong"))
		return
	} else {
		json.NewEncoder(w).Encode("Comment Added")
	}
}
