package handler

import (
	"fmt"

	"github.com/NghiaLeopard/Go-Ecommerce-Backend/global"
	IHandler "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/handler/interfaces"
	IRequest "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/handler/request"
	IUseCase "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/usecase/interfaces"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/pkg/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ProductTypeHandler struct {
	ProductTypeUseCase IUseCase.ProductType
}

func NewProductTypeHandler(ProductTypeUseCase IUseCase.ProductType) IHandler.ProductType {
	return &ProductTypeHandler{ProductTypeUseCase: ProductTypeUseCase}
}

func (r *ProductTypeHandler) CreateProductType(ctx *gin.Context) {
	var req IRequest.CreateProductType
	if err := ctx.ShouldBindJSON(&req); err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		response.ErrorResponse(ctx, "Body is invalid or not exist", 400)
		return
	}

	ProductType, err, codeStatus := r.ProductTypeUseCase.CreateProductType(ctx, req.Name, req.Slug)

	if err != nil {
		response.ErrorResponse(ctx, err.Error(), codeStatus)
		return
	}

	global.Logger.Error("create ProductType", zap.String("Status", "Error"))
	response.SuccessResponse(ctx, "Create ProductType success", codeStatus, ProductType)
}

func (c *ProductTypeHandler) GetProductType(ctx *gin.Context) {
	var req IRequest.GetProductType
	if err := ctx.ShouldBindUri(&req); err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		response.ErrorResponse(ctx, "Body is invalid or not exist", 400)
		return
	}

	fmt.Println(req.ID)

	ProductType, err, codeStatus := c.ProductTypeUseCase.GetProductTypeUseCase(ctx, req.ID)

	if err != nil {
		response.ErrorResponse(ctx, err.Error(), codeStatus)
		return
	}

	global.Logger.Error("get ProductType", zap.String("Status", "Error"))
	response.SuccessResponse(ctx, "Get ProductType success", codeStatus, ProductType)

}

func (c *ProductTypeHandler) GetAllProductType(ctx *gin.Context) {
	var req IRequest.GetAllProductType
	if err := ctx.ShouldBindQuery(&req); err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		response.ErrorResponse(ctx, "Body is invalid or not exist", 400)
		return
	}

	ProductType, err, codeStatus := c.ProductTypeUseCase.GetAllProductTypeUseCase(ctx, req)

	if err != nil {
		response.ErrorResponse(ctx, err.Error(), codeStatus)
		return
	}

	global.Logger.Error("get ProductType", zap.String("Status", "Error"))
	response.SuccessResponse(ctx, "Get ProductType success", codeStatus, ProductType)

}

func (c *ProductTypeHandler) UpdateProductType(ctx *gin.Context) {
	var params IRequest.GetParamsUpdateProductType
	var body IRequest.GetBodyUpdateProductType
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

	ProductType, err, codeStatus := c.ProductTypeUseCase.UpdateProductTypeUseCase(ctx, params.ID, body.Name, body.Slug)

	if err != nil {
		response.ErrorResponse(ctx, err.Error(), codeStatus)
		return
	}

	global.Logger.Error("get ProductType", zap.String("Status", "Error"))
	response.SuccessResponse(ctx, "Get ProductType success", codeStatus, ProductType)
}

func (c *ProductTypeHandler) DeleteProductType(ctx *gin.Context) {
	var req IRequest.DeleteProductType
	if err := ctx.ShouldBindUri(&req); err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		response.ErrorResponse(ctx, "Body is invalid or not exist", 400)
		return
	}

	err, codeStatus := c.ProductTypeUseCase.DeleteProductTypeUseCase(ctx, req.ID)

	if err != nil {
		response.ErrorResponse(ctx, err.Error(), codeStatus)
		return
	}

	global.Logger.Error("get ProductType", zap.String("Status", "Error"))
	response.SuccessResponse(ctx, "Delete ProductType success", codeStatus, "")
}

func (c *ProductTypeHandler) DeleteManyProductType(ctx *gin.Context) {
	var req IRequest.DeleteManyProductType
	if err := ctx.ShouldBindJSON(&req); err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		response.ErrorResponse(ctx, "Body is invalid or not exist", 400)
		return
	}

	err, codeStatus := c.ProductTypeUseCase.DeleteManyProductTypeUseCase(ctx, req.ArrayId)

	if err != nil {
		response.ErrorResponse(ctx, err.Error(), codeStatus)
		return
	}

	global.Logger.Error("get ProductType", zap.String("Status", "Error"))
	response.SuccessResponse(ctx, "Delete ProductType success", codeStatus, "")
}
