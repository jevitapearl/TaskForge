package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/golang-migrate/migrate/v4"
	"github.com/jevitapearl/TaskForge/internal/config"
	"github.com/jevitapearl/TaskForge/internal/database"
	"github.com/jevitapearl/TaskForge/internal/middleware"
	"github.com/jevitapearl/TaskForge/internal/repository"
	"github.com/jevitapearl/TaskForge/internal/router"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {

	// Loading configs
	cfg, err := config.LoadDBconfig()
	if err != nil {
		log.Fatal(err)
	}

	// Migrations
	m, err := migrate.New("file://internal/migrations", fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", cfg.User, cfg.Password, cfg.Host, cfg.DBName))
	if err != nil {
		log.Fatal(err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}


	// DB init
	db, err := database.NewPostgres(cfg)
	if err != nil {
		log.Fatal(err)
	}


	// Router init
	repo := repository.NewPostgresRepository(db)
	r := router.New(repo)

	// Middleware wrapping
	wrapped := middleware.LoggingMiddleware(middleware.RecoveryMiddleware(r.Handler()))

	// Serving HTTP req
	log.Println("Server running on port 8080")
	err = http.ListenAndServe(":8080", wrapped)
	if err != nil {
		panic(err)
	}

}
