package db

import (
	"database/sql"
	_ "github.com/lib/pq"
)

var db *sql.DB

func DB() *sql.DB {
	if db != nil {
		return db
	}
	db, _ = sql.Open("postgres",
		"host=db user=postgres password=postgres dbname=app sslmode=disable")
	return db
}