package config

import (
	"log"

	"github.com/IhsanAlhakim/socmed-backend-go/internal/env"
	"github.com/joho/godotenv"
)

type DBConfig struct {
	Dsn          string
	DbDriver     string
	MaxOpenConns int
	MaxIdleConns int
	MaxIdleTime  string
}

type Config struct {
	AppName  string
	Port     string
	DBConfig DBConfig

	JWTSignKey      string
	JWTContextKey   string
	TokenCookieName string
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Failed to load .env file", err)
	}

	return &Config{
		AppName: env.GetString("APPNAME", "socmed"),
		Port:    env.GetString("PORT", "8000"),
		DBConfig: DBConfig{
			Dsn:          env.GetString("DSN", "postgres://postgres:admin123@localhost:5432/socmed"),
			DbDriver:     env.GetString("DB_DRIVER", "pgx"),
			MaxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			MaxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			MaxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
		JWTSignKey:      env.GetString("JWT_SIGNATURE_KEY", "very-secret-key"),
		JWTContextKey:   env.GetString("JWT_CONTEXT_KEY", "userInfo"),
		TokenCookieName: env.GetString("TOKEN_COOKIE_NAME", "cookie-token-name"),
	}
}
