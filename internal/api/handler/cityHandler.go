package handler

import (
	IHandler "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/handler/interfaces"
	IRequest "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/handler/request"
	IUseCase "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/usecase/interfaces"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/pkg/response"
	"github.com/gin-gonic/gin"
)

type CityHandler struct {
	CityUseCase IUseCase.ICityUseCase
}

func NewCityHandler(cityUseCase IUseCase.ICityUseCase) IHandler.ICityHandler {
	return &CityHandler{CityUseCase: cityUseCase}
}

// CreateCity implements IHandler.ICityHandler.
func (c *CityHandler) CreateCity(ctx *gin.Context) {
	var req IRequest.CreateCity
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(ctx, "Body is invalid or not exist", 400)
		return
	}

	city, err, codeStatus := c.CityUseCase.CreateCityUseCase(ctx, req.Name)

	if err != nil {
		response.ErrorResponse(ctx, err.Error(), codeStatus)
		return
	}

	response.SuccessResponse(ctx, "Create user success", codeStatus, city)

}

// CreateCity implements IHandler.ICityHandler.
func (c *CityHandler) GetCity(ctx *gin.Context) {
	var req IRequest.GetCity
	if err := ctx.ShouldBindUri(&req); err != nil {
		response.ErrorResponse(ctx, "Body is invalid or not exist", 400)
		return
	}

	city, err, codeStatus := c.CityUseCase.GetCityUseCase(ctx, req.ID)

	if err != nil {
		response.ErrorResponse(ctx, err.Error(), codeStatus)
		return
	}

	response.SuccessResponse(ctx, "Get user success", codeStatus, city)

}
