package follows

import (
	"log"
	"net/http"
	"strconv"

	"github.com/IhsanAlhakim/socmed-backend-go/internal/auth"
	"github.com/IhsanAlhakim/socmed-backend-go/internal/httpjson"
	"github.com/IhsanAlhakim/socmed-backend-go/internal/users"
	"github.com/IhsanAlhakim/socmed-backend-go/internal/validation"
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
	userId, err := auth.GetJWTSub(r)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var payload FollowParam
	if err := httpjson.Decode(r, &payload); err != nil {
		if err == httpjson.ErrEmptyBody {
			http.Error(w, err.Error(), http.StatusBadRequest)
		} else {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	err = h.service.Follow(int64(userId), &payload)
	if err != nil {
		switch {
		case validation.IsErrValidation(err):
			http.Error(w, err.Error(), http.StatusBadRequest)
		case err == ErrFollowedNotFound || err == users.ErrUserNotFound:
			http.Error(w, err.Error(), http.StatusNotFound)
		case err == ErrUserAlreadyFollowed:
			http.Error(w, err.Error(), http.StatusConflict)
		default:
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	httpjson.Respond(w, httpjson.ResponseBody{
		Message: "Follow Data Created!",
	}, http.StatusCreated)
}

func (h *Handler) Unfollow(w http.ResponseWriter, r *http.Request) {
	userId, err := auth.GetJWTSub(r)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var payload FollowParam
	if err := httpjson.Decode(r, &payload); err != nil {
		if err == httpjson.ErrEmptyBody {
			http.Error(w, err.Error(), http.StatusBadRequest)
		} else {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	err = h.service.Unfollow(int64(userId), &payload)
	if err != nil {
		switch {
		case validation.IsErrValidation(err):
			http.Error(w, err.Error(), http.StatusBadRequest)
		case err == users.ErrUserNotFound:
			http.Error(w, err.Error(), http.StatusNotFound)
		default:
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	httpjson.Respond(w, httpjson.ResponseBody{
		Message: "Follow Data Deleted!",
	}, http.StatusOK)
}
