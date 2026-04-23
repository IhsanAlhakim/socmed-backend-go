package comments

import (
	"database/sql"
)

func NewStore(db *sql.DB) StoreInterface {
	return &PostgresStore{db: db}
}

type PostgresStore struct {
	db *sql.DB
}

func (pgs *PostgresStore) CreateComment(commentData *CreateCommentParam) error {
	query := `
	INSERT INTO comments (user_id, post_id, content)
	VALUES ($1, $2, $3)
	`

	_, err := pgs.db.Exec(query, commentData.UserId, commentData.PostId, commentData.Content)

	if err != nil {
		return err
	}

	return nil
}
