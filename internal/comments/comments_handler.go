package comments

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

func (h *Handler) CreateComment(w http.ResponseWriter, r *http.Request) {
	var payload CreateCommentParam
	if err := httpjson.Decode(r, &payload); err != nil {
		log.Println(err)
		if err == httpjson.ErrEmptyBody {
			http.Error(w, err.Error(), http.StatusBadRequest)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	err := h.service.CreateComment(&payload)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	httpjson.Respond(w, httpjson.ResponseBody{
		Message: "Comment Created!",
	}, http.StatusCreated)
}

func (h *Handler) Getcomments(w http.ResponseWriter, r *http.Request) {
	postId := r.PathValue("postId")

	postIdInt, err := strconv.Atoi(postId)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	comments, err := h.service.Getcomments(int64(postIdInt))
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	httpjson.Respond(w, httpjson.ResponseBody{
		Data: comments,
	}, http.StatusOK)
}

func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	commentId := r.PathValue("commentId")

	commentIdInt, err := strconv.Atoi(commentId)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = h.service.DeleteComment(int64(commentIdInt))
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	httpjson.Respond(w, httpjson.ResponseBody{
		Message: "Comment Deleted",
	}, http.StatusOK)
}
