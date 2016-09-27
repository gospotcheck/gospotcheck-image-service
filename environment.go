package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvironmentFile() {
	err := godotenv.Load()
	if err == nil {
		log.Println("Loaded environment variables from .env file")
	}
}

func GetEnv(varName string, required bool) string {
	var value = os.Getenv(varName)
	if required && value == "" {
		log.Fatalf("Missing required environment variable: %v", varName)
	}
	return value
}
