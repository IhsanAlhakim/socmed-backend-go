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
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	userStore := users.NewStore(app.db)
	userService := users.NewService(userStore, app.config)
	userHandler := users.NewHandler(userService)
	mux.HandleFunc("POST /users", userHandler.CreateUser)
	mux.HandleFunc("PUT /users/{id}", userHandler.UpdateUser)
	mux.HandleFunc("DELETE /users/{id}", userHandler.DeleteUser)
	mux.HandleFunc("POST /sessions", userHandler.SignIn)
	mux.HandleFunc("DELETE /sessions", userHandler.SignOut)

	postStore := posts.NewStore(app.db)
	postService := posts.NewService(postStore)
	postHandler := posts.NewHandler(postService)
	mux.HandleFunc("GET /posts", postHandler.GetPosts)
	mux.HandleFunc("GET /users/{userId}/followed-posts", postHandler.GetFollowedPosts)
	mux.HandleFunc("GET /posts/{id}", postHandler.GetPostById)
	mux.HandleFunc("POST /posts", postHandler.CreatePost)
	mux.HandleFunc("DELETE /posts/{id}", postHandler.DeletePost)

	followStore := follows.NewStore(app.db)
	followService := follows.NewService(followStore)
	followHandler := follows.NewHandler(followService)
	mux.HandleFunc("GET /users/{userId}/followers", followHandler.GetFollower)
	mux.HandleFunc("GET /users/{userId}/following", followHandler.GetFollowed)
	mux.HandleFunc("POST /users/{userId}/follow", followHandler.Follow)
	mux.HandleFunc("DELETE /users/{userId}/follow", followHandler.Unfollow)

	commentStore := comments.NewStore(app.db)
	commentService := comments.NewService(commentStore)
	commentHandler := comments.NewHandler(commentService)
	mux.HandleFunc("POST /posts/{postId}/comments", commentHandler.CreateComment)
	mux.HandleFunc("GET /posts/{postId}/comments", commentHandler.Getcomments)
	mux.HandleFunc("DELETE /comments/{commentId}", commentHandler.DeleteComment)

	postLikesStore := plikes.NewStore(app.db)
	postLikesService := plikes.NewService(postLikesStore)
	postLikesHandler := plikes.NewHandler(postLikesService)
	mux.HandleFunc("POST /posts/{postId}/likes", postLikesHandler.LikePost)
	mux.HandleFunc("DELETE /posts/{postId}/likes", postLikesHandler.UnlikePost)
	mux.HandleFunc("GET /posts/{postId}/likes", postLikesHandler.GetPostLiker)
	return mux
}
