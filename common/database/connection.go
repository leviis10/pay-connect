package database

import (
	"fmt"
	"log"
	"os"
	customerModel "pay-connect/features/customers"
	loginHistoryModel "pay-connect/features/login-histories"
	paymentModel "pay-connect/features/payments/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
}

func Migrate() {
	DB.AutoMigrate(&customerModel.Customer{})
	DB.AutoMigrate(&paymentModel.Payment{})
	DB.AutoMigrate(&loginHistoryModel.LoginHistory{})
}
