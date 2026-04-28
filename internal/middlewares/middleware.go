package middlewares

import "github.com/IhsanAlhakim/socmed-backend-go/internal/config"

type Middleware struct {
	config *config.Config
}

func New(config *config.Config) *Middleware {
	return &Middleware{config: config}
}
