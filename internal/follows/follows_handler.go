package follows

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

func (h *Handler) GetFollower(w http.ResponseWriter, r *http.Request) {
	userId := r.PathValue("userId")
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	follower, err := h.service.GetFollower(int64(userIdInt))
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	httpjson.Respond(w, httpjson.ResponseBody{
		Data: follower,
	}, http.StatusOK)
}

func (h *Handler) GetFollowed(w http.ResponseWriter, r *http.Request) {
	userId := r.PathValue("userId")
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	followed, err := h.service.GetFollowed(int64(userIdInt))
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	httpjson.Respond(w, httpjson.ResponseBody{
		Data: followed,
	}, http.StatusOK)
}

func (h *Handler) Follow(w http.ResponseWriter, r *http.Request) {
	userId := r.PathValue("userId")
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var payload FollowParam
	if err := httpjson.Decode(r, &payload); err != nil {
		log.Println(err)
		if err == httpjson.ErrEmptyBody {
			http.Error(w, err.Error(), http.StatusBadRequest)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	err = h.service.Follow(int64(userIdInt), &payload)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	httpjson.Respond(w, httpjson.ResponseBody{
		Message: "Follow Data Created!",
	}, http.StatusCreated)
}

func (h *Handler) Unfollow(w http.ResponseWriter, r *http.Request) {
	userId := r.PathValue("userId")
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var payload FollowParam
	if err := httpjson.Decode(r, &payload); err != nil {
		log.Println(err)
		if err == httpjson.ErrEmptyBody {
			http.Error(w, err.Error(), http.StatusBadRequest)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	err = h.service.Unfollow(int64(userIdInt), &payload)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	httpjson.Respond(w, httpjson.ResponseBody{
		Message: "Follow Data Deleted!",
	}, http.StatusOK)
}
