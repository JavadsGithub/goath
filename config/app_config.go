package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	JWT_SECRET string
	DB_NAME    string
}

var AppConfig Config

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	AppConfig.DB_NAME = os.Getenv("DB_NAME")
	AppConfig.JWT_SECRET = os.Getenv("JWT_SECRET")

	if AppConfig.DB_NAME == "" {
		log.Fatal("Env: DB_NAME is required")
	}
	if AppConfig.JWT_SECRET == "" {
		log.Fatal("Env: JWT_SECRET is required")
	}
}
