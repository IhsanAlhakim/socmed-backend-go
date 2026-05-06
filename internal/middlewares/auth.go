package middlewares

import (
	"context"
	"errors"
	"net/http"

	"github.com/IhsanAlhakim/socmed-backend-go/internal/auth"
)

func (m *Middleware) Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			storedToken, err := r.Cookie(m.jwtAuth.TokenCookieName)
			if err != nil {
				http.Error(w, "Invalid credentials", http.StatusUnauthorized)
				return
			}
			claims, err := m.jwtAuth.VerifyToken(storedToken.Value)
			if err != nil {
				if errors.Is(err, auth.ErrInvalidSigningMethod) {
					http.Error(w, "Invalid credentials", http.StatusUnauthorized)
					return
				}
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			ctx := context.WithValue(r.Context(), m.jwtAuth.ContextKey, claims)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
}
