package users

import (
	"net/http"
	"strconv"
	"time"

	"github.com/IhsanAlhakim/socmed-backend-go/internal/auth"
	"github.com/IhsanAlhakim/socmed-backend-go/internal/config"
)

func NewService(store StoreInterface, config *config.Config) ServiceInterface {
	return &Service{
		store:  store,
		config: config,
	}
}

type Service struct {
	store  StoreInterface
	config *config.Config
}

func (svc *Service) SignIn(payload *SignInParam) (*http.Cookie, error) {
	user, err := svc.store.GetUserByEmail(payload.Email)
	if err != nil {
		return nil, err
	}

	if err := auth.VerifyPassword(user.Password, payload.Password); err != nil {
		return nil, err
	}

	userIdString := strconv.Itoa(int(user.ID))

	token, err := auth.GenerateToken(userIdString, svc.config.AppName, svc.config.JWTSignKey)
	if err != nil {
		return nil, err
	}

	cookie := &http.Cookie{
		Name:     svc.config.TokenCookieName,
		Value:    token,
		Expires:  time.Now().Add(time.Duration(1) * time.Hour),
		HttpOnly: true,
		Path:     "/",
	}

	return cookie, nil
}

func (svc *Service) CreateUser(payload *CreateUserParam) error {
	// input validation

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
