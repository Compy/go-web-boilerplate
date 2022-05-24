package main

import (
	"fmt"
	"go-web-boilerplate/database"
	"os"

	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("[ERROR] Could not load .env file. Configuration values will be pulled from the OS environment instead")
	}

	if len(os.Getenv("JWT_SECRET")) < 32 {
		log.Fatalf("[ERROR] JWT_SECRET must be at least 32 chars long")
		return
	}

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbPort := os.Getenv("DB_PORT")
	dbDatabase := os.Getenv("DB_DATABASE")

	database.Connect(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPass, dbHost, dbPort, dbDatabase))
	database.Migrate()

	router := initRouter()
	router.Run(":8080")
}
