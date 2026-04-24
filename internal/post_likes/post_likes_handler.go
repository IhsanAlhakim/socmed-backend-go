package plikes

import (
	"log"
	"net/http"

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
	var postLikeData PostLikeParam
	if err := httpjson.Decode(r, &postLikeData); err != nil {
		log.Println(err)
		if err == httpjson.ErrEmptyBody {
			http.Error(w, err.Error(), http.StatusBadRequest)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	err := h.service.LikePost(&postLikeData)
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
	var postLikeData PostLikeParam
	if err := httpjson.Decode(r, &postLikeData); err != nil {
		log.Println(err)
		if err == httpjson.ErrEmptyBody {
			http.Error(w, err.Error(), http.StatusBadRequest)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	err := h.service.UnlikePost(&postLikeData)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	httpjson.Respond(w, httpjson.ResponseBody{
		Message: "Post Unliked...",
	}, http.StatusOK)
}
