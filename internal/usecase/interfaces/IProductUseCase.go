package IUseCase

import (
	IRequest "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/handler/request"
	IResponse "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/handler/response"
	"github.com/gin-gonic/gin"
)

type Product interface {
	CreateProduct(ctx *gin.Context, req IRequest.CreateProduct) (IResponse.Product, error, int)
	GetProductUseCase(ctx *gin.Context, id int64) (IResponse.GetProduct, error, int)
	GetProductRelatedUseCase(ctx *gin.Context, req IRequest.GetAllProductRelated) (IResponse.GetAllProductRelated, error, int)

	// Public
	GetProductBySlugUseCase(ctx *gin.Context, slug string, isViewed bool) (IResponse.GetProduct, error, int)
	GetProductPublicByIdUseCase(ctx *gin.Context, productId int64, isViewed bool) (IResponse.GetProduct, error, int)

	// Me Action
	GetAllProductMeLikedUseCase(ctx *gin.Context, req IRequest.GetAllProductLiked) (IResponse.GetAllMeLiked, error, int)
	GetAllProductMeViewedUseCase(ctx *gin.Context, req IRequest.GetAllProductViewed) (IResponse.GetAllMeViewed, error, int)

	// Get all product
	GetAllProductAdminUseCase(ctx *gin.Context, req IRequest.GetAllProductAdmin) (IResponse.GetAllProductAdmin, error, int)
	GetAllProductPublicUseCase(ctx *gin.Context, req IRequest.GetAllProductPublic) (IResponse.GetAllProductPublic, error, int)

	// Update
	UpdateProductUseCase(ctx *gin.Context, id int64, body IRequest.UpdateProduct) (IResponse.UpdateProduct, error, int)

	// Delete
	DeleteProductUseCase(ctx *gin.Context, id int64) (error, int)
	DeleteManyProductUseCase(ctx *gin.Context, id []int64) (error, int)

	// Like
	LikeProductUseCase(ctx *gin.Context, id int64) (error, int)
	UnLikeProductUseCase(ctx *gin.Context, id int64) (error, int)
}
