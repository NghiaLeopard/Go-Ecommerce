package handler

import (
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/global"
	IHandler "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/handler/interfaces"
	IRequest "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/handler/request"
	IResponse "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/handler/response"
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
// @Router 					/api/products [post]
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

	global.Logger.Info("create Product", zap.String("Status", "Success"))
	response.SuccessResponse(ctx, "Create Product success", codeStatus, Product)
}

// GetAllProductAdmin 		godoc
// @security 			BearerAuth
// @Summary 			Get all Product
// @Description 		Get all Product
// @Param 				request query IRequest.GetAllProductAdmin true "get all product admin"
// @Produce 			application/json
// @Tags 				Product
// @Success 			200 {array} []IResponse.GetAllProductAdmin{}
// @Router 				/api/products [get]
func (c *ProductHandler) GetAllProductAdmin(ctx *gin.Context) {
	var req IRequest.GetAllProductAdmin
	if err := ctx.ShouldBindQuery(&req); err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		response.ErrorResponse(ctx, "Body is invalid or not exist", 400)
		return
	}

	Product, err, codeStatus := c.ProductUseCase.GetAllProductAdminUseCase(ctx, req)

	if err != nil {
		response.ErrorResponse(ctx, err.Error(), codeStatus)
		return
	}

	global.Logger.Error("get Product", zap.String("Status", "Success"))
	response.SuccessResponse(ctx, "Get Product success", codeStatus, Product)
}

// GetAllProductPublic 		godoc
// @security 			BearerAuth
// @Summary 			Get all Product
// @Description 		Get all Product
// @Param 				request query IRequest.GetAllProductPublic true "get all product type"
// @Produce 			application/json
// @Tags 				Product
// @Success 			200 {array} []IResponse.GetAllProductPublic{}
// @Router 				/api/products/public [get]
func (c *ProductHandler) GetAllProductPublic(ctx *gin.Context) {
	var req IRequest.GetAllProductPublic
	if err := ctx.ShouldBindQuery(&req); err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		response.ErrorResponse(ctx, "Body is invalid or not exist", 400)
		return
	}

	Product, err, codeStatus := c.ProductUseCase.GetAllProductPublicUseCase(ctx, req)

	if err != nil {
		response.ErrorResponse(ctx, err.Error(), codeStatus)
		return
	}

	global.Logger.Error("get Product", zap.String("Status", "Success"))
	response.SuccessResponse(ctx, "Get Product success", codeStatus, Product)
}

// GetAllProductMeLiked 		godoc
// @security 					BearerAuth
// @Summary 					Get all Product
// @Description 				Get all Product
// @Param 						request query IRequest.GetAllProductLiked true "get all product me liked"
// @Produce 					application/json
// @Tags 						Product
// @Success 					200 {array} []IResponse.Product{}
// @Router 						/api/products/liked/me [get]
func (c *ProductHandler) GetAllProductMeLiked(ctx *gin.Context) {
	var req IRequest.GetAllProductLiked
	if err := ctx.ShouldBindQuery(&req); err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		response.ErrorResponse(ctx, "Body is invalid or not exist", 400)
		return
	}

	Product, err, codeStatus := c.ProductUseCase.GetAllProductMeLikedUseCase(ctx, req)

	if err != nil {
		response.ErrorResponse(ctx, err.Error(), codeStatus)
		return
	}

	global.Logger.Error("get Product", zap.String("Status", "success"))
	response.SuccessResponse(ctx, "Get Product success", codeStatus, Product)
}

// GetAllProductMeViewed 		godoc
// @security 					BearerAuth
// @Summary 					Get all Product me viewed
// @Description 				Get all Product me viewed
// @Param 						request query IRequest.GetAllProductViewed true "get all product type"
// @Produce 					application/json
// @Tags 						Product
// @Success 					200 {array} []IResponse.GetAllMeLiked{}
// @Router 						/api/products/viewed/me [get]
func (c *ProductHandler) GetAllProductMeViewed(ctx *gin.Context) {
	var req IRequest.GetAllProductViewed
	if err := ctx.ShouldBindQuery(&req); err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		response.ErrorResponse(ctx, "Body is invalid or not exist", 400)
		return
	}

	Product, err, codeStatus := c.ProductUseCase.GetAllProductMeViewedUseCase(ctx, req)

	if err != nil {
		response.ErrorResponse(ctx, err.Error(), codeStatus)
		return
	}

	global.Logger.Error("get Product", zap.String("Status", "success"))
	response.SuccessResponse(ctx, "Get Product success", codeStatus, Product)
}

