package database

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

type HandlerDB struct {
	db *sql.DB
}

func Init() *HandlerDB {
	db, err := sql.Open("sqlite", "./blurt.db")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(v1)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	db.SetMaxOpenConns(1)

	return &HandlerDB{
		db: db,
	}
}

func (h *HandlerDB) Close() {
	h.db.Close()
}
