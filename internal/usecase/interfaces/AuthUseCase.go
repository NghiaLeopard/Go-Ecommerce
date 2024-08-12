package IUseCase

import (
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/handler/response"
	"github.com/gin-gonic/gin"
)

type IAuthUseCase interface {
	LoginUseCase(ctx *gin.Context,email string, password string) (response.LoginResponse,error)
	RegisterUseCase(ctx *gin.Context,email string, password string) error
}