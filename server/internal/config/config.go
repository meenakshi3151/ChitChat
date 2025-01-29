package config

import (
	"log"
	"fmt"
	"github.com/joho/godotenv"
)

// LoadEnv loads the environment variables from .env file
func LoadEnv() {
	err := godotenv.Load()
    fmt.Println("Loading env file")
	if err != nil {
		log.Println("Warning: No .env file found, using system environment variables")
	}
}
