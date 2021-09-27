package env

import (
	"log"

	"github.com/joho/godotenv"
)

// Load enviroments
func Load() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file", err)
	}
}
