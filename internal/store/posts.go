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

func (pgs *PostsPostgresStore) Get() ([]Post, error) {

	query := `
	SELECT id, title, user_id, content, created_at
	FROM posts
	`
	rows, err := pgs.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []Post

	for rows.Next() {
		var each Post
		err := rows.Scan(&each.ID, &each.Title, &each.UserId, &each.Content, &each.CreatedAt)
		if err != nil {
			return nil, err
		}
		result = append(result, each)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func (pgs *PostsPostgresStore) GetById(postId int64) (Post, error) {

	query := `
	SELECT id, title, user_id, content, created_at
	FROM posts
	WHERE id = $1
	`

	var result Post

	err := pgs.db.QueryRow(query, postId).Scan(&result.ID, &result.Title, &result.UserId, &result.Content, &result.CreatedAt)
	if err != nil {
		return Post{}, err
	}

	return result, nil
}
