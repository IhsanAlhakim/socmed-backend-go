package handlers

import (
	"log"
	"net/http"

	"github.com/IhsanAlhakim/socmed-backend-go/internal/httpjson"
	"github.com/IhsanAlhakim/socmed-backend-go/internal/services"
	"github.com/IhsanAlhakim/socmed-backend-go/internal/store"
)

type FollowHandler struct {
	service services.FollowServiceInterface
}

func NewFollowHandler(service services.FollowServiceInterface) *FollowHandler {
	return &FollowHandler{
		service: service,
	}
}

func (h *FollowHandler) CreateFollow(w http.ResponseWriter, r *http.Request) {
	var followData store.Follow
	if err := httpjson.Decode(r, &followData); err != nil {
		log.Println(err)
		if err == httpjson.ErrEmptyBody {
			http.Error(w, err.Error(), http.StatusBadRequest)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	err := h.service.CreateFollow(&followData)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	httpjson.Respond(w, httpjson.ResponseBody{
		Message: "Follow Data Created!",
	}, http.StatusCreated)
}
