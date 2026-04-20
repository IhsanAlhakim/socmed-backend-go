package handlers

import (
	"log"
	"net/http"

	"github.com/IhsanAlhakim/socmed-backend-go/internal/httpjson"
	"github.com/IhsanAlhakim/socmed-backend-go/internal/services"
	"github.com/IhsanAlhakim/socmed-backend-go/internal/store"
)

type PostHandler struct {
	service services.PostService
}

func NewPostHandler(service services.PostService) *PostHandler {
	return &PostHandler{
		service: service,
	}
}

func (h *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	var post store.Post
	if err := httpjson.Decode(r, &post); err != nil {
		log.Println(err)
		if err == httpjson.ErrEmptyBody {
			http.Error(w, err.Error(), http.StatusBadRequest)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	err := h.service.CreatePost(post)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	httpjson.Respond(w, httpjson.ResponseBody{
		Message: "Post Created!",
	}, http.StatusCreated)
}
