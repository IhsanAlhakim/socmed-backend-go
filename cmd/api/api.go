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
