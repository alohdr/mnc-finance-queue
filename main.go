package main

import (
	"github.com/joho/godotenv"
	"log"
	"mnc-finance-queue/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	err = routes.SetupRoutes()
	if err != nil {
		log.Println("[Main][Server][Error]: ", err)
		return
	}
}
