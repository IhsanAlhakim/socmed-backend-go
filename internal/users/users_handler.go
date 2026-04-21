package users

import (
	"log"
	"net/http"
	"strconv"

	"github.com/IhsanAlhakim/socmed-backend-go/internal/httpjson"
)

func NewHandler(service ServiceInterface) *Handler {
	return &Handler{
		service: service,
	}
}

type Handler struct {
	service ServiceInterface
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := httpjson.Decode(r, &user); err != nil {
		log.Println(err)
		if err == httpjson.ErrEmptyBody {
			http.Error(w, err.Error(), http.StatusBadRequest)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	err := h.service.CreateUser(&user)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	httpjson.Respond(w, httpjson.ResponseBody{
		Message: "User Created!",
	}, http.StatusCreated)
}

func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	userId := r.PathValue("id")
	userIdInt, err := strconv.Atoi(userId)

	var updatedUser User
	if err := httpjson.Decode(r, &updatedUser); err != nil {
		log.Println(err)
		if err == httpjson.ErrEmptyBody {
			http.Error(w, err.Error(), http.StatusBadRequest)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	err = h.service.UpdateUser(int64(userIdInt), &updatedUser)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	httpjson.Respond(w, httpjson.ResponseBody{
		Message: "User Updated!",
	}, http.StatusOK)
}

func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	userId := r.PathValue("id")
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = h.service.DeleteUser(int64(userIdInt))
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	httpjson.Respond(w, httpjson.ResponseBody{
		Message: "User Deleted!",
	}, http.StatusOK)
}
