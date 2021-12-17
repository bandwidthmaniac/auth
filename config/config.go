package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Get(key string) string {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	return os.Getenv(key)
}