// GetProduct 				godoc
// @security 				BearerAuth
// @Summary 				Get Product by id
// @Description 			Get Product by id
// @Param productId  		path int true "product ID"
// @Produce 				application/json
// @Tags 					Product
// @Success 				200 {object} IResponse.GetAllMeViewed{}
// @Router 					/api/products/{productId} [get]
func (c *ProductHandler) GetProduct(ctx *gin.Context) {
	var req IRequest.GetProduct
	var _ *IResponse.Product
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

	global.Logger.Info("get Product", zap.String("Status", "Success"))
	response.SuccessResponse(ctx, "Get Product success", codeStatus, Product)
}

// GetProductRelated 			godoc
// @security 					BearerAuth
// @Summary 					Get product related
// @Description 				Get product related
// @Param  						request query IRequest.GetAllProductRelated true "product related"
// @Produce 					application/json
// @Tags 						Product
// @Success 					200 {object} IResponse.GetAllProductRelated{}
// @Router 						/api/products/related [get]
func (c *ProductHandler) GetProductRelated(ctx *gin.Context) {
	var req IRequest.GetAllProductRelated
	if err := ctx.ShouldBindQuery(&req); err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		response.ErrorResponse(ctx, "Body is invalid or not exist", 400)
		return
	}

	Product, err, codeStatus := c.ProductUseCase.GetProductRelatedUseCase(ctx, req)

	if err != nil {
		response.ErrorResponse(ctx, err.Error(), codeStatus)
		return
	}

	global.Logger.Info("get Product", zap.String("Status", "Success"))
	response.SuccessResponse(ctx, "Get Product success", codeStatus, Product)
}

// GetProductBySlug 		godoc
// @security 				BearerAuth
// @Summary 				Get Product by slug
// @Description 			Get Product by slug
// @Param productSlug  		path string true "product slug"
// @Param 					request query IRequest.GetParamsIsViewed true "is viewed product"
// @Produce 				application/json
// @Tags 					Product
// @Success 				200 {object} IResponse.Product{}
// @Router 					/api/products/public/slug/{productSlug} [get]
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

	global.Logger.Info("get Product", zap.String("Status", "Success"))
	response.SuccessResponse(ctx, "Get Product success", codeStatus, Product)
}

// GetProductPublicById 	godoc
// @security 				BearerAuth
// @Summary 				Get Product public by id
// @Description 			Get Product public by id
// @Param productId  		path int true "product ID"
// @Param 					request query IRequest.GetParamsIsViewed true "is viewed product"
// @Produce 				application/json
// @Tags 					Product
// @Success 				200 {object} IResponse.Product{}
// @Router 					/api/products/public/{productId} [get]
func (c *ProductHandler) GetProductPublicById(ctx *gin.Context) {
	var reqUrl IRequest.GetProductPublicById
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

	Product, err, codeStatus := c.ProductUseCase.GetProductPublicByIdUseCase(ctx, reqUrl.ID, reqViewed.IsViewed)

	if err != nil {
		response.ErrorResponse(ctx, err.Error(), codeStatus)
		return
	}

	global.Logger.Info("get Product", zap.String("Status", "Success"))
	response.SuccessResponse(ctx, "Get Product success", codeStatus, Product)

}

