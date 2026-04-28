package users

import (
	"database/sql"
)

func NewStore(db *sql.DB) StoreInterface {
	return &PostgresStore{db: db}
}

type PostgresStore struct {
	db *sql.DB
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
		return &User{}, err
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
		return err
	}

	return nil
}

func (pgs *PostgresStore) UpdateUser(userId int64, payload *UpdateUserParam) error {

	query := `
	UPDATE users 
	SET username = $1, email = $2 
	WHERE id = $3
	`
	_, err := pgs.db.Exec(query, payload.Username, payload.Email, userId)
	if err != nil {
		return err
	}

	return nil
}

func (pgs *PostgresStore) DeleteUser(userId int64) error {

	query := `
	DELETE FROM users 
	WHERE id = $1
	`
	_, err := pgs.db.Exec(query, userId)
	if err != nil {
		return err
	}

	return nil
}
