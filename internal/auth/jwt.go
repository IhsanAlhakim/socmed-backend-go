package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type MyClaims struct {
	jwt.RegisteredClaims
}

var (
	ErrInvalidSigningMethod = errors.New("Invalid Signing Method")
)

var (
	LOGIN_EXPIRATION_DURATION = time.Duration(1) * time.Hour
	JWT_SIGNING_METHOD        = jwt.SigningMethodHS256
)

type JWTAuthenticator struct {
	issuer          string
	signKey         string
	TokenCookieName string
	ContextKey      string
}

func NewJWTAuthenticator(issuer string, signKey string, contextKey string, tokenCookieName string) *JWTAuthenticator {
	return &JWTAuthenticator{
		issuer:          issuer,
		signKey:         signKey,
		ContextKey:      contextKey,
		TokenCookieName: tokenCookieName,
	}
}

func (ja *JWTAuthenticator) GenerateToken(userId string) (string, error) {
	claims := MyClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   userId,
			Issuer:    ja.issuer,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(LOGIN_EXPIRATION_DURATION)),
		},
	}

	token := jwt.NewWithClaims(
		JWT_SIGNING_METHOD,
		claims,
	)

	signKeyByte := []byte(ja.signKey)

	signedToken, err := token.SignedString(signKeyByte)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func (ja *JWTAuthenticator) VerifyToken(tokenString string) (jwt.Claims, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
		if method, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidSigningMethod
		} else if method != JWT_SIGNING_METHOD {
			return nil, ErrInvalidSigningMethod
		}

		signKeyByte := []byte(ja.signKey)
		return signKeyByte, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, err
	}
	return claims, nil
}
