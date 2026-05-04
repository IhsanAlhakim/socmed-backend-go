package middlewares

import (
	"context"
	"errors"
	"net/http"

	"github.com/IhsanAlhakim/socmed-backend-go/internal/auth"
)

var ContextWithUserInfoKey = "userInfo"

func (m *Middleware) Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			storedToken, err := r.Cookie(m.config.TokenCookieName)
			if err != nil {
				http.Error(w, "Invalid credentials", http.StatusUnauthorized)
				return
			}
			claims, err := auth.VerifyToken(storedToken.Value, m.config.JWTSignKey)
			if err != nil {
				if errors.Is(err, auth.ErrInvalidSigningMethod) {
					http.Error(w, "Invalid credentials", http.StatusUnauthorized)
					return
				}
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			ctx := context.WithValue(r.Context(), ContextWithUserInfoKey, claims)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
}
