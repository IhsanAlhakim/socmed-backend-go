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
	ErrInvalidToken         = errors.New("Invalid Token")
	ErrInvalidSigningMethod = errors.New("Invalid Signing Method")
)

var (
	LOGIN_EXPIRATION_DURATION = time.Duration(1) * time.Hour
	JWT_SIGNING_METHOD        = jwt.SigningMethodHS256
)

func GenerateToken(userId string, issuer string, JWTSignKey string) (string, error) {
	claims := MyClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   userId,
			Issuer:    issuer,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(LOGIN_EXPIRATION_DURATION)),
		},
	}

	token := jwt.NewWithClaims(
		JWT_SIGNING_METHOD,
		claims,
	)

	jwtSignKeyByte := []byte(JWTSignKey)

	signedToken, err := token.SignedString(jwtSignKeyByte)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func VerifyToken(tokenString string, JWTSignKey string) (jwt.Claims, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
		if method, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidSigningMethod
		} else if method != JWT_SIGNING_METHOD {
			return nil, ErrInvalidSigningMethod
		}

		jwtSignKeyByte := []byte(JWTSignKey)
		return jwtSignKeyByte, nil
	})

	if err != nil {
		return nil, ErrInvalidToken
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, ErrInvalidToken
	}
	return claims, nil
}
