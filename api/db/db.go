package db

import (
	"database/sql"
	h "pogo/api/helpers"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB() *sql.DB {

	// Sqlite database connection
	db, errOpenDB := sql.Open("sqlite3", "./pogo.db")
	h.CheckErr(errOpenDB)

	return db
}
