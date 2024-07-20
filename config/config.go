package config

import (
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

var UserId = uuid.MustParse("e8bc8014-4a5a-4e91-a6d2-3b5b6f77d3e0")
var cfg = LoadConfig()

type Config struct {
	SERVER_PORT string
	HOST        string
	USER        string
	DBNAME      string
	PASSWORD    string
	DB_PORT     string
	API_VERSION string
	JWTSECERT   string
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	return &Config{
		SERVER_PORT: os.Getenv("SERVER_PORT"),
		HOST:        os.Getenv("HOST"),
		USER:        os.Getenv("USER"),
		DBNAME:      os.Getenv("DBNAME"),
		PASSWORD:    os.Getenv("PASSWORD"),
		DB_PORT:     os.Getenv("DB_PORT"),
		API_VERSION: os.Getenv("API_VERSION"),
		JWTSECERT:   os.Getenv("SECERT_KEY"),
	}
}
