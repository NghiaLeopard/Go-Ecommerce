package IRequest

import _ "encoding/base64"

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type ChangePasswordRequest struct {
	CurrentPassword string `json:"currentPassword" binding:"required,min=6"`
	NewPassword     string `json:"newPassword" binding:"required,min=6"`
}

type ResetPasswordRequest struct {
	NewPassword string `json:"newPassword" binding:"required,min=6"`
	SecretKey   string `json:"secretKey" binding:"required"`
}

type ForgotPasswordRequest struct {
	Email string `json:"email" binding:"required,email"`
}

type UpdateAuthMe struct {
	Address     string `json:"address" binding:"required"`
	Avatar      string `json:"avatar" binding:"required"`
	City        int64  `json:"city" binding:"required" format:"int64"`
	FirstName   string `json:"firstName" binding:"required"`
	LastName    string `json:"lastName" binding:"required"`
	MiddleName  string `json:"middleName" binding:"required"`
	PhoneNumber int64  `json:"phoneNumber" binding:"required" format:"int64"`
	Image       string `json:"image" binding:"required"`
}
