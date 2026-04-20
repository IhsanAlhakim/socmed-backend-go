package store

import "database/sql"

type Storage struct {
	Users interface {
		Create(userData *User) error
		Update(userData *User) error
		Delete(userId int64) error
	}
	Posts interface {
		Create(postData *Post) error
		Delete(postId int64) error
	}
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{
		Users: &UsersPostgresStore{db: db},
		Posts: &PostsPostgresStore{db: db},
	}
}
