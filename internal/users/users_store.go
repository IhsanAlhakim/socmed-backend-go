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

func (pgs *PostgresStore) Create(user *User) error {
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

func (pgs *PostgresStore) Update(userId int64, updatedUser *User) error {

	query := `
	UPDATE users 
	SET username = $1, email = $2 
	WHERE id = $3
	`
	_, err := pgs.db.Exec(query, updatedUser.Username, updatedUser.Email, userId)
	if err != nil {
		return err
	}

	return nil
}

func (pgs *PostgresStore) Delete(userId int64) error {

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
