package handler

import (
	IHandler "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/handler/interfaces"
	IUseCase "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/usecase/interfaces"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/pkg/response"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	AuthUseCase IUseCase.IAuthUseCase
}

func NewAuthHandler(authUseCase IUseCase.IAuthUseCase) IHandler.IAuthHandler {
	return &AuthHandler{AuthUseCase: authUseCase}
}

// LoginUser implements IHandler.IAuthHandler.
func (a *AuthHandler) LoginUser(ctx *gin.Context) {
	panic("unimplemented")
}

// SignUpUser implements IHandler.IAuthHandler.
func (a *AuthHandler) SignUpUser(ctx *gin.Context) {
	response.SuccessResponse(ctx,"Register success",200,"")
}


