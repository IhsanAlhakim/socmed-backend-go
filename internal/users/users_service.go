package users

import (
	"net/http"
	"strconv"
	"time"

	"github.com/IhsanAlhakim/socmed-backend-go/internal/auth"
	"github.com/IhsanAlhakim/socmed-backend-go/internal/validation"
)

func NewService(store StoreInterface, jwtAuth *auth.JWTAuthenticator) ServiceInterface {
	return &Service{
		store:   store,
		jwtAuth: jwtAuth,
	}
}

type Service struct {
	store   StoreInterface
	jwtAuth *auth.JWTAuthenticator
}

func (svc *Service) GetUserById(userId int64) (*User, error) {
	user, err := svc.store.GetUserById(userId)
	if err != nil {
		return &User{}, err
	}
	return user, nil
}

func (svc *Service) GetUserByUsername(username string) (*User, error) {
	user, err := svc.store.GetUserByUsername(username)
	if err != nil {
		return &User{}, err
	}
	return user, nil
}

func (svc *Service) SignIn(payload *SignInParam) (*http.Cookie, error) {
	if err := validation.Validate.Struct(payload); err != nil {
		return nil, validation.NewError(err)
	}

	user, err := svc.store.GetUserByEmail(payload.Email)
	if err != nil {
		return nil, err
	}

	if err := auth.VerifyPassword(user.Password, payload.Password); err != nil {
		return nil, err
	}

	userIdString := strconv.Itoa(int(user.ID))

	token, err := svc.jwtAuth.GenerateToken(userIdString)
	if err != nil {
		return nil, err
	}

	cookie := &http.Cookie{
		Name:     svc.jwtAuth.TokenCookieName,
		Value:    token,
		Expires:  time.Now().Add(time.Duration(1) * time.Hour),
		HttpOnly: true,
		Path:     "/",
	}

	return cookie, nil
}

func (svc *Service) SignOut() *http.Cookie {
	cookie := &http.Cookie{
		Name:    svc.jwtAuth.TokenCookieName,
		Expires: time.Unix(0, 0),
		MaxAge:  -1,
	}
	return cookie
}

func (svc *Service) CreateUser(payload *CreateUserParam) error {
	if err := validation.Validate.Struct(payload); err != nil {
		return validation.NewError(err)
	}

	// hash password
	hashedPassword, err := auth.GenerateHashPassword(payload.Password)
	if err != nil {
		return err
	}
	payload.Password = string(hashedPassword)

	err = svc.store.CreateUser(payload)
	if err != nil {
		return err
	}

	return nil
}

func (svc *Service) UpdateUser(userId int64, payload *UpdateUserParam) error {
	if err := validation.Validate.Struct(payload); err != nil {
		return validation.NewError(err)
	}

	err := svc.store.UpdateUser(userId, payload)
	if err != nil {
		return err
	}
	return nil
}

func (svc *Service) DeleteUser(userId int64) error {

	err := svc.store.DeleteUser(userId)
	if err != nil {
		return err
	}
	return nil
}
