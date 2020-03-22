package models

// Post Author sutruct
type Author struct {
	ID       string `json:"id"`
	UserName string `json:"username"`
	Email    string `json:"email"`
}

// Post Struct
type Post struct {
	ID        string   `json:"id"`
	Title     string   `json:"title"`
	Body      string   `json:"body"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
	Likes     string   `json:"likes"`
	Author    Author   `json:"author"`
	Comments  Comments `json:"comments"`
}

type Posts []Post

// Own Post Struct
type OwnPost struct {
	ID        string   `json:"id"`
	Title     string   `json:"title"`
	Body      string   `json:"body"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
	Likes     string   `json:"likes"`
	Comments  Comments `json:"comments"`
}

type OwnPosts []OwnPost

// Updated Post Struct
type UpdatedPost struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	UserID    string `json:"user_id"`
	Likes     string `json:"likes"`
}

// Send Post Struct
type SendPost struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserID string `json:"user_id"`
}

// Error Struct
type Error struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

// Post Comment Struct
type Comment struct {
	ID            string `json:"id"`
	Comment       string `json:"comment"`
	CommentAuthor Author `json:"author"`
}

type Comments []Comment

// Comment Struct
type AddComment struct {
	Comment string `json:"comment"`
	PostID  string `json:"post_id"`
	UserID  string `json:"user_id"`
}

// Comment Struct
type GetComment struct {
	ID            string `json:"id"`
	Comment       string `json:"comment"`
	PostID        string `json:"post_id"`
	UserID        string `json:"user_id"`
	CommentAuthor Author `json:"author"`
}

type GetComments []GetComment
