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
