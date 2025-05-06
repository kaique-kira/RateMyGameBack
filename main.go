package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }
  fmt.Println("Environment variables loaded successfully")
  fmt.Println("API_KEY:", os.Getenv("API_KEY"))
}