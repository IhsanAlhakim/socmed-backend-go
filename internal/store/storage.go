package store

import "database/sql"

type Storage struct {
	Users interface {
		Create(userData *User) error
		Update(userId int64, userData *User) error
		Delete(userId int64) error
	}
	Posts interface {
		Create(postData *Post) error
		Delete(postId int64) error
		Get() ([]Post, error)
		GetById(postId int64) (Post, error)
	}
	Follows interface {
		Create(followData *Follow) error
		Delete(followData *Follow) error
	}
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{
		Users:   &UsersPostgresStore{db: db},
		Posts:   &PostsPostgresStore{db: db},
		Follows: &FollowPostgresStore{db: db},
	}
}
