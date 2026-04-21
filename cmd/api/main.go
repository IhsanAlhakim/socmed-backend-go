package main

import (
	"log"

	"github.com/IhsanAlhakim/socmed-backend-go/internal/config"
	"github.com/IhsanAlhakim/socmed-backend-go/internal/database"
)

func main() {
	cfg := config.Load()

	db, err := database.New(cfg.DBConfig)
	if err != nil {
		log.Fatal("Failed connecting to database: ", err)
	}
	defer db.Close()

	app := newApp(db, cfg)

	mux := app.mount()

	if err := app.run(mux); err != nil {
		log.Fatal("Error starting server")
	}
}
