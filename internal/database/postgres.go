package database

import (
	"database/sql"
	"fmt"

	"github.com/jevitapearl/TaskForge/internal/config"
)

func NewPostgres(cfg *config.DBConfig) (*sql.DB, error) {

	dsn := fmt.Sprintf("host=%v dbname=%v user=%v password=%v port=%v sslmode=disable", cfg.Host, cfg.DBName, cfg.User, cfg.Password, cfg.DBPort)
	db, err := sql.Open("postgres", dsn)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
