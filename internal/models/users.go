package models

import (
	"database/sql"
	"errors"
)

type Users struct {
	database *sql.DB
}

// TODO: Make this into a singleton.

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
	query := "SELECT EXISTS (SELECT 1 FROM users WHERE username = ? AND password = ? LIMIT 1)"
	row := users.database.QueryRow(query, username, password)

	var exists bool
	if err := row.Scan(&exists); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, err
	}
	return exists, nil
}
