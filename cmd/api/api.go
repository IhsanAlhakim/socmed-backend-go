package main

import (
	"log"
	"net/http"
	"time"
)

type application struct {
}

func newApp() *application {
	return &application{}
}

func (app *application) mount() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	return mux
}

func (app *application) run(mux http.Handler) error {
	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 30,
		IdleTimeout:  time.Minute,
	}

	log.Println("Server has started at :8080")
	return server.ListenAndServe()
}
