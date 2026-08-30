package main

import (
	"log"
	"net/http"

	"github.com/jevitapearl/TaskForge/internal/config"
	"github.com/jevitapearl/TaskForge/internal/database"
	"github.com/jevitapearl/TaskForge/internal/middleware"
	"github.com/jevitapearl/TaskForge/internal/repository"
	"github.com/jevitapearl/TaskForge/internal/server"
)

func main() {
	cfg, err := config.LoadDBconfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.NewPostgres(cfg)
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewPostgresRepository(db)
	s := server.New(repo)

	wrapped := middleware.LoggingMiddleware(s.Router())

	log.Println("Server running on port 8080")
	err = http.ListenAndServe(":8080", wrapped)
	if err != nil {
		panic(err)
	}

}
