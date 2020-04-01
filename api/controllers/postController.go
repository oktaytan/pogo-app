package controllers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"pogo/api/db"
	h "pogo/api/helpers"
	m "pogo/api/models"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	_ "github.com/mattn/go-sqlite3"
)

// Veritabanı bağlantısı yapılıyor
var mainDB = db.InitDB()

var upgrader = websocket.Upgrader{}

func Live(w http.ResponseWriter, r *http.Request) {

	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	c, err := upgrader.Upgrade(w, r, nil)
	h.CheckErr(err)
	defer c.Close()

	for {
		mt, message, err := c.ReadMessage()
		h.CheckErr(err)
		log.Printf("recv: %s", message)

		err = c.WriteMessage(mt, message)
		h.CheckErr(err)
	}
}

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
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(h.Error("Gönderi bulunamadı!"))
		return
	} else {
		json.NewEncoder(w).Encode(post)
	}
}

// En çok beğeni alan gönderiler
func MostLikedPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "aplication/json")

	// Gönderiler için sql sorgusu çalıştırılıyor
	rows, err := mainDB.Query("SELECT * FROM posts AS p INNER JOIN (SELECT id, username, email FROM users) AS u ON p.user_id = u.id ORDER BY p.likes DESC")

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

// Göderi Arama
func SearchPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "aplication/json")

	type SearchText struct {
		Search string `json:"search"`
	}

	var searchText SearchText

	// Gönderini içeriği istek gövdesinden alınıyor
	_ = json.NewDecoder(r.Body).Decode(&searchText)

	searchText.Search = "%" + searchText.Search + "%"

	// Aranacak gönderiler için sql sorgusu hazırlanıyor
	stmt, err := mainDB.Prepare("SELECT * FROM posts AS p INNER JOIN (SELECT id, username, email FROM users) AS u ON p.user_id = u.id WHERE p.title LIKE ? OR p.body LIKE ? ORDER BY p.created_at DESC")
	h.CheckErr(err)

	// Aranacak gönderiler için sql sorgusu çalıştırılıyor
	rows, errQuery := stmt.Query(searchText.Search, searchText.Search)
	h.CheckErr(errQuery)

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

// Gönderi oluşturma
func CreatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "aplication/json")

	var post m.SendPost
	var result m.Success

	// Gönderini içeriği istek gövdesinden alınıyor
	_ = json.NewDecoder(r.Body).Decode(&post)

	// Yeni gönderiyi veritabanına ekleyecek sorgu hazırlanıyor
	stmt, err := mainDB.Prepare("INSERT INTO posts(title, body, user_id) VALUES (?, ?, ?)")
	h.CheckErr(err)

	// Yeni gönderiyi veritabanına ekleyecek sorgu çalıştırılıyor
	_, errExec := stmt.Exec(post.Title, post.Body, post.UserID)
	h.CheckErr(errExec)

	result.Success = true
	result.Message = "Gönderi paylaşıldı."

	post.Result = result

	json.NewEncoder(w).Encode(post)
}

// Gönderi beğenme
func UpdatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "aplication/json")

	// Güncellenecek gönderinin id' si istek parametresinden alınıyor
	params := mux.Vars(r)

	var likes m.Likes

	// Güncellenecek gönderinin yeni içeriği istek gövdesinden alınıyor
	_ = json.NewDecoder(r.Body).Decode(&likes)

	// Gönderiyi güncelleyecek sorgu hazırlanıyor
	stmt, err := mainDB.Prepare("UPDATE posts SET likes = ? WHERE id = ?")
	h.CheckErr(err)

	// Gönderiyi güncelleyecek sorgu çalıştırılıyor
	result, errExec := stmt.Exec(likes.LikeCount, params["id"])
	h.CheckErr(errExec)

	rowAffected, errLast := result.RowsAffected()
	h.CheckErr(errLast)

	var likeResult m.LikesResult
	var resultLike m.Success

	resultLike.Success = true
	resultLike.Message = "Gönderi Beğenildi"

	likeResult.Result = resultLike
	likeResult.PostID = params["id"]
	likeResult.Likes = likes.LikeCount

	// Gönderinin veritabanında olup olmadığı kontrol ediliyor
	if rowAffected == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(h.Error("Gönderi bulunamadı!"))
		return
	} else {
		json.NewEncoder(w).Encode(likeResult)
	}
}