// UpdateProduct 			godoc
// @security 				BearerAuth
// @Summary 				Update Product
// @Description 			Update Product
// @Param productId 		path int true "Update Product"
// @Param 					tags body IRequest.UpdateProduct true "Update Product"
// @Produce 				application/json
// @Tags 					Product
// @Success 				200 {object} IResponse.UpdateProduct{}
// @Router 					/api/products/{productId} [put]
func (c *ProductHandler) UpdateProduct(ctx *gin.Context) {
	var params IRequest.UpdateProductUrl
	var body IRequest.UpdateProduct
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

	Product, err, codeStatus := c.ProductUseCase.UpdateProductUseCase(ctx, params.ID, body)

	if err != nil {
		response.ErrorResponse(ctx, err.Error(), codeStatus)
		return
	}

	global.Logger.Error("get Product", zap.String("Status", "Error"))
	response.SuccessResponse(ctx, "Get Product success", codeStatus, Product)
}

// DeleteProduct 			godoc
// @security 				BearerAuth
// @Summary 				Delete Product
// @Description 			Delete Product
// @Param productId 		path int true "Delete Product"
// @Produce 				application/json
// @Tags 					Product
// @Success 				200 {string} string [delete Product success]
// @Router 					/api/products/{productId} [delete]
func (c *ProductHandler) DeleteProduct(ctx *gin.Context) {
	var req IRequest.DeleteProduct
	if err := ctx.ShouldBindUri(&req); err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "error"))
		response.ErrorResponse(ctx, "Body is invalid or not exist", 400)
		return
	}

	err, codeStatus := c.ProductUseCase.DeleteProductUseCase(ctx, req.ID)

	if err != nil {
		response.ErrorResponse(ctx, err.Error(), codeStatus)
		return
	}

	global.Logger.Info("get Product", zap.String("Status", "Success"))
	response.SuccessResponse(ctx, "Delete Product success", codeStatus, map[string]int{"_id": 1})
}

// DeleteManyProduct 		godoc
// @security 				BearerAuth
// @Summary 				Delete many Product
// @Description 			Delete many Product
// @Param 					tags body IRequest.DeleteManyProduct true "DeleteMany Product"
// @Produce 				application/json
// @Tags 					Product
// @Success 				200 {string} string "Delete many Product success"
// @Router 					/api/products/delete-many [delete]
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

	global.Logger.Info("get Product", zap.String("Status", "Success"))
	response.SuccessResponse(ctx, "Delete Product success", codeStatus, map[string]int{"_id": 1})
}

// LikeProduct 				godoc
// @security 				BearerAuth
// @Summary 				Like product
// @Description 			Like product
// @Param 					tags body IRequest.LikeProduct true "like product"
// @Produce 				application/json
// @Tags 					Product
// @Success 				200 {string} string "Like product success"
// @Router 					/api/products/like [post]
func (c *ProductHandler) LikeProduct(ctx *gin.Context) {
	var req IRequest.LikeProduct
	if err := ctx.ShouldBindJSON(&req); err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		response.ErrorResponse(ctx, "Body is invalid or not exist", 400)
		return
	}

	err, codeStatus := c.ProductUseCase.LikeProductUseCase(ctx, req.ID)

	if err != nil {
		response.ErrorResponse(ctx, err.Error(), codeStatus)
		return
	}

	global.Logger.Info("get Product", zap.String("Status", "Success"))
	response.SuccessResponse(ctx, "Like Product success", codeStatus, "")
}

// UnLikeProduct 			godoc
// @security 				BearerAuth
// @Summary 				Unlike product
// @Description 			Unlike product
// @Param 					tags body IRequest.UnLikeProduct true "unlike product"
// @Produce 				application/json
// @Tags 					Product
// @Success 				200 {string} string "Unlike product success"
// @Router 					/api/products/unlike [post]
func (c *ProductHandler) UnLikeProduct(ctx *gin.Context) {
	var req IRequest.UnLikeProduct
	if err := ctx.ShouldBindJSON(&req); err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		response.ErrorResponse(ctx, "Body is invalid or not exist", 400)
		return
	}

	err, codeStatus := c.ProductUseCase.UnLikeProductUseCase(ctx, req.ID)

	if err != nil {
		response.ErrorResponse(ctx, err.Error(), codeStatus)
		return
	}

	global.Logger.Info("get Product", zap.String("Status", "Success"))
	response.SuccessResponse(ctx, "Unlike Product success", codeStatus, "")
}
