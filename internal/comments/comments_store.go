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

func (pgs *PostgresStore) Getcomments(postId int64) (*[]Comment, error) {
	query := `
	SELECT c.id, u.username, c.content, c.created_at
	FROM comments c
	JOIN users u ON c.user_id = u.id
	WHERE c.post_id = $1 
	`
	rows, err := pgs.db.Query(query, postId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []Comment

	for rows.Next() {
		var each Comment
		err := rows.Scan(&each.ID, &each.Username, &each.Content, &each.CreatedAt)
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

func (pgs *PostgresStore) DeleteComment(commentId int64) error {
	query := `
	DELETE FROM comments
	WHERE id = $1
	`

	_, err := pgs.db.Exec(query, commentId)

	if err != nil {
		return err
	}

	return nil
}
