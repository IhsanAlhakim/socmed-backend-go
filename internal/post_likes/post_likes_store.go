package plikes

import (
	"database/sql"
	"strings"

	"github.com/IhsanAlhakim/socmed-backend-go/internal/posts"
	"github.com/IhsanAlhakim/socmed-backend-go/internal/users"
)

func NewStore(db *sql.DB) StoreInterface {
	return &PostgresStore{db: db}
}

type PostgresStore struct {
	db *sql.DB
}

func (pgs *PostgresStore) LikePost(postId int64, userId int64) error {
	query := `
	INSERT INTO post_likes (post_id, user_id)
	VALUES ($1, $2)
	`

	_, err := pgs.db.Exec(query, postId, userId)

	if err != nil {
		switch {
		case strings.Contains(err.Error(), `duplicate key value violates unique constraint "post_likes_postid_userid_unique"`):
			return ErrPostAlreadyLiked
		case strings.Contains(err.Error(), `insert or update on table "post_likes" violates foreign key constraint "post_likes_user_id_fkey`):
			return users.ErrUserNotFound
		case strings.Contains(err.Error(), `insert or update on table "post_likes" violates foreign key constraint "post_likes_post_id_fkey"`):
			return posts.ErrPostNotFound
		default:
			return err
		}
	}

	return nil
}

func (pgs *PostgresStore) UnlikePost(postId int64, userId int64) error {
	query := `
	DELETE FROM post_likes
	WHERE post_id = $1
	AND user_id = $2
	`

	result, err := pgs.db.Exec(query, postId, userId)

	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return ErrLikeNotFound
	}

	return nil
}

func (pgs *PostgresStore) GetPostLiker(postId int64) (*[]PostLike, error) {
	query := `
	SELECT pl.user_id, u.username 
	FROM post_likes pl
	JOIN users u ON pl.user_id = u.id
	WHERE pl.post_id = $1
	`

	rows, err := pgs.db.Query(query, postId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []PostLike

	for rows.Next() {
		var each PostLike
		err := rows.Scan(&each.UserId, &each.Username)
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

func (pgs *PostgresStore) GetPostLikesCount(postId int64) (*PostLikesCount, error) {
	query := `
	SELECT COUNT(user_id) 
	FROM post_likes
	WHERE post_id = $1
	`

	var result PostLikesCount

	err := pgs.db.QueryRow(query, postId).Scan(&result.LikesCount)
	if err != nil {
		return &PostLikesCount{}, err
	}

	return &result, nil
}
