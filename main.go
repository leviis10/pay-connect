package main

import (
	"log"
	"pay-connect/common/database"
	"pay-connect/features/auth"
	"pay-connect/features/payments"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load Environment Variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Database Connection
	database.Connect()
	database.Migrate()

	router := gin.Default()

	auth.RegisterRoutes(router)
	payments.RegisterRoutes(router)

	router.Run(":8080")
}
