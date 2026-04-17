package main

import (
	"log"
	"net/http"
	"time"

	"github.com/IhsanAlhakim/socmed-backend-go/internal/config"
	"github.com/IhsanAlhakim/socmed-backend-go/internal/handlers"
	"github.com/IhsanAlhakim/socmed-backend-go/internal/services"
	"github.com/IhsanAlhakim/socmed-backend-go/internal/store"
)

type application struct {
	store  store.Storage
	config config.Config
}

func newApp(storage store.Storage, config config.Config) *application {
	return &application{
		store:  storage,
		config: config,
	}
}

func (app *application) mount() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	userService := services.NewUserService(app.store)
	userHandler := handlers.NewUserHandler(*userService)
	mux.HandleFunc("POST /users", userHandler.CreateUser)

	return mux
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
