package models

import (
	loginHistoryModel "pay-connect/features/login-histories"
	paymentModel "pay-connect/features/payments/models"

	"time"
)

type Customer struct {
	ID              uint
	Username        string `gorm:"unique"`
	Password        string
	Email           string `gorm:"unique"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	LoginHistory    []loginHistoryModel.LoginHistory
	PaymentReceived []paymentModel.Payment `gorm:"foreignKey:ReceiverCustomerID"`
	PaymentSent     []paymentModel.Payment `gorm:"foreignKey:SenderCustomerID"`
}
