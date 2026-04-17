package store

import (
	"database/sql"
	"time"
)

type User struct {
	ID       int64     `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	CreateAt time.Time `json:"created_at"`
}

type UsersPostgresStore struct {
	db *sql.DB
}

func (pgs *UsersPostgresStore) Create(user *User) error {
	query := `
	INSERT INTO users (username, password, email)
	VALUES ($1, $2, $3)
	`

	_, err := pgs.db.Exec(query, user.Username, user.Password, user.Email)

	if err != nil {
		return err
	}

	return nil
}
