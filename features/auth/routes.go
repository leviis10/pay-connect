package auth

import (
	"pay-connect/common/database"
	"pay-connect/features/auth/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	authGroup := router.Group("/api/v1/auth")

	authGroup.POST("/register", handlers.RegisterUser(database.DB))
	authGroup.POST("/login", handlers.UserLogin(database.DB))
}
