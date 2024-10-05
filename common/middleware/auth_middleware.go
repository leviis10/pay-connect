package middleware

import (
	"net/http"
	"pay-connect/common/response"
	"pay-connect/common/utils"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.Request.Header.Get("Authorization")
		if tokenString == "" {
			ctx.JSON(http.StatusUnauthorized, response.ErrorResponse{
				Status: "Error",
				Error:  "Unauthorized",
			})
			ctx.Abort()
			return
		}

		tokenString = tokenString[len("Bearer "):]
		claims, err := utils.ValidateToken(tokenString)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, response.ErrorResponse{
				Status: "Error",
				Error:  err.Error(),
			})
			ctx.Abort()
			return
		}

		ctx.Set("id", claims["id"])
		ctx.Set("username", claims["username"])
		ctx.Set("email", claims["email"])
		ctx.Next()
	}
}
