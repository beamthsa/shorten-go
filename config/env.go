package config

import (
	"github.com/opn-ooo/gin-boilerplate/config/database"
	"os"
)

type GoDotENV struct {
	Port           string                  `json:"port"`
	GinMode        string                  `json:"ginMode"`
	PostgresConfig database.PostgresConfig `json:"postgres"`
}

// GetGoDotENV |return| prepare/format/init of configs from .env
func GetGoDotENV() GoDotENV {
	return GoDotENV{
		// Gin
		Port:    ":" + os.Getenv("PORT"),
		GinMode: os.Getenv("GIN_MODE"),
		// PostgresConfig
		PostgresConfig: setPostgres(),
	}
}

func setPostgres() database.PostgresConfig {
	return database.PostgresConfig{
		Host:     os.Getenv("POSTGRES_HOST"),
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Database: os.Getenv("POSTGRES_DBNAME"),
		Port:     os.Getenv("POSTGRES_PORT"),
	}
}
