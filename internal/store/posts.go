package store

import (
	"database/sql"
	"time"
)

type Post struct {
	ID        int64     `json:"id"`
	UserId    int64     `json:"user_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

type PostsPostgresStore struct {
	db *sql.DB
}

func (pgs *PostsPostgresStore) Create(post *Post) error {
	query := `
	INSERT INTO posts (user_id, title, content)
	VALUES ($1, $2, $3)
	`

	_, err := pgs.db.Exec(query, post.UserId, post.Title, post.Content)

	if err != nil {
		return err
	}

	return nil
}

func (pgs *PostsPostgresStore) Delete(postId int64) error {

	query := `
	DELETE FROM posts 
	WHERE id = $1
	`
	_, err := pgs.db.Exec(query, postId)
	if err != nil {
		return err
	}

	return nil
}
