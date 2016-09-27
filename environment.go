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

func Getenv(varName string) string {
	return os.Getenv(varName)
}

func GetRequiredenv(varName string) string {
	value := Getenv(varName)
	if value == "" {
		log.Fatalf("Missing required environment variable: %v", varName)
	}
	return value
}
