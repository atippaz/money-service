package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	SERVER_PORT string
	HOST        string
	USER        string
	DBNAME      string
	PASSWORD    string
	DB_PORT     string
	API_VERSION string
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	cfg := &Config{
		SERVER_PORT: os.Getenv("SERVER_PORT"),
		HOST:        os.Getenv("HOST"),
		USER:        os.Getenv("USER"),
		DBNAME:      os.Getenv("DBNAME"),
		PASSWORD:    os.Getenv("PASSWORD"),
		DB_PORT:     os.Getenv("DB_PORT"),
		API_VERSION: os.Getenv("API_VERSION"),
	}

	return cfg
}
