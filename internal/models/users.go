package models

import (
	"database/sql"
)

type Users struct {
	database *sql.DB
}

func NewUsers(database *sql.DB) *Users {
	return &Users{
		database: database,
	}
}

func (users *Users) Migrate() error {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		username TEXT PRIMARY KEY,
		password TEXT NOT NULL
	)`

	_, err := users.database.Exec(query)
	return err
}

func (users *Users) Create(username, password string) error {
	_, err := users.database.Exec("INSERT INTO users (username, password) VALUES (?, ?)")
	return err
}

func (users *Users) Exists(username, password string) (bool, error) {
	// TODO: Check if a row with the given username and password exists.
	return false, nil
}
