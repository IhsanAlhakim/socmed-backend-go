package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/IhsanAlhakim/socmed-backend-go/internal/comments"
	"github.com/IhsanAlhakim/socmed-backend-go/internal/config"
	"github.com/IhsanAlhakim/socmed-backend-go/internal/env"
	"github.com/IhsanAlhakim/socmed-backend-go/internal/follows"
	"github.com/IhsanAlhakim/socmed-backend-go/internal/middlewares"
	plikes "github.com/IhsanAlhakim/socmed-backend-go/internal/post_likes"
	"github.com/IhsanAlhakim/socmed-backend-go/internal/posts"
	"github.com/IhsanAlhakim/socmed-backend-go/internal/users"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
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

	// TODO :   add graceful shutdown

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

	m := middlewares.New(app.config)

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{env.GetString("ALLOWED_ORIGIN", "http://localhost:5173")},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	// TODO : add CORS

	userStore := users.NewStore(app.db)
	userService := users.NewService(userStore, app.config)
	userHandler := users.NewHandler(userService)

	postStore := posts.NewStore(app.db)
	postService := posts.NewService(postStore)
	postHandler := posts.NewHandler(postService)

	followStore := follows.NewStore(app.db)
	followService := follows.NewService(followStore)
	followHandler := follows.NewHandler(followService)

	commentStore := comments.NewStore(app.db)
	commentService := comments.NewService(commentStore)
	commentHandler := comments.NewHandler(commentService)

	postLikesStore := plikes.NewStore(app.db)
	postLikesService := plikes.NewService(postLikesStore)
	postLikesHandler := plikes.NewHandler(postLikesService)

	r.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	r.HandleFunc("POST /users", userHandler.CreateUser)
	r.HandleFunc("POST /sessions", userHandler.SignIn)

	// TODO : add input validation

	// Endpoint with auth middleware
	r.Route("/", func(r chi.Router) {
		r.Use(m.Auth)

		r.HandleFunc("PUT /users/{id}", userHandler.UpdateUser)
		r.HandleFunc("DELETE /users/{id}", userHandler.DeleteUser)
		r.HandleFunc("DELETE /sessions", userHandler.SignOut)

		r.HandleFunc("GET /posts", postHandler.GetPosts)
		r.HandleFunc("GET /users/{userId}/followed-posts", postHandler.GetFollowedPosts)
		r.HandleFunc("GET /posts/{id}", postHandler.GetPostById)
		r.HandleFunc("POST /posts", postHandler.CreatePost)
		r.HandleFunc("DELETE /posts/{id}", postHandler.DeletePost)

		r.HandleFunc("GET /users/{userId}/followers", followHandler.GetFollower)
		r.HandleFunc("GET /users/{userId}/following", followHandler.GetFollowed)
		r.HandleFunc("POST /users/{userId}/follow", followHandler.Follow)
		r.HandleFunc("DELETE /users/{userId}/follow", followHandler.Unfollow)

		r.HandleFunc("POST /posts/{postId}/comments", commentHandler.CreateComment)
		r.HandleFunc("GET /posts/{postId}/comments", commentHandler.Getcomments)
		r.HandleFunc("DELETE /comments/{commentId}", commentHandler.DeleteComment)

		r.HandleFunc("POST /posts/{postId}/likes", postLikesHandler.LikePost)
		r.HandleFunc("DELETE /posts/{postId}/likes", postLikesHandler.UnlikePost)
		r.HandleFunc("GET /posts/{postId}/likes", postLikesHandler.GetPostLiker)
	})

	return r
}
