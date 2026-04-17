package store

import "database/sql"

type Storage struct {
	Users interface {
		Create(userData *User) error
	}
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{
		Users: &UsersPostgresStore{db: db},
	}
}
