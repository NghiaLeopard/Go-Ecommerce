package IHandler

import "github.com/gin-gonic/gin"

type Auth interface {
	LoginUser(ctx *gin.Context)
	SignUpUser(ctx *gin.Context)
	LogoutUser(ctx *gin.Context)
	ChangePasswordUser(ctx *gin.Context)
	ForgotPasswordUser(ctx *gin.Context)
	ResetPasswordUser(ctx *gin.Context)
	GetAuthMe(ctx *gin.Context)
	RefreshToken(ctx *gin.Context)
}
