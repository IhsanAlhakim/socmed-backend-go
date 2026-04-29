package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/IhsanAlhakim/socmed-backend-go/internal/comments"
	"github.com/IhsanAlhakim/socmed-backend-go/internal/config"
	"github.com/IhsanAlhakim/socmed-backend-go/internal/follows"
	plikes "github.com/IhsanAlhakim/socmed-backend-go/internal/post_likes"
	"github.com/IhsanAlhakim/socmed-backend-go/internal/posts"
	"github.com/IhsanAlhakim/socmed-backend-go/internal/users"
	"github.com/go-chi/chi/v5"
)

type application struct {
	db     *sql.DB
	config *config.Config
}

func (app *application) run(mux http.Handler) error {
	server := &http.Server{
		Addr:         ":" + app.config.Port,
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 30,
		IdleTimeout:  time.Minute,
	}

	log.Println("Server has started at :8080")
	return server.ListenAndServe()
}

func newApp(db *sql.DB, config *config.Config) *application {
	return &application{
		db:     db,
		config: config,
	}
}

func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	r.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	userStore := users.NewStore(app.db)
	userService := users.NewService(userStore, app.config)
	userHandler := users.NewHandler(userService)
	r.HandleFunc("POST /users", userHandler.CreateUser)
	r.HandleFunc("PUT /users/{id}", userHandler.UpdateUser)
	r.HandleFunc("DELETE /users/{id}", userHandler.DeleteUser)
	r.HandleFunc("POST /sessions", userHandler.SignIn)
	r.HandleFunc("DELETE /sessions", userHandler.SignOut)

	postStore := posts.NewStore(app.db)
	postService := posts.NewService(postStore)
	postHandler := posts.NewHandler(postService)
	r.HandleFunc("GET /posts", postHandler.GetPosts)
	r.HandleFunc("GET /users/{userId}/followed-posts", postHandler.GetFollowedPosts)
	r.HandleFunc("GET /posts/{id}", postHandler.GetPostById)
	r.HandleFunc("POST /posts", postHandler.CreatePost)
	r.HandleFunc("DELETE /posts/{id}", postHandler.DeletePost)

	followStore := follows.NewStore(app.db)
	followService := follows.NewService(followStore)
	followHandler := follows.NewHandler(followService)
	r.HandleFunc("GET /users/{userId}/followers", followHandler.GetFollower)
	r.HandleFunc("GET /users/{userId}/following", followHandler.GetFollowed)
	r.HandleFunc("POST /users/{userId}/follow", followHandler.Follow)
	r.HandleFunc("DELETE /users/{userId}/follow", followHandler.Unfollow)

	commentStore := comments.NewStore(app.db)
	commentService := comments.NewService(commentStore)
	commentHandler := comments.NewHandler(commentService)
	r.HandleFunc("POST /posts/{postId}/comments", commentHandler.CreateComment)
	r.HandleFunc("GET /posts/{postId}/comments", commentHandler.Getcomments)
	r.HandleFunc("DELETE /comments/{commentId}", commentHandler.DeleteComment)

	postLikesStore := plikes.NewStore(app.db)
	postLikesService := plikes.NewService(postLikesStore)
	postLikesHandler := plikes.NewHandler(postLikesService)
	r.HandleFunc("POST /posts/{postId}/likes", postLikesHandler.LikePost)
	r.HandleFunc("DELETE /posts/{postId}/likes", postLikesHandler.UnlikePost)
	r.HandleFunc("GET /posts/{postId}/likes", postLikesHandler.GetPostLiker)
	return r
}
