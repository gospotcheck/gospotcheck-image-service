package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadEnvironmentFile sets environment variables specified in .env file.
func LoadEnvironmentFile() {
	err := godotenv.Load()
	if err == nil {
		log.Println("Loaded environment variables from .env file")
	}
}

// Getenv returns value of environment variable varnName.
func Getenv(varName string) string {
	return os.Getenv(varName)
}

/*
GetRequiredenv returns value of environment variable varnName.  Logs fatal
error if environment variable varName not set.
*/
func GetRequiredenv(varName string) string {
	value := Getenv(varName)
	if value == "" {
		log.Fatalf("Missing required environment variable: %v", varName)
	}
	return value
}
