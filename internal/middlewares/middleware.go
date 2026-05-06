package middlewares

import (
	"github.com/IhsanAlhakim/socmed-backend-go/internal/auth"
)

type Middleware struct {
	jwtAuth *auth.JWTAuthenticator
}

func New(jwtAuth *auth.JWTAuthenticator) *Middleware {
	return &Middleware{jwtAuth: jwtAuth}
}
