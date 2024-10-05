package handlers

import (
	"net/http"
	"pay-connect/common/response"
	"pay-connect/common/utils"
	"pay-connect/features/auth/dto"
	customerModels "pay-connect/features/customers"
	loginHistoryModel "pay-connect/features/login-histories"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterUser(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var json *dto.UserRegisterDTO

		if err := ctx.ShouldBindJSON(&json); err != nil {
			ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
				Status: "Error",
				Error:  err.Error(),
			})
			return
		}

		hashedPassword, err := utils.HashPassword(json.Password)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
				Status: "error",
				Error:  "Internal Server Error",
			})
			return
		}

		customer := customerModels.Customer{
			Username: json.Username,
			Email:    json.Email,
			Password: hashedPassword,
		}

		result := db.Create(&customer)
		if result.Error != nil {
			ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
				Status: "Error",
				Error:  result.Error.Error(),
			})
			return
		}

		ctx.JSON(http.StatusCreated, response.SuccessResponse{
			Status:  "Success",
			Message: "Registration Successfull",
		})
	}
}

func UserLogin(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var customer customerModels.Customer
		var json struct {
			Username string `binding:"required"`
			Password string `binding:"required"`
		}

		if err := ctx.ShouldBindJSON(&json); err != nil {
			ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
				Status: "Error",
				Error:  err.Error(),
			})
			return
		}

		if err := db.Where("username = ?", json.Username).First(&customer).Error; err != nil {
			ctx.JSON(http.StatusUnauthorized, response.ErrorResponse{
				Status: "Error",
				Error:  "Wrong username or password",
			})
			return
		}

		loginHistory := loginHistoryModel.LoginHistory{
			CustomerID: customer.ID,
		}

		if !utils.CompareHashPassword(customer.Password, json.Password) {
			loginHistory.Status = loginHistoryModel.LoginStatusFailure
			result := db.Create(&loginHistory)
			if result.Error != nil {
				ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
					Status: "Error",
					Error:  result.Error.Error(),
				})
				return
			}
			ctx.JSON(http.StatusUnauthorized, response.ErrorResponse{
				Status: "Error",
				Error:  "Wrong username or password",
			})
			return
		}

		loginHistory.Status = loginHistoryModel.LoginStatusSuccess
		result := db.Create(&loginHistory)
		if result.Error != nil {
			ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
				Status: "Error",
				Error:  result.Error.Error(),
			})
			return
		}

		token, err := utils.GenerateToken(customer)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
				Status: "Error",
				Error:  err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, response.SuccessResponse{
			Status:  "Success",
			Message: "Logged in",
			Data:    map[string]string{"token": token},
		})
	}
}
