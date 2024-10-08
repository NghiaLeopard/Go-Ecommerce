package IUseCase

import (
	IRequest "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/handler/request"
	IResponse "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/handler/response"
	"github.com/gin-gonic/gin"
)

type Auth interface {
	LoginUseCase(ctx *gin.Context, email string, password string) (IResponse.Login, error, int)
	RegisterUseCase(ctx *gin.Context, email string, password string) error
	LogoutUseCase(ctx *gin.Context) (error, int)

	GetAuthMeUserCase(ctx *gin.Context) (IResponse.AuthMe, error, int)
	UpdateAuthMeUserCase(ctx *gin.Context, req IRequest.UpdateAuthMe) (IResponse.UpdateAuthMe, error, int)

	ChangePasswordUseCase(ctx *gin.Context, currentPassword string, newPassword string) (error, int)
	ForgotPasswordUseCase(ctx *gin.Context, email string) (error, int)
	ResetPasswordUseCase(ctx *gin.Context, newPassword string, secretKey string) (error, int)
	RefreshTokenUseCase(ctx *gin.Context) (IResponse.GetAccessToken, error, int)
}
