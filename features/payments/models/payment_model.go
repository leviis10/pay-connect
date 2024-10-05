package models

import "time"

type PaymentStatus string

const (
	PaymentStatusCompleted PaymentStatus = "completed"
	PaymentStatusPending   PaymentStatus = "pending"
)

type Payment struct {
	ID                 uint
	Amount             float64
	SenderCustomerID   uint
	ReceiverCustomerID uint
	Status             PaymentStatus `gorm:"type:payment_status"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
}
