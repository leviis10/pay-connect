package dto

type UserRegisterDTO struct {
	Username string `binding:"required"`
	Email    string `binding:"required,email"`
	Password string `binding:"required"`
}
