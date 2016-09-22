package main

import (
    "github.com/joho/godotenv"
    "log"
)

func LoadEnvironment() {
  err := godotenv.Load()
  if err == nil {
    log.Println("Loaded environment variables from .env file")
  }
}
