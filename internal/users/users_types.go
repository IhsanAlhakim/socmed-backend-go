package users

import (
	"net/http"
	"time"
)

// Interface
type StoreInterface interface {
	CreateUser(payload *CreateUserParam) error
	UpdateUser(userId int64, payload *UpdateUserParam) error
	DeleteUser(userId int64) error
	GetUserByEmail(email string) (*User, error)
}

type ServiceInterface interface {
	CreateUser(payload *CreateUserParam) error
	UpdateUser(userId int64, payload *UpdateUserParam) error
	DeleteUser(userId int64) error
	SignIn(payload *SignInParam) (*http.Cookie, error)
	SignOut() *http.Cookie
}

// Struct
type User struct {
	ID       int64     `json:"id,omitempty"`
	Username string    `json:"username,omitempty"`
	Email    string    `json:"email,omitempty"`
	Password string    `json:"password,omitempty"`
	CreateAt time.Time `json:"created_at,omitempty"`
}

type CreateUserParam struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUserParam struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type SignInParam struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
