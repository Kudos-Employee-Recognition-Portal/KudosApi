package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	// Load environment variables from .env file with godotenv package.
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file.")
	}
	// Instantiate App object from app.go.
	app := App{}
	app.Init(os.Getenv("API_DB_USERNAME"), os.Getenv("API_DB_PASSWORD"), os.Getenv("API_DB_NAME"))
	app.Run(":8080")
}

