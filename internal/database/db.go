package database

import (
	"database/sql"
	"log"
	"time"

	"github.com/IhsanAlhakim/socmed-backend-go/internal/config"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func New(dbConfig config.DBConfig) (*sql.DB, error) {
	db, err := sql.Open("pgx", dbConfig.Dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(dbConfig.MaxOpenConns)
	db.SetMaxIdleConns(dbConfig.MaxIdleConns)

	duration, err := time.ParseDuration(dbConfig.MaxIdleTime)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxIdleTime(duration)

	if err := db.Ping(); err != nil {
		return nil, err
	}

	log.Println("Connected to database")
	return db, nil
}
