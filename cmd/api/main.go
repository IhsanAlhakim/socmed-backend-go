package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	app := newApp()

	if err := app.run(mux); err != nil {
		log.Fatal("Error starting server")
	}
}
