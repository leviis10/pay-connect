package dto

type CreatePaymentDTO struct {
	Amount             float64 `json:"amount" binding:"required"`
	ReceiverCustomerID uint    `json:"receiver_customer_id" binding:"required"`
}
