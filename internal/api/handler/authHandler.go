package handler

import (
	IHandler "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/handler/interfaces"
	IRequest "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/handler/request"
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
	var req *IRequest.RegisterRequest

	err := ctx.ShouldBindJSON(&req)

	if err != nil {
		response.ErrorResponse(ctx,"Email or password is invalid",400)
		return
	}

	user, err := a.AuthUseCase.LoginUseCase(ctx,req.Email,req.Password)


}

// SignUpUser implements IHandler.IAuthHandler.
func (a *AuthHandler) SignUpUser(ctx *gin.Context) {
	var req *IRequest.RegisterRequest

	err := ctx.ShouldBindJSON(&req)

	if err != nil {
		response.ErrorResponse(ctx,"Email or password is invalid",400)
		return
	}

	err = a.AuthUseCase.RegisterUseCase(ctx,req.Email,req.Password)

	if err != nil {
		response.ErrorResponse(ctx,err.Error(),500)
		return
	}
	
	response.SuccessResponse(ctx,"Register success",200,"")

}


