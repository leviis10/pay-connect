package payments

import (
	"pay-connect/common/database"
	"pay-connect/common/middleware"
	"pay-connect/features/payments/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	paymentsGroup := router.Group("/api/v1/payments")

	paymentsGroup.POST("/", middleware.AuthMiddleware(), handlers.CreatePayment(database.DB))
	paymentsGroup.PATCH("/:id/complete", middleware.AuthMiddleware(), handlers.UpdatePaymentStatus(database.DB))
}
