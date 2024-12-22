package user

import (
	"database/sql"
	"time"
)

func (u *User) Create(db *sql.DB) (sql.Result, error) {
	password, err := HashPassword(u.Password)
	if err != nil {
		return nil, err
	}

	result, err := db.Exec("INSERT INTO users (username, password, created_at) VALUES ($1, $2, $3)", u.Username, password, time.Now())
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (u *User) Get(db *sql.DB) error {
	row := db.QueryRow("SELECT * FROM users WHERE username = $1", u.Username)
	err := row.Scan(&u.ID, &u.Username, &u.Password, &u.Created_at)
	if err != nil {
		return err
	}

	return nil
}

func (u *User) Delete(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM users WHERE username = $1", u.Username)
	if err != nil {
		return err
	}

	return nil
}