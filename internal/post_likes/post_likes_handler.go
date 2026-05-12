package plikes

import (
	"log"
	"net/http"
	"strconv"

	"github.com/IhsanAlhakim/socmed-backend-go/internal/auth"
	"github.com/IhsanAlhakim/socmed-backend-go/internal/httpjson"
	"github.com/IhsanAlhakim/socmed-backend-go/internal/posts"
	"github.com/IhsanAlhakim/socmed-backend-go/internal/users"
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
	postId := r.PathValue("postId")
	postIdInt, err := strconv.Atoi(postId)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userId, err := auth.GetJWTSub(r)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = h.service.LikePost(int64(postIdInt), int64(userId))
	if err != nil {
		switch {
		case err == ErrPostAlreadyLiked:
			http.Error(w, err.Error(), http.StatusConflict)
		case err == users.ErrUserNotFound || err == posts.ErrPostNotFound:
			http.Error(w, err.Error(), http.StatusNotFound)
		default:
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	httpjson.Respond(w, httpjson.ResponseBody{
		Message: "Post Liked!",
	}, http.StatusCreated)
}

func (h *Handler) UnlikePost(w http.ResponseWriter, r *http.Request) {
	postId := r.PathValue("postId")
	postIdInt, err := strconv.Atoi(postId)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userId, err := auth.GetJWTSub(r)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = h.service.UnlikePost(int64(postIdInt), int64(userId))
	if err != nil {
		if err == ErrLikeNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
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

func (h *Handler) GetPostLikesCount(w http.ResponseWriter, r *http.Request) {
	postId := r.PathValue("postId")
	postIdInt, err := strconv.Atoi(postId)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	postLikesCount, err := h.service.GetPostLikesCount(int64(postIdInt))
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	httpjson.Respond(w, httpjson.ResponseBody{
		Data: postLikesCount,
	}, http.StatusOK)
}
