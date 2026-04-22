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
	followedId := r.PathValue("followedId")
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
	}, http.StatusOK)
}

func (h *Handler) GetFollowed(w http.ResponseWriter, r *http.Request) {
	followerId := r.PathValue("followerId")
	followerIdInt, err := strconv.Atoi(followerId)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	followed, err := h.service.GetFollowed(int64(followerIdInt))
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
	var followData FollowDataparam
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

func (h *Handler) Unfollow(w http.ResponseWriter, r *http.Request) {
	var followData FollowDataparam
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
	}, http.StatusOK)
}
