package IUseCase

import (
	IResponse "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/handler/response"
	"github.com/gin-gonic/gin"
)

type ProductType interface {
	CreateProductType(ctx *gin.Context, name string, slug string) (IResponse.ProductType, error, int)
	GetProductTypeUseCase(ctx *gin.Context, id int) (IResponse.ProductType, error, int)
	UpdateProductTypeUseCase(ctx *gin.Context, id int, name string, slug string) (IResponse.ProductType, error, int)
	DeleteProductTypeUseCase(ctx *gin.Context, id int) (error, int)
	DeleteManyProductTypeUseCase(ctx *gin.Context, id []int) (error, int)
}
