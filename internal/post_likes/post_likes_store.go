package plikes

import "database/sql"

func NewStore(db *sql.DB) StoreInterface {
	return &PostgresStore{db: db}
}

type PostgresStore struct {
	db *sql.DB
}

func (pgs *PostgresStore) LikePost(postLikeData *PostLikeParam) error {
	query := `
	INSERT INTO post_likes (post_id, user_id)
	VALUES ($1, $2)
	`

	_, err := pgs.db.Exec(query, postLikeData.PostId, postLikeData.UserId)

	if err != nil {
		return err
	}

	return nil
}

func (pgs *PostgresStore) UnlikePost(postLikeData *PostLikeParam) error {
	query := `
	DELETE FROM post_likes
	WHERE post_id = $1
	AND user_id = $2
	`

	_, err := pgs.db.Exec(query, postLikeData.PostId, postLikeData.UserId)

	if err != nil {
		return err
	}

	return nil
}
