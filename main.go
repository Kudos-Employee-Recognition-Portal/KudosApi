package main

import (
	"github.com/Kudos-Employee-Recognition-Portal/KudosApi/app"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	// Redirect server logging to file.
	//logfile, err := os.OpenFile("dev.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer logfile.Close()
	//log.SetOutput(logfile)
	//log.Print("Log file start.")

	// Load environment variables from .env file with godotenv package.
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file.")
	}
	// Instantiate App object from app package/app.go.
	api := app.App{}
	// Initialize database connection.
	api.InitDB(
		os.Getenv("API_DB_USERNAME"),
		os.Getenv("API_DB_PASSWORD"),
		os.Getenv("API_DB_HOST"),
		os.Getenv("API_DB_PORT"),
		os.Getenv("API_DB_NAME"))
	// Build routes.
	api.InitRouter()
	// Serve.
	api.Run(os.Getenv("API_PORT"))
}
