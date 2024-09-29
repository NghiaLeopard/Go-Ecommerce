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

type ProductTypeHandler struct {
	ProductTypeUseCase IUseCase.ProductType
}

func NewProductTypeHandler(ProductTypeUseCase IUseCase.ProductType) IHandler.ProductType {
	return &ProductTypeHandler{ProductTypeUseCase: ProductTypeUseCase}
}

// CreateProductType 		godoc
// @security 				BearerAuth
// @Summary 				Create ProductType
// @Description 			Create ProductType
// @Param 					tags body IRequest.CreateProductType true "Create ProductType"
// @Produce 				application/json
// @Tags 					ProductType
// @Success 				200 {object} IResponse.ProductType{}
// @Router 					/api/product-types [post]
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

	global.Logger.Info("create ProductType", zap.String("Status", "success"))
	response.SuccessResponse(ctx, "Create ProductType success", codeStatus, ProductType)
}

// GetAllProductType 	godoc
// @security 			BearerAuth
// @Summary 			Get all ProductType
// @Description 		Get all ProductType
// @Param 				request query IRequest.GetAllProductType true "get all product type"
// @Produce 			application/json
// @Tags 				ProductType
// @Success 			200 {array} []IResponse.ProductType{}
// @Router 				/api/product-types [get]
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

	global.Logger.Info("get ProductType", zap.String("Status", "Success"))
	response.SuccessResponse(ctx, "Get ProductType success", codeStatus, ProductType)

}

// GetProductType 			godoc
// @security 				BearerAuth
// @Summary 				Get ProductType by id
// @Description 			Get ProductType by id
// @Param ProductTypeId  	path int true "User ID"
// @Produce 				application/json
// @Tags 					ProductType
// @Success 				200 {object} IResponse.ProductType{}
// @Router 					/api/product-types/{ProductTypeId} [get]
func (c *ProductTypeHandler) GetProductType(ctx *gin.Context) {
	var req IRequest.GetProductType
	if err := ctx.ShouldBindUri(&req); err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		response.ErrorResponse(ctx, "Body is invalid or not exist", 400)
		return
	}

	ProductType, err, codeStatus := c.ProductTypeUseCase.GetProductTypeUseCase(ctx, req.ID)

	if err != nil {
		response.ErrorResponse(ctx, err.Error(), codeStatus)
		return
	}

	global.Logger.Info("get ProductType", zap.String("Status", "Success"))
	response.SuccessResponse(ctx, "Get ProductType success", codeStatus, ProductType)

}

// UpdateProductType 		godoc
// @security 				BearerAuth
// @Summary 				Update ProductType
// @Description 			Update ProductType
// @Param ProductTypeId 	path int true "Update ProductType"
// @Param 					tags body IRequest.GetBodyUpdateProductType true "Update ProductType"
// @Produce 				application/json
// @Tags 					ProductType
// @Success 				200 {object} IResponse.ProductType{}
// @Router 					/api/product-types/{ProductTypeId} [put]
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

	global.Logger.Info("get ProductType", zap.String("Status", "Success"))
	response.SuccessResponse(ctx, "Get ProductType success", codeStatus, ProductType)
}

// DeleteProductType 		godoc
// @security 				BearerAuth
// @Summary 				Delete ProductType
// @Description 			Delete ProductType
// @Param ProductTypeId 	path int true "Delete ProductType"
// @Produce 				application/json
// @Tags 					ProductType
// @Success 				200 {string} string [delete ProductType success]
// @Router 					/api/product-types/{ProductTypeId} [delete]
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

	global.Logger.Info("get ProductType", zap.String("Status", "Success"))
	response.SuccessResponse(ctx, "Delete ProductType success", codeStatus, "")
}

// DeleteManyProductType 	godoc
// @security 				BearerAuth
// @Summary 				Delete many ProductType
// @Description 			Delete many ProductType
// @Param 					tags body IRequest.DeleteManyProductType true "DeleteMany ProductType"
// @Produce 				application/json
// @Tags 					ProductType
// @Success 				200 {string} string "Delete many ProductType success"
// @Router 					/api/product-types/delete-many [delete]
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

	global.Logger.Info("get ProductType", zap.String("Status", "Success"))
	response.SuccessResponse(ctx, "Delete ProductType success", codeStatus, "")
}
