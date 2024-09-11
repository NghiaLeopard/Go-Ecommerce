package handler

import (
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/global"
	IHandler "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/handler/interfaces"
	IRequest "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/handler/request"
	IUseCase "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/usecase/interfaces"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/pkg/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AuthHandler struct {
	AuthUseCase IUseCase.Auth
}

func NewAuthHandler(authUseCase IUseCase.Auth) IHandler.Auth {
	return &AuthHandler{AuthUseCase: authUseCase}
}

func (a *AuthHandler) LoginUser(ctx *gin.Context) {
	var req *IRequest.RegisterRequest

	err := ctx.ShouldBindJSON(&req)

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		response.ErrorResponse(ctx, "Email or password is invalid", 400)
		return
	}

	user, err, statusCode := a.AuthUseCase.LoginUseCase(ctx, req.Email, req.Password)

	if err != nil {
		response.ErrorResponse(ctx, err.Error(), statusCode)
		return
	}

	global.Logger.Info("Login", zap.String("Status", "Success"))
	response.SuccessResponse(ctx, "Login success", 201, user)
}

func (a *AuthHandler) SignUpUser(ctx *gin.Context) {
	var req *IRequest.RegisterRequest

	err := ctx.ShouldBindJSON(&req)

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		response.ErrorResponse(ctx, "Email or password is invalid", 400)
		return
	}

	err = a.AuthUseCase.RegisterUseCase(ctx, req.Email, req.Password)

	if err != nil {
		response.ErrorResponse(ctx, err.Error(), 500)
		return
	}

	global.Logger.Info("Register", zap.String("Status", "Success"))
	response.SuccessResponse(ctx, "Register success", 201, "")
}

func (a *AuthHandler) LogoutUser(ctx *gin.Context) {
	err, statusCode := a.AuthUseCase.LogoutUseCase(ctx)

	if err != nil {
		response.ErrorResponse(ctx, err.Error(), 500)
		return
	}

	global.Logger.Info("Logout", zap.String("Status", "Success"))
	response.SuccessResponse(ctx, "Logout success", statusCode, "")
}

func (a *AuthHandler) ChangePasswordUser(ctx *gin.Context) {
	var req *IRequest.ChangePasswordRequest

	err := ctx.ShouldBindJSON(&req)

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		response.ErrorResponse(ctx, "Current password or new password is invalid", 400)
		return
	}

	err, statusCode := a.AuthUseCase.ChangePasswordUseCase(ctx, req.CurrentPassword, req.NewPassword)

	if err != nil {
		response.ErrorResponse(ctx, err.Error(), statusCode)
		return
	}

	global.Logger.Info("Change password", zap.String("Status", "Success"))
	response.SuccessResponse(ctx, "change password success", statusCode, "")
}

func (a *AuthHandler) ForgotPasswordUser(ctx *gin.Context) {
	var req *IRequest.ForgotPasswordRequest

	err := ctx.ShouldBindJSON(&req)

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		response.ErrorResponse(ctx, "Email is invalid", 400)
		return
	}

	err, statusCode := a.AuthUseCase.ForgotPasswordUseCase(ctx, req.Email)

	if err != nil {
		response.ErrorResponse(ctx, err.Error(), statusCode)
		return
	}

	global.Logger.Info("Forgot password", zap.String("Status", "Success"))
	response.SuccessResponse(ctx, "Send gmail", statusCode, "")
}

func (a *AuthHandler) ResetPasswordUser(ctx *gin.Context) {
	var res *IRequest.ResetPasswordRequest

	err := ctx.ShouldBindJSON(&res)

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		response.ErrorResponse(ctx, "value invalid", 400)
		return
	}

	err, statusCode := a.AuthUseCase.ResetPasswordUseCase(ctx, res.NewPassword, res.SecretKey)

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		response.ErrorResponse(ctx, err.Error(), statusCode)
		return
	}

	global.Logger.Info("Reset password", zap.String("Status", "Success"))
	response.SuccessResponse(ctx, "Reset password success", statusCode, "")
}

func (a *AuthHandler) GetAuthMe(ctx *gin.Context) {
	authMe, err, codeStatus := a.AuthUseCase.GetAuthMeUserCase(ctx)

	if err != nil {
		response.ErrorResponse(ctx, err.Error(), codeStatus)
	}

	global.Logger.Info("Get auth", zap.String("Status", "Success"))
	response.SuccessResponse(ctx, "Get auth me success", codeStatus, authMe)
}
