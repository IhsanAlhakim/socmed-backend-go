package comments

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

func (h *Handler) CreateComment(w http.ResponseWriter, r *http.Request) {
	var commentData CreateCommentParam
	if err := httpjson.Decode(r, &commentData); err != nil {
		log.Println(err)
		if err == httpjson.ErrEmptyBody {
			http.Error(w, err.Error(), http.StatusBadRequest)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	err := h.service.CreateComment(&commentData)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	httpjson.Respond(w, httpjson.ResponseBody{
		Message: "Comment Created!",
	}, http.StatusCreated)
}
