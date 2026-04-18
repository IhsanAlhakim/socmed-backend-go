package services

import (
	"github.com/IhsanAlhakim/socmed-backend-go/internal/store"
)

type UserServiceInterface interface {
	CreateUser(userData store.User) error
	UpdateUser(userData store.User) error
	DeleteUser(userId int64)
}

type UserService struct {
	storage store.Storage
}

func NewUserService(storage store.Storage) *UserService {
	return &UserService{
		storage: storage,
	}
}

func (usvc *UserService) CreateUser(userdata store.User) error {
	// input validation

	// hash password

	err := usvc.storage.Users.Create(&userdata)
	if err != nil {
		return err
	}

	return nil
}

func (usvc *UserService) UpdateUser(userdata store.User) error {
	err := usvc.storage.Users.Update(&userdata)
	if err != nil {
		return err
	}
	return nil
}

func (usvc *UserService) DeleteUser(userId int64) error {

	err := usvc.storage.Users.Delete(userId)
	if err != nil {
		return err
	}
	return nil
}
