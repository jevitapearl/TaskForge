package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	m, err := migrate.New(
		"file://internal/migrations",
		fmt.Sprintf(
			"postgres://%s:%s@%s:%v/%s?sslmode=disable",
			cfg.User,
			cfg.Password,
			cfg.Host,
			cfg.DBPort,
			cfg.DBName,
		),
	)

	if err != nil {
		log.Fatal(err)
	}

	defer m.Close()
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
	wrapped := middleware.RecoveryMiddleware(r.Handler())

	server := &http.Server{
		Addr:    ":8080",
		Handler: wrapped,
	}

	// Serving HTTP req

	sigCh := make(chan os.Signal, 1)

	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Println("Server running on port 8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Println("Server stopped:", err)
		}
	}()

	sig := <-sigCh
	log.Println("Shutdown signal received:", sig)

	shutdownCtx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Println("Graceful shutdown failed:", err)
	}

	db.Close()
	log.Println("Server shutdown complete")

}
