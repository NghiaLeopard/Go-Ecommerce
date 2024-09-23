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

type CityHandler struct {
	CityUseCase IUseCase.City
}

func NewCityHandler(cityUseCase IUseCase.City) IHandler.City {
	return &CityHandler{CityUseCase: cityUseCase}
}

func (c *CityHandler) CreateCity(ctx *gin.Context) {
	var req IRequest.CreateCity
	if err := ctx.ShouldBindJSON(&req); err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		response.ErrorResponse(ctx, "Body is invalid or not exist", 400)
		return
	}

	city, err, codeStatus := c.CityUseCase.CreateCityUseCase(ctx, req.Name)

	if err != nil {
		response.ErrorResponse(ctx, err.Error(), codeStatus)
		return
	}

	global.Logger.Error("create city", zap.String("Status", "Error"))
	response.SuccessResponse(ctx, "Create city success", codeStatus, city)
}

func (c *CityHandler) GetCity(ctx *gin.Context) {
	var req IRequest.GetCity
	if err := ctx.ShouldBindUri(&req); err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		response.ErrorResponse(ctx, "Body is invalid or not exist", 400)
		return
	}

	city, err, codeStatus := c.CityUseCase.GetCityUseCase(ctx, req.ID)

	if err != nil {
		response.ErrorResponse(ctx, err.Error(), codeStatus)
		return
	}

	global.Logger.Error("get city", zap.String("Status", "Error"))
	response.SuccessResponse(ctx, "Get city success", codeStatus, city)
}

func (c *CityHandler) GetAllCity(ctx *gin.Context) {
	var req IRequest.GetAllCity
	if err := ctx.ShouldBindJSON(&req); err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		response.ErrorResponse(ctx, "Body is invalid or not exist", 400)
		return
	}

	city, err, codeStatus := c.CityUseCase.GetAllCityUseCase(ctx, req.Page, req.Limit, req.Search, req.Order)

	if err != nil {
		response.ErrorResponse(ctx, err.Error(), codeStatus)
		return
	}

	global.Logger.Error("get city", zap.String("Status", "Error"))
	response.SuccessResponse(ctx, "Get city success", codeStatus, city)
}

func (c *CityHandler) UpdateCity(ctx *gin.Context) {
	var params IRequest.GetParamsUpdateCity
	var body IRequest.GetBodyUpdateCity
	if err := ctx.ShouldBindUri(&params); err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		response.ErrorResponse(ctx, "Body is invalid or not exist", 400)
		return
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		response.ErrorResponse(ctx, "Body is invalid or not exist", 400)
		return
	}

	city, err, codeStatus := c.CityUseCase.UpdateCityUseCase(ctx, params.ID, body.Name)

	if err != nil {
		response.ErrorResponse(ctx, err.Error(), codeStatus)
		return
	}

	global.Logger.Error("get city", zap.String("Status", "Error"))
	response.SuccessResponse(ctx, "Get city success", codeStatus, city)
}

func (c *CityHandler) DeleteCity(ctx *gin.Context) {
	var req IRequest.DeleteCity
	if err := ctx.ShouldBindUri(&req); err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		response.ErrorResponse(ctx, "Body is invalid or not exist", 400)
		return
	}

	err, codeStatus := c.CityUseCase.DeleteCityUseCase(ctx, req.ID)

	if err != nil {
		response.ErrorResponse(ctx, err.Error(), codeStatus)
		return
	}

	global.Logger.Error("get city", zap.String("Status", "Error"))
	response.SuccessResponse(ctx, "Delete city success", codeStatus, "")
}

func (c *CityHandler) DeleteManyCity(ctx *gin.Context) {
	var req IRequest.DeleteManyCity
	if err := ctx.ShouldBindJSON(&req); err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		response.ErrorResponse(ctx, "Body is invalid or not exist", 400)
		return
	}

	err, codeStatus := c.CityUseCase.DeleteManyCityUseCase(ctx, req.ArrayId)

	if err != nil {
		response.ErrorResponse(ctx, err.Error(), codeStatus)
		return
	}

	global.Logger.Error("get city", zap.String("Status", "Error"))
	response.SuccessResponse(ctx, "Delete city success", codeStatus, "")
}
