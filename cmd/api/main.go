package main

import (
	"log"
)

func main() {
	app := newApp()

	mux := app.mount()

	if err := app.run(mux); err != nil {
		log.Fatal("Error starting server")
	}
}
