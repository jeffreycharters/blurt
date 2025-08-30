package db

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

func Init() *sql.DB {
	db, err := sql.Open("sqlite", "./blurt.db")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(v1)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
