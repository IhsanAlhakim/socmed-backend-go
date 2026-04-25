package posts

import (
	"database/sql"
)

func NewStore(db *sql.DB) StoreInterface {
	return &PostgresStore{db: db}
}

type PostgresStore struct {
	db *sql.DB
}

func (pgs *PostgresStore) Create(payload *CreatePostParam) error {
	query := `
	INSERT INTO posts (user_id, title, content)
	VALUES ($1, $2, $3)
	`

	_, err := pgs.db.Exec(query, payload.UserId, payload.Title, payload.Content)

	if err != nil {
		return err
	}

	return nil
}

func (pgs *PostgresStore) Delete(postId int64) error {

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

func (pgs *PostgresStore) GetPosts() (*[]Post, error) {

	query := `
	SELECT p.id, u.username as creator, p.title, p.content, p.created_at
	FROM posts p
	JOIN users u ON p.user_id = u.id
	`
	rows, err := pgs.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []Post

	for rows.Next() {
		var each Post
		err := rows.Scan(&each.ID, &each.Creator, &each.Title, &each.Content, &each.CreatedAt)
		if err != nil {
			return nil, err
		}
		result = append(result, each)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &result, nil
}

func (pgs *PostgresStore) GetFollowedPosts(followerId int64) (*[]Post, error) {

	query := `
	SELECT p.id, u.username as creator, p.title, p.content, p.created_at
	FROM posts p
	JOIN users u ON p.user_id = u.id
	WHERE p.user_id IN (SELECT followed_id FROM follows WHERE follower_id = $1)
	`
	rows, err := pgs.db.Query(query, followerId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []Post

	for rows.Next() {
		var each Post
		err := rows.Scan(&each.ID, &each.Creator, &each.Title, &each.Content, &each.CreatedAt)
		if err != nil {
			return nil, err
		}
		result = append(result, each)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &result, nil
}

func (pgs *PostgresStore) GetById(postId int64) (*Post, error) {

	query := `
	SELECT id, title, user_id, content, created_at
	FROM posts
	WHERE id = $1
	`

	var result Post

	err := pgs.db.QueryRow(query, postId).Scan(&result.ID, &result.Title, &result.UserId, &result.Content, &result.CreatedAt)
	if err != nil {
		return &Post{}, err
	}

	return &result, nil
}
