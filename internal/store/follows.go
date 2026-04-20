package store

import "database/sql"

type Follow struct {
	ID         int64 `json:"id"`
	FollowedId int64 `json:"followed_id"`
	FollowerId int64 `json:"follower_id"`
}

type FollowPostgresStore struct {
	db *sql.DB
}

func (pgs *FollowPostgresStore) Create(followData *Follow) error {
	query := `
	INSERT INTO follows (followed_id, follower_id)
	VALUES ($1, $2)
	`

	_, err := pgs.db.Exec(query, followData.FollowedId, followData.FollowerId)

	if err != nil {
		return err
	}

	return nil
}

func (pgs *FollowPostgresStore) Delete(followData *Follow) error {
	query := `
	DELETE FROM follows
	WHERE followed_id = $1
	AND
	follower_id = $2
	`

	_, err := pgs.db.Exec(query, followData.FollowedId, followData.FollowerId)

	if err != nil {
		return err
	}

	return nil
}
