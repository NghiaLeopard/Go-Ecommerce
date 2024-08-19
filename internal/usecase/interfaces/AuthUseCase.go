package IUseCase

import (
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/handler/response"
	"github.com/gin-gonic/gin"
)

type IAuthUseCase interface {
	LoginUseCase(ctx *gin.Context, email string, password string) (response.LoginResponse, error, int)
	RegisterUseCase(ctx *gin.Context, email string, password string) error
	LogoutUseCase(ctx *gin.Context) (error, int)
	ChangePasswordUseCase(ctx *gin.Context,currentPassword string,newPassword string) (error,int)
}
