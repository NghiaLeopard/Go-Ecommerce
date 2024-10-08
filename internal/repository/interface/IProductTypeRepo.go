package IRepository

import (
	db "github.com/NghiaLeopard/Go-Ecommerce-Backend/db/sqlc"
	IRequest "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/handler/request"
	"github.com/gin-gonic/gin"
)

type ProductType interface {
	CreateProductType(ctx *gin.Context, name string, slug string) (db.ProductType, error)
	GetAllProductType(ctx *gin.Context, req IRequest.GetAllProductType) ([]db.ListProductTypeRow, error)
	GetProductTypeById(ctx *gin.Context, id int64) (db.ProductType, error)
	GetProductTypeByName(ctx *gin.Context, name string) (db.ProductType, error)
	UpdateProductType(ctx *gin.Context, id int64, name string, slug string) (db.ProductType, error)
	DeleteProductType(ctx *gin.Context, id int64) error
	DeleteManyProductType(ctx *gin.Context, arrayId []int64) error
}
