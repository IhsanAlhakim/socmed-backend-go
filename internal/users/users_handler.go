package users

import (
	"database/sql"
	"errors"
	"log"
	"net/http"

	"github.com/IhsanAlhakim/socmed-backend-go/internal/auth"
	"github.com/IhsanAlhakim/socmed-backend-go/internal/httpjson"
	"github.com/IhsanAlhakim/socmed-backend-go/internal/validation"
	"golang.org/x/crypto/bcrypt"
)

func NewHandler(service ServiceInterface) *Handler {
	return &Handler{
		service: service,
	}
}

type Handler struct {
	service ServiceInterface
}

func (h *Handler) GetUserById(w http.ResponseWriter, r *http.Request) {
	userId, err := auth.GetJWTSub(r)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user, err := h.service.GetUserById(int64(userId))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	httpjson.Respond(w, httpjson.ResponseBody{
		Data: user,
	}, http.StatusOK)
}

func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request) {
	var payload SignInParam
	if err := httpjson.Decode(r, &payload); err != nil {
		log.Println(err)
		if err == httpjson.ErrEmptyBody {
			http.Error(w, err.Error(), http.StatusBadRequest)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	cookie, err := h.service.SignIn(&payload)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "user not found", http.StatusNotFound)
		} else if err == bcrypt.ErrMismatchedHashAndPassword {
			http.Error(w, "invalid credentials", http.StatusUnauthorized)
		} else if errors.As(err, &validation.ErrValidation) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	http.SetCookie(w, cookie)
	httpjson.Respond(w, httpjson.ResponseBody{
		Message: "Sign In Successful",
	}, http.StatusOK)
}

func (h *Handler) SignOut(w http.ResponseWriter, r *http.Request) {

	cookie := h.service.SignOut()

	http.SetCookie(w, cookie)
	httpjson.Respond(w, httpjson.ResponseBody{
		Message: "Sign Out Successful",
	}, http.StatusOK)
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var payload CreateUserParam
	if err := httpjson.Decode(r, &payload); err != nil {
		if err == httpjson.ErrEmptyBody {
			http.Error(w, err.Error(), http.StatusBadRequest)
		} else {
			log.Println(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	err := h.service.CreateUser(&payload)
	if err != nil {
		if errors.As(err, &validation.ErrValidation) || err == ErrDuplicateEmail || err == ErrDuplicateUsername {
			http.Error(w, err.Error(), http.StatusBadRequest)
		} else {
			log.Println(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	httpjson.Respond(w, httpjson.ResponseBody{
		Message: "User Created!",
	}, http.StatusCreated)
}

func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	userId, err := auth.GetJWTSub(r)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var payload UpdateUserParam
	if err := httpjson.Decode(r, &payload); err != nil {
		log.Println(err)
		if err == httpjson.ErrEmptyBody {
			http.Error(w, err.Error(), http.StatusBadRequest)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	err = h.service.UpdateUser(int64(userId), &payload)
	if err != nil {
		log.Println(err)
		if errors.As(err, &validation.ErrValidation) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	httpjson.Respond(w, httpjson.ResponseBody{
		Message: "User Updated!",
	}, http.StatusOK)
}

func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	userId, err := auth.GetJWTSub(r)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = h.service.DeleteUser(int64(userId))
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	httpjson.Respond(w, httpjson.ResponseBody{
		Message: "User Deleted!",
	}, http.StatusOK)
}
