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

// Veritabanı bağlantısı yapılıyor
var mainDB = db.InitDB()

// Tüm gönderiler
func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "aplication/json")

	// Tüm gönderiler için sql sorgusu çalıştırılıyor
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

		// Her gönderiye ait yorumları getirecek sorgu hazırlanıyor
		stmt, err := mainDB.Prepare("SELECT * FROM comments AS c INNER JOIN ( SELECT id, username, email FROM users) u ON c.user_id = u.id WHERE c.post_id = ?")

		var commentsRow *sql.Rows

		// Her gönderiye ait yorumları getirecek sorgu çalıştırılıyor
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

// Kullanıcının kendi göderileri
func GetOwnPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "aplication/json")

	// Kullanıcı id' si istek parametresinden alınıyor
	params := mux.Vars(r)

	// Kullanıcının kendi gönderilerini getirecek sorgu hazırlanıyor
	stmt, err := mainDB.Prepare("SELECT posts.id, posts.title, posts.body, posts.created_at, posts.updated_at, posts.user_id, posts.likes, users.id FROM posts INNER JOIN users ON users.id = posts.user_id WHERE users.username = ? ORDER BY posts.created_at DESC")
	h.CheckErr(err)

	// Kullanıcının kendi gönderilerini getirecek sorgu çalıştırılıyor
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

		// Kullanıcının gönderilerine yapılmış yorumları getirecek sorgu hazırlanıyor
		stmt, err := mainDB.Prepare("SELECT * FROM comments AS c INNER JOIN ( SELECT id, username, email FROM users) u ON c.user_id = u.id WHERE c.post_id = ?")
		h.CheckErr(err)

		var commentsRow *sql.Rows

		// Kullanıcının gönderilerine yapılmış yorumları getirecek sorgu çalıştırılıyor
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

// Seçili ID' ye göre gönderi getirme
func GetPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "aplication/json")

	// Getirilecek gönderinin id' si istek parametresinden alınıyor
	params := mux.Vars(r)

	// Id' si belirli olan göderiyi getirecek sorgu hazırlanıyor
	stmt, err := mainDB.Prepare("SELECT * FROM posts AS p INNER JOIN (SELECT id, username, email from users) AS u ON p.user_id = u.id WHERE p.id = ?")
	h.CheckErr(err)

	// Id' si belirli olan göderiyi getirecek sorgu çalıştırlıyor
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

		// Gönderiye ait yorumları getirecek sorgu hazırlanıyor
		stmtNew, err := mainDB.Prepare("SELECT * FROM comments AS c INNER JOIN ( SELECT id, username, email FROM users) u ON c.user_id = u.id WHERE c.post_id = ?")

		var commentsRow *sql.Rows

		// Gönderiye ait yorumları getirecek sorgu çalıştırılıyor
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

	// Gönderinin veritabanında olup olmadığı kontrol ediliyor
	isID, _ := strconv.ParseInt(post.ID, 10, 0)

	if isID == 0 {
		json.NewEncoder(w).Encode(h.Error("Gönderi bulunamadı!"))
		return
	} else {
		json.NewEncoder(w).Encode(post)
	}
}

// Gönderi oluşturma
func CreatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "aplication/json")

	var post m.SendPost

	// Gönderini içeriği istek gövdesinden alınıyor
	_ = json.NewDecoder(r.Body).Decode(&post)

	// Yeni gönderi için id oluşturuluyor
	post.ID = strconv.Itoa(rand.Intn(100000000000))

	// Yeni gönderiyi veritabanına ekleyecek sorgu hazırlanıyor
	stmt, err := mainDB.Prepare("INSERT INTO posts(id, title, body, user_id) VALUES (?, ?, ?, ?)")
	h.CheckErr(err)

	// Yeni gönderiyi veritabanına ekleyecek sorgu çalıştırılıyor
	_, errExec := stmt.Exec(post.ID, post.Title, post.Body, post.UserID)
	h.CheckErr(errExec)

	json.NewEncoder(w).Encode(post)
}

// Gönderi güncelleme
func UpdatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "aplication/json")

	// Güncellenecek gönderinin id' si istek parametresinden alınıyor
	params := mux.Vars(r)

	var updatedPost m.UpdatedPost

	// Güncellenecek gönderinin yeni içeriği istek gövdesinden alınıyor
	_ = json.NewDecoder(r.Body).Decode(&updatedPost)

	// Gönderiyi güncelleyecek sorgu hazırlanıyor
	stmt, err := mainDB.Prepare("UPDATE posts SET id = ?, title = ?, body = ?, created_at = ?, updated_at = ?, user_id = ?, likes = ? WHERE id = ?")
	h.CheckErr(err)

	// Gönderiyi güncelleyecek sorgu çalıştırılıyor
	result, errExec := stmt.Exec(updatedPost.ID, updatedPost.Title, updatedPost.Body, updatedPost.CreatedAt, updatedPost.UpdatedAt, updatedPost.UserID, updatedPost.Likes, params["id"])
	h.CheckErr(errExec)

	rowAffected, errLast := result.RowsAffected()
	h.CheckErr(errLast)

	// Gönderinin veritabanında olup olmadığı kontrol ediliyor
	if rowAffected == 0 {
		json.NewEncoder(w).Encode(h.Error("Gönderi bulunamadı!"))
		return
	} else {
		json.NewEncoder(w).Encode(updatedPost)
	}
}

// Gönderi silme
func DeletePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "aplication/json")

	// Silinecek gönderinin id' si istek parametresinden alınıyor
	params := mux.Vars(r)

	// Gönderiyi silecek sorgu hazırlanıyor
	stmt, err := mainDB.Prepare("DELETE FROM posts WHERE id = ?")
	h.CheckErr(err)

	// Gönderiyi silecek sorgu çalıştırılıyor
	result, errExec := stmt.Exec(params["id"])
	h.CheckErr(errExec)

	rows, errRow := result.RowsAffected()
	h.CheckErr(errRow)

	// Gönderinin veritabanında olup olmadığı kontrol ediliyor
	if rows == 0 {
		json.NewEncoder(w).Encode(h.Error("Gönderi bulunamadı!"))
		return
	} else {
		json.NewEncoder(w).Encode("Gönderi silindi!")
	}
}

// Tüm yorumları getirme
func GetAllComments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "aplication/json")

	var comments m.GetComments

	// Yorumları getirecek sorgu çalıştırılıyor
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

// Yeni yorum ekleme
func AddComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "aplication/json")

	var addComment m.AddComment

	// Yorum içeriği istek gövdesinden alınıyor
	_ = json.NewDecoder(r.Body).Decode(&addComment)

	// Veritabanına yorum ekleyecek sorgu hazırlanıyor
	stmt, err := mainDB.Prepare("INSERT INTO comments( comment, post_id, user_id) VALUES(?, ?, ?)")
	h.CheckErr(err)

	// Veritabanına yorum ekleyecek sorgu çalıştırılıyor
	result, errExec := stmt.Exec(&addComment.Comment, &addComment.PostID, &addComment.UserID)
	h.CheckErr(errExec)

	rows, errRow := result.RowsAffected()
	h.CheckErr(errRow)

	// Hata kontrolü yapılıyor
	if rows == 0 {
		json.NewEncoder(w).Encode(h.Error("Bir terslik oldu!"))
		return
	} else {
		json.NewEncoder(w).Encode("Yorum eklendi.")
	}
}
