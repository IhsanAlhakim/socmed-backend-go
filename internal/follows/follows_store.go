package follows

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

func (pgs *PostgresStore) GetFollower(userId int64) (*[]Follow, error) {
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

	rows, err := pgs.db.Query(query, userId)
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

func (pgs *PostgresStore) GetFollowed(userId int64) (*[]Follow, error) {

	query := `
	SELECT f.followed_id, u.username as followed_name
	FROM follows f
	JOIN users u ON f.followed_id = u.id
	WHERE follower_id = $1
	`

	rows, err := pgs.db.Query(query, userId)
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

func (pgs *PostgresStore) Follow(userId int64, payload *FollowParam) error {
	query := `
	INSERT INTO follows (followed_id, follower_id)
	VALUES ($1, $2)
	`

	_, err := pgs.db.Exec(query, payload.FollowedId, userId)

	if err != nil {
		switch {
		case strings.Contains(err.Error(), `insert or update on table "follows" violates foreign key constraint "fk_follow_followed_id"`):
			return ErrFollowedNotFound
		case strings.Contains(err.Error(), `insert or update on table "follows" violates foreign key constraint "fk_follow_follower_id"`):
			return users.ErrUserNotFound
		case strings.Contains(err.Error(), `duplicate key value violates unique constraint "follows_followedid_followerid_unique"`):
			return ErrUserAlreadyFollowed
		default:
			return err
		}
	}

	return nil
}

func (pgs *PostgresStore) Unfollow(userId int64, payload *FollowParam) error {
	query := `
	DELETE FROM follows
	WHERE followed_id = $1
	AND
	follower_id = $2
	`

	result, err := pgs.db.Exec(query, payload.FollowedId, userId)

	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return users.ErrUserNotFound
	}

	return nil
}
