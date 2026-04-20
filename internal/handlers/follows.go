package handlers

import (
	"log"
	"net/http"
	"strconv"

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

func (h *FollowHandler) GetFollower(w http.ResponseWriter, r *http.Request) {
	followedId := r.PathValue("followerId")
	followedIdInt, err := strconv.Atoi(followedId)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	follower, err := h.service.GetFollower(int64(followedIdInt))
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	httpjson.Respond(w, httpjson.ResponseBody{
		Data: follower,
	}, http.StatusCreated)
}

func (h *FollowHandler) Follow(w http.ResponseWriter, r *http.Request) {
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

	err := h.service.Follow(&followData)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	httpjson.Respond(w, httpjson.ResponseBody{
		Message: "Follow Data Created!",
	}, http.StatusCreated)
}

func (h *FollowHandler) Unfollow(w http.ResponseWriter, r *http.Request) {
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

	err := h.service.Unfollow(&followData)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	httpjson.Respond(w, httpjson.ResponseBody{
		Message: "Follow Data Deleted!",
	}, http.StatusCreated)
}
