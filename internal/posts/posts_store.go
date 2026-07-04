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

func (pgs *PostgresStore) GetPosts(userId int64) (*[]Post, error) {

	// query := `
	// SELECT p.id, p.user_id, u.username as creator, p.content, p.created_at,
	// EXISTS(SELECT pl.id from post_likes pl WHERE pl.post_id = p.id AND pl.user_id = $1 LIMIT 1) as liked
	// FROM posts p
	// JOIN users u ON p.user_id = u.id
	// `

	query := `
	SELECT p.id, p.user_id, u.username as creator, p.content, p.created_at,
	(SELECT COUNT(*) FROM post_likes pl WHERE pl.post_id = p.id AND pl.deleted = false) as like_count,
	(SELECT COUNT(*) FROM comments c WHERE c.post_id = p.id) as comment_count, 
	CASE WHEN pl.deleted = FALSE THEN TRUE
	ELSE FALSE
	END AS liked
	FROM posts p
	JOIN users u ON p.user_id = u.id
	LEFT JOIN post_likes pl ON p.id = pl.post_id AND pl.user_id = $1
	`
	rows, err := pgs.db.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []Post

	for rows.Next() {
		var each Post
		err := rows.Scan(&each.ID, &each.UserId, &each.Creator, &each.Content, &each.CreatedAt, &each.LikeCount, &each.CommentCount, &each.Liked)
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

func (pgs *PostgresStore) GetPostsByUsername(userId int64, otherUsername string) (*[]Post, error) {

	query := `
	SELECT p.id, p.user_id, u.username as creator, p.content, p.created_at,
	CASE WHEN pl.deleted = FALSE THEN TRUE
	ELSE FALSE
	END AS liked
	FROM posts p
	JOIN users u ON p.user_id = u.id
	LEFT JOIN post_likes pl ON p.id = pl.post_id AND pl.user_id = $1
	WHERE u.username = $2
	`
	rows, err := pgs.db.Query(query, userId, otherUsername)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []Post

	for rows.Next() {
		var each Post
		err := rows.Scan(&each.ID, &each.UserId, &each.Creator, &each.Content, &each.CreatedAt, &each.Liked)
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
	SELECT p.id, p.user_id, u.username as creator, p.content, p.created_at,
	CASE WHEN pl.deleted = FALSE THEN TRUE
	ELSE FALSE
	END AS liked
	FROM posts p
	JOIN users u ON p.user_id = u.id
	LEFT JOIN post_likes pl ON p.id = pl.post_id AND pl.user_id = $1
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
		err := rows.Scan(&each.ID, &each.UserId, &each.Creator, &each.Content, &each.CreatedAt, &each.Liked)
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

func (pgs *PostgresStore) GetPostById(postId int64, userId int64) (*Post, error) {

	query := `
	SELECT p.id, p.user_id, u.username as creator, p.content, p.created_at,
	CASE WHEN pl.deleted = FALSE THEN TRUE
	ELSE FALSE
	END AS liked
	FROM posts p
	JOIN users u ON p.user_id = u.id
	LEFT JOIN post_likes pl ON p.id = pl.post_id AND pl.user_id = $1
	WHERE p.id = $2
	`

	var result Post

	err := pgs.db.QueryRow(query, userId, postId).Scan(&result.ID, &result.UserId, &result.Creator, &result.Content, &result.CreatedAt, &result.Liked)
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
	SELECT p.id, p.user_id, u.username as creator, p.content, p.created_at,
	CASE WHEN pl.deleted = FALSE THEN TRUE
	ELSE FALSE
	END AS liked
	FROM posts p
	JOIN users u ON p.user_id = u.id
	LEFT JOIN post_likes pl ON p.id = pl.post_id AND pl.user_id = $1
	WHERE p.id IN (SELECT post_id from post_likes WHERE user_id = $1 AND deleted = FALSE)
	`
	rows, err := pgs.db.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []Post

	for rows.Next() {
		var each Post
		err := rows.Scan(&each.ID, &each.UserId, &each.Creator, &each.Content, &each.CreatedAt, &each.Liked)
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
