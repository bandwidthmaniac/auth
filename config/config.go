package config

import (
	"os"

	"github.com/joho/godotenv"
)

func Get(key string) string {
	// Load .env file
	godotenv.Load()
	return os.Getenv(key)
}
