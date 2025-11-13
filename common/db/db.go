package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var dbs = ConnectDb()

func ConnectDb() *sql.DB {
	db, err := sql.Open("sqlite3", "./EXCLUSIVE_BASE_QZ.db")
	checkErr(err)
	return db
}
