package IHandler

import "github.com/gin-gonic/gin"

type ProductType interface {
	CreateProductType(ctx *gin.Context)
	GetAllProductType(ctx *gin.Context)
	GetProductType(ctx *gin.Context)
	UpdateProductType(ctx *gin.Context)
	DeleteProductType(ctx *gin.Context)
	DeleteManyProductType(ctx *gin.Context)
}
