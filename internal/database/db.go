package database

import (
	"database/sql"
	"log"

	"github.com/IhsanAlhakim/socmed-backend-go/internal/config"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func New(dbConfig config.DBConfig) (*sql.DB, error) {
	db, err := sql.Open(dbConfig.DbDriver, dbConfig.Dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	log.Println("Connected to database")
	return db, nil
}
