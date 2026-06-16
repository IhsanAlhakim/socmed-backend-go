package posts

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

func (h *Handler) GetPosts(w http.ResponseWriter, r *http.Request) {

	userId, err := auth.GetJWTSub(r)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	posts, err := h.service.GetPosts(int64(userId))
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	httpjson.Respond(w, httpjson.ResponseBody{
		Data: posts,
	}, http.StatusOK)
}

func (h *Handler) GetPostsByUsername(w http.ResponseWriter, r *http.Request) {
	userId, err := auth.GetJWTSub(r)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	otherUsername := r.PathValue("username")

	posts, err := h.service.GetPostsByUsername(int64(userId), otherUsername)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	httpjson.Respond(w, httpjson.ResponseBody{
		Data: posts,
	}, http.StatusOK)
}

func (h *Handler) GetFollowedPosts(w http.ResponseWriter, r *http.Request) {
	userId, err := auth.GetJWTSub(r)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	posts, err := h.service.GetFollowedPosts(int64(userId))
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	httpjson.Respond(w, httpjson.ResponseBody{
		Data: posts,
	}, http.StatusOK)
}

func (h *Handler) GetLikedPosts(w http.ResponseWriter, r *http.Request) {
	userId, err := auth.GetJWTSub(r)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	posts, err := h.service.GetLikedPosts(int64(userId))
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	httpjson.Respond(w, httpjson.ResponseBody{
		Data: posts,
	}, http.StatusOK)
}

func (h *Handler) GetPostById(w http.ResponseWriter, r *http.Request) {
	postId := r.PathValue("id")

	postIdInt, err := strconv.Atoi(postId)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	post, err := h.service.GetPostById(int64(postIdInt))
	if err != nil {
		if err == ErrPostNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	httpjson.Respond(w, httpjson.ResponseBody{
		Data: post,
	}, http.StatusOK)
}

func (h *Handler) CreatePost(w http.ResponseWriter, r *http.Request) {
	userId, err := auth.GetJWTSub(r)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var payload CreatePostParam
	if err := httpjson.Decode(r, &payload); err != nil {
		if err == httpjson.ErrEmptyBody {
			http.Error(w, err.Error(), http.StatusBadRequest)
		} else {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	err = h.service.CreatePost(int64(userId), &payload)
	if err != nil {
		switch {
		case err == users.ErrUserNotFound:
			http.Error(w, err.Error(), http.StatusNotFound)
		case validation.IsErrValidation(err):
			http.Error(w, err.Error(), http.StatusBadRequest)
		default:
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	httpjson.Respond(w, httpjson.ResponseBody{
		Message: "Post Created!",
	}, http.StatusCreated)
}

func (h *Handler) DeletePost(w http.ResponseWriter, r *http.Request) {
	postId := r.PathValue("id")
	postIdInt, err := strconv.Atoi(postId)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = h.service.DeletePost(int64(postIdInt))
	if err != nil {
		if err == ErrPostNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	httpjson.Respond(w, httpjson.ResponseBody{
		Message: "Post Deleted!",
	}, http.StatusOK)
}
