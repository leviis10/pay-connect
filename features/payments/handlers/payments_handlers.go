package handlers

import (
	"net/http"
	"pay-connect/common/response"
	"pay-connect/features/payments/dto"
	"pay-connect/features/payments/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreatePayment(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var json *dto.CreatePaymentDTO

		if err := ctx.ShouldBindJSON(&json); err != nil {
			ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
				Status: "Error",
				Error:  err.Error(),
			})
			return
		}

		payment := models.Payment{
			Amount:             json.Amount,
			SenderCustomerID:   uint(ctx.MustGet("id").(float64)),
			ReceiverCustomerID: json.ReceiverCustomerID,
			Status:             models.PaymentStatusPending,
		}

		result := db.Create(&payment)
		if result.Error != nil {
			ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
				Status: "Error",
				Error:  result.Error.Error(),
			})
			return
		}

		ctx.JSON(http.StatusCreated, response.SuccessResponse{
			Status:  "Success",
			Message: "Payment created successfully",
			Data:    payment,
		})
	}
}

func UpdatePaymentStatus(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId, ok := ctx.MustGet("id").(float64)
		if !ok {
			ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
				Status: "Error",
				Error:  "Invalid user ID type",
			})
			return
		}

		paymentIdStr := ctx.Param("id")
		paymentId, err := strconv.ParseUint(paymentIdStr, 10, 32)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
				Status: "Error",
				Error:  "Invalid ID format",
			})
			return
		}

		var payment models.Payment
		result := db.Where(&models.Payment{ID: uint(paymentId), SenderCustomerID: uint(userId)}).First(&payment)
		if result.RowsAffected == 0 {
			ctx.JSON(http.StatusNotFound, response.ErrorResponse{
				Status: "Error",
				Error:  "Payment not found",
			})
			return
		}

		payment.Status = models.PaymentStatusCompleted
		db.Save(&payment)
		ctx.JSON(http.StatusOK, response.SuccessResponse{
			Status:  "Success",
			Message: "Payment status updated successfully",
		})
	}
}
