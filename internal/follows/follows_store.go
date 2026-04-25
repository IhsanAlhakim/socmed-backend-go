package follows

import "database/sql"

func NewStore(db *sql.DB) StoreInterface {
	return &PostgresStore{db: db}
}

type PostgresStore struct {
	db *sql.DB
}

func (pgs *PostgresStore) GetFollower(followedId int64) (*[]Follow, error) {
	// query2 := `
	// SELECT f.followed_id, u1.username as followed_name, f.follower_id, u2.username as follower_name
	// FROM follows f
	// JOIN users u1 ON f.followed_id = u1.id
	// JOIN users u2 ON f.follower_id = u2.id
	// WHERE followed_id = $1
	// `

	query := `
	SELECT f.follower_id, u.username as follower_name
	FROM follows f
	JOIN users u ON f.follower_id = u.id
	WHERE followed_id = $1
	`

	rows, err := pgs.db.Query(query, followedId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []Follow

	for rows.Next() {
		var each Follow
		err := rows.Scan(&each.FollowerId, &each.FollowerName)
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

func (pgs *PostgresStore) GetFollowed(followerId int64) (*[]Follow, error) {

	query := `
	SELECT f.followed_id, u.username as followed_name
	FROM follows f
	JOIN users u ON f.followed_id = u.id
	WHERE follower_id = $1
	`

	rows, err := pgs.db.Query(query, followerId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []Follow

	for rows.Next() {
		var each Follow
		err := rows.Scan(&each.FollowedId, &each.FollowedName)
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

func (pgs *PostgresStore) Create(payload *FollowParam) error {
	query := `
	INSERT INTO follows (followed_id, follower_id)
	VALUES ($1, $2)
	`

	_, err := pgs.db.Exec(query, payload.FollowedId, payload.FollowerId)

	if err != nil {
		return err
	}

	return nil
}

func (pgs *PostgresStore) Delete(payload *FollowParam) error {
	query := `
	DELETE FROM follows
	WHERE followed_id = $1
	AND
	follower_id = $2
	`

	_, err := pgs.db.Exec(query, payload.FollowedId, payload.FollowerId)

	if err != nil {
		return err
	}

	return nil
}
