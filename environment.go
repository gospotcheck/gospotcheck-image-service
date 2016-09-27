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

func GetEnv(varName string) string {
	return os.Getenv(varName)
}

func GetRequiredEnv(varName string) string {
	value := GetEnv(varName)
	if value == "" {
		log.Fatalf("Missing required environment variable: %v", varName)
	}
	return value
}
