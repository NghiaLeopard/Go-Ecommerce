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

type ProductHandler struct {
	ProductUseCase IUseCase.Product
}

func NewProductHandler(ProductUseCase IUseCase.Product) IHandler.Product {
	return &ProductHandler{ProductUseCase: ProductUseCase}
}

// CreateProduct 			godoc
// @security 				BearerAuth
// @Summary 				Create Product
// @Description 			Create Product
// @Param 					tags body IRequest.CreateProduct true "Create Product"
// @Produce 				application/json
// @Tags 					Product
// @Success 				200 {object} IResponse.Product{}
// @Router 					/api/product-types [post]
func (r *ProductHandler) CreateProduct(ctx *gin.Context) {
	var req IRequest.CreateProduct
	if err := ctx.ShouldBindJSON(&req); err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		response.ErrorResponse(ctx, "Body is invalid or not exist", 400)
		return
	}

	Product, err, codeStatus := r.ProductUseCase.CreateProduct(ctx, req)

	if err != nil {
		response.ErrorResponse(ctx, err.Error(), codeStatus)
		return
	}

	global.Logger.Error("create Product", zap.String("Status", "Error"))
	response.SuccessResponse(ctx, "Create Product success", codeStatus, Product)
}

// GetAllProduct 	godoc
// @security 			BearerAuth
// @Summary 			Get all Product
// @Description 		Get all Product
// @Param 				request query IRequest.GetAllProduct true "get all product type"
// @Produce 			application/json
// @Tags 				Product
// @Success 			200 {array} []IResponse.Product{}
// @Router 				/api/product-types [get]
// func (c *ProductHandler) GetAllProduct(ctx *gin.Context) {
// 	var req IRequest.GetAllProduct
// 	if err := ctx.ShouldBindQuery(&req); err != nil {
// 		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
// 		response.ErrorResponse(ctx, "Body is invalid or not exist", 400)
// 		return
// 	}

// 	Product, err, codeStatus := c.ProductUseCase.GetAllProductUseCase(ctx, req)

// 	if err != nil {
// 		response.ErrorResponse(ctx, err.Error(), codeStatus)
// 		return
// 	}

// 	global.Logger.Error("get Product", zap.String("Status", "Error"))
// 	response.SuccessResponse(ctx, "Get Product success", codeStatus, Product)

// }

// GetProduct 				godoc
// @security 				BearerAuth
// @Summary 				Get Product by id
// @Description 			Get Product by id
// @Param ProductId  		path int true "product ID"
// @Produce 				application/json
// @Tags 					Product
// @Success 				200 {object} IResponse.Product{}
// @Router 					/api/product/{ProductId} [get]
func (c *ProductHandler) GetProduct(ctx *gin.Context) {
	var req IRequest.GetProduct
	if err := ctx.ShouldBindUri(&req); err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		response.ErrorResponse(ctx, "Body is invalid or not exist", 400)
		return
	}

	Product, err, codeStatus := c.ProductUseCase.GetProductUseCase(ctx, req.ID)

	if err != nil {
		response.ErrorResponse(ctx, err.Error(), codeStatus)
		return
	}

	global.Logger.Error("get Product", zap.String("Status", "Error"))
	response.SuccessResponse(ctx, "Get Product success", codeStatus, Product)
}

// GetProduct 				godoc
// @security 				BearerAuth
// @Summary 				Get Product by id
// @Description 			Get Product by id
// @Param ProductId  		path int true "product ID"
// @Produce 				application/json
// @Tags 					Product
// @Success 				200 {object} IResponse.Product{}
// @Router 					/api/product/{ProductId} [get]
func (c *ProductHandler) GetProductBySlug(ctx *gin.Context) {
	var reqUrl IRequest.GetProductBySlug
	if err := ctx.ShouldBindUri(&reqUrl); err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		response.ErrorResponse(ctx, "Body is invalid or not exist", 400)
		return
	}

	var reqViewed IRequest.GetParamsIsViewed
	if err := ctx.ShouldBindQuery(&reqViewed); err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		response.ErrorResponse(ctx, "Param is invalid or not exist", 400)
		return
	}

	Product, err, codeStatus := c.ProductUseCase.GetProductBySlugUseCase(ctx, reqUrl.Slug, reqViewed.IsViewed)

	if err != nil {
		response.ErrorResponse(ctx, err.Error(), codeStatus)
		return
	}

	global.Logger.Error("get Product", zap.String("Status", "Error"))
	response.SuccessResponse(ctx, "Get Product success", codeStatus, Product)

}

