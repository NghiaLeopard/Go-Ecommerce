package IUseCase

import (
	IRequest "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/handler/request"
	IResponse "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/handler/response"
	"github.com/gin-gonic/gin"
)

type Product interface {
	CreateProduct(ctx *gin.Context, req IRequest.CreateProduct) (IResponse.Product, error, int)
	GetProductUseCase(ctx *gin.Context, id int64) (IResponse.GetProduct, error, int)
	GetProductBySlugUseCase(ctx *gin.Context, id int64,isViewed bool) (IResponse.GetProduct, error, int)
	// GetAllProductUseCase(ctx *gin.Context, req IRequest.GetAllProduct) ([]db.Product, error, int)
	// UpdateProductUseCase(ctx *gin.Context, id int, name string, slug string) (IResponse.Product, error, int)
	DeleteProductUseCase(ctx *gin.Context, id int64) (error, int)
	DeleteManyProductUseCase(ctx *gin.Context, id []int64) (error, int)
}
