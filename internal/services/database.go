package services

import (
	"database/sql"
	"log"
)

const DB_NAME = "data.db"

// TODO: Make this into a singleton.

func NewDatabase() *sql.DB {
	database, err := sql.Open("sqlite3", DB_NAME)
	if err != nil {
		log.Fatalf("failed to open database: %s", err)
	}
	return database
}
