package IHandler

import "github.com/gin-gonic/gin"

type Product interface {
	CreateProduct(ctx *gin.Context)
	GetProduct(ctx *gin.Context)
	GetProductRelated(ctx *gin.Context)
	GetProductPublicById(ctx *gin.Context)
	GetProductBySlug(ctx *gin.Context)
	GetAllProductMeLiked(ctx *gin.Context)
	GetAllProductMeViewed(ctx *gin.Context)
	GetAllProductAdmin(ctx *gin.Context)
	GetAllProductPublic(ctx *gin.Context)
	UpdateProduct(ctx *gin.Context)
	DeleteProduct(ctx *gin.Context)
	DeleteManyProduct(ctx *gin.Context)
	LikeProduct(ctx *gin.Context)
	UnLikeProduct(ctx *gin.Context)
}
