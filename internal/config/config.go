package config

import (
	"log"

	"github.com/IhsanAlhakim/socmed-backend-go/internal/env"
	"github.com/joho/godotenv"
)

type DBConfig struct {
	Dsn string
}

type Config struct {
	Port     string
	DBConfig DBConfig
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Failed to load .env file", err)
	}

	return &Config{
		Port: env.GetString("PORT", "8000"),
		DBConfig: DBConfig{
			Dsn: env.GetString("DSN", "postgres://postgres:admin123@localhost:5432/socmed"),
		},
	}
}
