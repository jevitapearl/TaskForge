package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	User     string `env:"DBUSER"`
	DBPort   int    `env:"PORT"`
	Host     string `env:"HOST"`
	Password string `env:"DBPASSWORD"`
	DBName   string `env:"DBNAME"`
}

func LoadDBconfig() (*DBConfig, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}
	port, err := strconv.Atoi(os.Getenv("DBPORT"))
	if err != nil {
		return nil, err
	}

	cfg := &DBConfig{
		User:     os.Getenv("DBUSER"),
		DBPort:   port,
		Host:     os.Getenv("HOST"),
		Password: os.Getenv("DBPASSWORD"),
		DBName:   os.Getenv("DBNAME"),
	}

	return cfg, nil
}