// // UpdateProduct 		godoc
// // @security 				BearerAuth
// // @Summary 				Update Product
// // @Description 			Update Product
// // @Param ProductId 	path int true "Update Product"
// // @Param 					tags body IRequest.GetBodyUpdateProduct true "Update Product"
// // @Produce 				application/json
// // @Tags 					Product
// // @Success 				200 {object} IResponse.Product{}
// // @Router 					/api/product-types/{ProductId} [put]
// func (c *ProductHandler) UpdateProduct(ctx *gin.Context) {
// 	var params IRequest.GetParamsUpdateProduct
// 	var body IRequest.GetBodyUpdateProduct
// 	if err := ctx.ShouldBindUri(&params); err != nil {
// 		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
// 		response.ErrorResponse(ctx, "Body is invalid or not exist", 400)
// 		return
// 	}

// 	if err := ctx.ShouldBindJSON(&body); err != nil {
// 		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
// 		response.ErrorResponse(ctx, "Body is invalid or not exist", 400)
// 		return
// 	}

// 	Product, err, codeStatus := c.ProductUseCase.UpdateProductUseCase(ctx, params.ID, body.Name, body.Slug)

// 	if err != nil {
// 		response.ErrorResponse(ctx, err.Error(), codeStatus)
// 		return
// 	}

// 	global.Logger.Error("get Product", zap.String("Status", "Error"))
// 	response.SuccessResponse(ctx, "Get Product success", codeStatus, Product)
// }

// DeleteProduct 		godoc
// @security 				BearerAuth
// @Summary 				Delete Product
// @Description 			Delete Product
// @Param ProductId 	path int true "Delete Product"
// @Produce 				application/json
// @Tags 					Product
// @Success 				200 {string} string [delete Product success]
// @Router 					/api/product-types/{ProductId} [delete]
func (c *ProductHandler) DeleteProduct(ctx *gin.Context) {
	var req IRequest.DeleteProduct
	if err := ctx.ShouldBindUri(&req); err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		response.ErrorResponse(ctx, "Body is invalid or not exist", 400)
		return
	}

	err, codeStatus := c.ProductUseCase.DeleteProductUseCase(ctx, req.ID)

	if err != nil {
		response.ErrorResponse(ctx, err.Error(), codeStatus)
		return
	}

	global.Logger.Error("get Product", zap.String("Status", "Error"))
	response.SuccessResponse(ctx, "Delete Product success", codeStatus, "")
}

// DeleteManyProduct 	godoc
// @security 				BearerAuth
// @Summary 				Delete many Product
// @Description 			Delete many Product
// @Param 					tags body IRequest.DeleteManyProduct true "DeleteMany Product"
// @Produce 				application/json
// @Tags 					Product
// @Success 				200 {string} string "Delete many Product success"
// @Router 					/api/product-types/delete-many [delete]
func (c *ProductHandler) DeleteManyProduct(ctx *gin.Context) {
	var req IRequest.DeleteManyProduct
	if err := ctx.ShouldBindJSON(&req); err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		response.ErrorResponse(ctx, "Body is invalid or not exist", 400)
		return
	}

	err, codeStatus := c.ProductUseCase.DeleteManyProductUseCase(ctx, req.ArrayId)

	if err != nil {
		response.ErrorResponse(ctx, err.Error(), codeStatus)
		return
	}

	global.Logger.Error("get Product", zap.String("Status", "Error"))
	response.SuccessResponse(ctx, "Delete Product success", codeStatus, "")
}
