package plikes

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

func (h *Handler) LikePost(w http.ResponseWriter, r *http.Request) {
	var payload PostLikeParam
	if err := httpjson.Decode(r, &payload); err != nil {
		log.Println(err)
		if err == httpjson.ErrEmptyBody {
			http.Error(w, err.Error(), http.StatusBadRequest)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	err := h.service.LikePost(&payload)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	httpjson.Respond(w, httpjson.ResponseBody{
		Message: "Post Liked!",
	}, http.StatusCreated)
}

func (h *Handler) UnlikePost(w http.ResponseWriter, r *http.Request) {
	var payload PostLikeParam
	if err := httpjson.Decode(r, &payload); err != nil {
		log.Println(err)
		if err == httpjson.ErrEmptyBody {
			http.Error(w, err.Error(), http.StatusBadRequest)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	err := h.service.UnlikePost(&payload)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	httpjson.Respond(w, httpjson.ResponseBody{
		Message: "Post Unliked...",
	}, http.StatusOK)
}

func (h *Handler) GetPostLiker(w http.ResponseWriter, r *http.Request) {
	postId := r.PathValue("postId")
	postIdInt, err := strconv.Atoi(postId)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	follower, err := h.service.GetPostLiker(int64(postIdInt))
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	httpjson.Respond(w, httpjson.ResponseBody{
		Data: follower,
	}, http.StatusOK)
}
