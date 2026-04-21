package users

import (
	"time"
)

// Interface
type StoreInterface interface {
	Create(userData *User) error
	Update(userId int64, userData *User) error
	Delete(userId int64) error
}

type ServiceInterface interface {
	CreateUser(userData *User) error
	UpdateUser(userId int64, userData *User) error
	DeleteUser(userId int64) error
}

// Struct
type User struct {
	ID       int64     `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	CreateAt time.Time `json:"created_at"`
}
