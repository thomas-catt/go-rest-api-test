package utils

import (
	"database/sql"
	"log"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "users.db")

	if err != nil {
		log.Fatalln(err)
		return
	}

	q := `CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT NOT NULL,
		password TEXT NOT NULL
	)`

	_, err = DB.Exec(q)

	if err != nil {
		log.Fatalln(err)
		return
	}
}
