package posts

import (
	"database/sql"
	"strings"

	"github.com/IhsanAlhakim/socmed-backend-go/internal/users"
)

func NewStore(db *sql.DB) StoreInterface {
	return &PostgresStore{db: db}
}

type PostgresStore struct {
	db *sql.DB
}

func (pgs *PostgresStore) CreatePost(userId int64, payload *CreatePostParam) error {
	query := `
	INSERT INTO posts (user_id, content)
	VALUES ($1, $2)
	`

	_, err := pgs.db.Exec(query, userId, payload.Content)

	if err != nil {
		if strings.Contains(err.Error(), `insert or update on table "posts" violates foreign key constraint "fk_posts_user_id"`) {
			return users.ErrUserNotFound
		}
		return err
	}

	return nil
}

func (pgs *PostgresStore) DeletePost(postId int64) error {

	query := `
	DELETE FROM posts 
	WHERE id = $1
	`
	result, err := pgs.db.Exec(query, postId)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return ErrPostNotFound
	}

	return nil
}

func (pgs *PostgresStore) GetPosts() (*[]Post, error) {

	query := `
	SELECT p.id, u.username as creator, p.content, p.created_at
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
		err := rows.Scan(&each.ID, &each.Creator, &each.Content, &each.CreatedAt)
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

func (pgs *PostgresStore) GetFollowedPosts(userId int64) (*[]Post, error) {

	query := `
	SELECT p.id, u.username as creator, p.content, p.created_at
	FROM posts p
	JOIN users u ON p.user_id = u.id
	WHERE p.user_id IN (SELECT followed_id FROM follows WHERE follower_id = $1)
	`
	rows, err := pgs.db.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []Post

	for rows.Next() {
		var each Post
		err := rows.Scan(&each.ID, &each.Creator, &each.Content, &each.CreatedAt)
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
	SELECT id, user_id, content, created_at
	FROM posts
	WHERE id = $1
	`

	var result Post

	err := pgs.db.QueryRow(query, postId).Scan(&result.ID, &result.UserId, &result.Content, &result.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return &Post{}, ErrPostNotFound
		}
		return &Post{}, err
	}

	return &result, nil
}

func (pgs *PostgresStore) GetLikedPosts(userId int64) (*[]Post, error) {

	query := `
	SELECT p.id, u.username as creator, p.content, p.created_at
	FROM posts p
	JOIN users u ON p.user_id = u.id
	WHERE p.id IN (SELECT post_id from post_likes WHERE user_id = $1)
	`
	rows, err := pgs.db.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []Post

	for rows.Next() {
		var each Post
		err := rows.Scan(&each.ID, &each.Creator, &each.Content, &each.CreatedAt)
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