// Kullanıcı beğenileri ayarlama
func LikedByUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "aplication/json")

	var likedUser m.LikedUser

	// Güncellenecek beğeni durumunun yeni içeriği istek gövdesinden alınıyor
	_ = json.NewDecoder(r.Body).Decode(&likedUser)

	// Mevcut beğeni var mı diye bakılıyor
	stmtHas, errHas := mainDB.Prepare("SELECT * FROM user_likes WHERE post_id = ? AND user_id = ?")
	h.CheckErr(errHas)

	rowsHas, errHasExec := stmtHas.Query(&likedUser.PostID, &likedUser.UserID)
	h.CheckErr(errHasExec)

	for rowsHas.Next() {
		err := rowsHas.Scan(&likedUser.ID, &likedUser.PostID, &likedUser.UserID, &likedUser.Liked)
		h.CheckErr(err)
	}

	if likedUser.ID == 0 {
		likedUser.Liked = 1

		// Yeni gönderiyi veritabanına ekleyecek sorgu hazırlanıyor
		stmtInsert, errInsert := mainDB.Prepare("INSERT INTO user_likes(post_id, user_id, liked) VALUES (?, ?, ?)")
		h.CheckErr(errInsert)

		// Yeni gönderiyi veritabanına ekleyecek sorgu çalıştırılıyor
		resultInsert, errExecInsert := stmtInsert.Exec(likedUser.PostID, likedUser.UserID, likedUser.Liked)
		h.CheckErr(errExecInsert)

		rowsInsert, errRowInsert := resultInsert.RowsAffected()
		h.CheckErr(errRowInsert)

		// Gönderinin veritabanında olup olmadığı kontrol ediliyor
		if rowsInsert == 0 {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(h.Error("Gönderi bulunamadı!"))
			return
		} else {
			json.NewEncoder(w).Encode(likedUser)
		}
	} else {
		// Eğer gönderi beğenilmişse "liked" değeri "0" olarak, aksi takdirde "1" olarak ayarlanıyor
		if likedUser.Liked == 0 {
			likedUser.Liked = 1
		} else {
			likedUser.Liked = 0
		}

		// Yeni gönderiyi veritabanına ekleyecek sorgu hazırlanıyor
		stmt, err := mainDB.Prepare("UPDATE user_likes SET liked = ? WHERE post_id = ? AND user_id = ?")
		h.CheckErr(err)

		// Yeni gönderiyi veritabanına ekleyecek sorgu çalıştırılıyor
		result, errExec := stmt.Exec(likedUser.Liked, likedUser.PostID, likedUser.UserID)
		h.CheckErr(errExec)

		rows, errRow := result.RowsAffected()
		h.CheckErr(errRow)

		// Gönderinin veritabanında olup olmadığı kontrol ediliyor
		if rows == 0 {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(h.Error("Gönderi bulunamadı!"))
			return
		} else {
			json.NewEncoder(w).Encode(likedUser)
		}
	}
}

// Kullanıcının beğenileri
func GetLikes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "aplication/json")

	// Beğenileri getirilecek kullanıcının id' si istek parametresinden alınıyor
	params := mux.Vars(r)

	var likedPosts m.LikedPosts
	var userLikes m.UserLikes

	// Kullanıcının beğenilerini getirecek sorgu hazırlanıyor
	stmt, err := mainDB.Prepare("SELECT * FROM user_likes WHERE user_id = ?")
	h.CheckErr(err)

	// Kullanıcının beğenilerini getirecek sorgu calıştırılıyor
	rows, errExec := stmt.Query(params["id"])
	h.CheckErr(errExec)

	for rows.Next() {
		var likedPost m.LikedPost

		err := rows.Scan(&likedPost.ID, &likedPost.PostID, &userLikes.UserID, &likedPost.Liked)
		h.CheckErr(err)

		likedPosts = append(likedPosts, likedPost)
	}
	userLikes.UserID = params["id"]

	userLikes.LikedPosts = likedPosts

	json.NewEncoder(w).Encode(userLikes)
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
		w.WriteHeader(http.StatusNotFound)
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

	if addComment.Comment == "" {
		json.NewEncoder(w).Encode(h.Error("Yorum boş bırakılamaz!"))
	} else {
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
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(h.Error("Bir terslik oldu!"))
			return
		} else {
			json.NewEncoder(w).Encode("Yorum eklendi.")
		}
	}
}

// Kendi yorumunu silme
func DeleteComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "aplication/json")

	// Silinecek yorumun id' si istek parametresinden alınıyor
	params := mux.Vars(r)

	// Yorumu silecek sorgu hazırlanıyor
	stmt, err := mainDB.Prepare("DELETE FROM comments WHERE id = ?")
	h.CheckErr(err)

	// Yorumu silecek sorgu çalıştırılıyor
	result, errExec := stmt.Exec(params["id"])
	h.CheckErr(errExec)

	rows, errRow := result.RowsAffected()
	h.CheckErr(errRow)

	// GönderYorumuninin veritabanında olup olmadığı kontrol ediliyor
	if rows == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(h.Error("Yorum bulunamadı!"))
		return
	} else {
		json.NewEncoder(w).Encode("Yorum silindi!")
	}
}
