package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/IhsanAlhakim/socmed-backend-go/internal/auth"
	"github.com/IhsanAlhakim/socmed-backend-go/internal/comments"
	"github.com/IhsanAlhakim/socmed-backend-go/internal/config"
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
	db      *sql.DB
	config  *config.Config
	jwtAuth *auth.JWTAuthenticator
}

func newApp(db *sql.DB, config *config.Config, jwtAuth *auth.JWTAuthenticator) *application {
	return &application{
		db:      db,
		config:  config,
		jwtAuth: jwtAuth,
	}
}

func (app *application) run(mux http.Handler) error {
	server := &http.Server{
		Addr:         ":" + app.config.Port,
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 30,
		IdleTimeout:  time.Minute,
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	var errChan = make(chan error)

	// Run server in the background
	go func() {
		log.Printf("Server has started at port %s", app.config.Port)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errChan <- err
		}
	}()

	select {
	case err := <-errChan:
		return err
	case <-ctx.Done(): // Listen for the interrupt signal
		// Create shutdown context with 30-second timeout
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		// Trigger graceful shutdown
		fmt.Println("graceful shutdown...")
		if err := server.Shutdown(shutdownCtx); err != nil {
			return err
		}
	}
	return nil
}

func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	m := middlewares.New(app.jwtAuth)

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{app.config.AllowedOrigin},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	userStore := users.NewStore(app.db)
	userService := users.NewService(userStore, app.jwtAuth)
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

	// Endpoint with auth middleware
	r.Route("/", func(r chi.Router) {
		r.Use(m.Auth)

		r.HandleFunc("GET /users", userHandler.GetUserById)
		r.HandleFunc("GET /users/{username}", userHandler.GetUserByUsername)
		r.HandleFunc("PUT /users", userHandler.UpdateUser)
		r.HandleFunc("DELETE /users", userHandler.DeleteUser)
		r.HandleFunc("DELETE /sessions", userHandler.SignOut)

		r.HandleFunc("GET /posts", postHandler.GetPosts)
		r.HandleFunc("GET /users/{username}/posts", postHandler.GetPostsByUsername)
		r.HandleFunc("GET /users/followed-posts", postHandler.GetFollowedPosts)
		r.HandleFunc("GET /users/liked-posts", postHandler.GetLikedPosts)
		r.HandleFunc("GET /posts/{id}", postHandler.GetPostById)
		r.HandleFunc("POST /posts", postHandler.CreatePost)
		r.HandleFunc("DELETE /posts/{id}", postHandler.DeletePost)

		r.HandleFunc("GET /users/{userId}/followers", followHandler.GetFollower)
		r.HandleFunc("GET /users/{userId}/following", followHandler.GetFollowed)
		r.HandleFunc("POST /follow", followHandler.Follow)
		r.HandleFunc("DELETE /follow", followHandler.Unfollow)

		r.HandleFunc("POST /posts/{postId}/comments", commentHandler.CreateComment)
		r.HandleFunc("GET /posts/{postId}/comments", commentHandler.Getcomments)
		r.HandleFunc("DELETE /comments/{commentId}", commentHandler.DeleteComment)

		r.HandleFunc("POST /posts/{postId}/likes", postLikesHandler.LikePost)
		r.HandleFunc("DELETE /posts/{postId}/likes", postLikesHandler.UnlikePost)
		r.HandleFunc("GET /posts/{postId}/likes", postLikesHandler.GetPostLiker)
		r.HandleFunc("GET /posts/{postId}/likes/count", postLikesHandler.GetPostLikesCount)
	})

	return r
}
