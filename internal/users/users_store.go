package users

import (
	"database/sql"
	"strings"
)

func NewStore(db *sql.DB) StoreInterface {
	return &PostgresStore{db: db}
}

type PostgresStore struct {
	db *sql.DB
}

func (pgs *PostgresStore) GetUserById(userId int64) (*User, error) {

	query := `
	SELECT id, username, email, created_at
	FROM users
	WHERE id = $1
	`

	var result User

	err := pgs.db.QueryRow(query, userId).Scan(&result.ID, &result.Username, &result.Email, &result.CreateAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return &User{}, ErrUserNotFound
		} else {
			return &User{}, err
		}
	}

	return &result, nil
}

func (pgs *PostgresStore) GetUserByUsername(loggedInUserId int64, username string) (*User, error) {

	// query := `
	// SELECT id, username, created_at
	// FROM users
	// WHERE username = $1
	// `

	query := `
	SELECT u.id, u.username, u.created_at,
	CASE WHEN fl.deleted = FALSE THEN TRUE
	ELSE FALSE
	END AS followed
	FROM users u
	LEFT JOIN follows fl ON u.id = fl.followed_id AND fl.follower_id = $1
	WHERE u.username = $2
	`

	var result User

	err := pgs.db.QueryRow(query, loggedInUserId, username).Scan(&result.ID, &result.Username, &result.CreateAt, &result.Followed)
	if err != nil {
		if err == sql.ErrNoRows {
			return &User{}, ErrUserNotFound
		} else {
			return &User{}, err
		}
	}

	return &result, nil
}

func (pgs *PostgresStore) GetUserByEmail(email string) (*User, error) {

	query := `
	SELECT id, password
	FROM users
	WHERE email = $1
	`

	var result User

	err := pgs.db.QueryRow(query, email).Scan(&result.ID, &result.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return &User{}, ErrUserNotFound
		} else {
			return &User{}, err
		}
	}

	return &result, nil
}

func (pgs *PostgresStore) CreateUser(user *CreateUserParam) error {
	query := `
	INSERT INTO users (username, password, email)
	VALUES ($1, $2, $3)
	`

	_, err := pgs.db.Exec(query, user.Username, user.Password, user.Email)

	if err != nil {
		switch {
		case strings.Contains(err.Error(), `duplicate key value violates unique constraint "users_username_key"`):
			return ErrDuplicateUsername
		case strings.Contains(err.Error(), `duplicate key value violates unique constraint "users_email_key"`):
			return ErrDuplicateEmail
		default:
			return err
		}
	}

	return nil
}

func (pgs *PostgresStore) UpdateUser(userId int64, payload *UpdateUserParam) error {

	query := `
	UPDATE users 
	SET username = $1, email = $2 
	WHERE id = $3
	`
	result, err := pgs.db.Exec(query, payload.Username, payload.Email, userId)

	if err != nil {
		switch {
		case strings.Contains(err.Error(), `duplicate key value violates unique constraint "users_username_key"`):
			return ErrDuplicateUsername
		case strings.Contains(err.Error(), `duplicate key value violates unique constraint "users_email_key"`):
			return ErrDuplicateEmail
		default:
			return err
		}
	}

	count, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if count == 0 {
		return ErrUserNotFound
	}

	return nil
}

func (pgs *PostgresStore) DeleteUser(userId int64) error {

	query := `
	DELETE FROM users 
	WHERE id = $1
	`
	result, err := pgs.db.Exec(query, userId)
	if err != nil {
		return err
	}

	count, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if count == 0 {
		return ErrUserNotFound
	}

	return nil
}
