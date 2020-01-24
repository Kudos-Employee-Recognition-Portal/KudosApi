package main

import (
	"./app"
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
	api := app.App{}

	api.InitDB(os.Getenv("API_DB_USERNAME"), os.Getenv("API_DB_PASSWORD"), os.Getenv("API_DB_NAME"))

	api.InitRouter()

	api.Run(os.Getenv("API_PORT"))
}
