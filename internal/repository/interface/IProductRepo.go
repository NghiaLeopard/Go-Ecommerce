package IRepository

import (
	"sync"

	IRequest "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/handler/request"
	db "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Product interface {
	CreateProductNotDiscount(ctx *gin.Context, req IRequest.CreateProduct) (db.Product, error)
	CreateProductDiscount(ctx *gin.Context, req IRequest.CreateProduct) (db.Product, error)
	// GetAllProduct(ctx *gin.Context, req IRequest.GetAllProduct) ([]db.Product, error)
	GetProductById(ctx *gin.Context, id int64) (db.GetProductByIdRow, error)
	GetProductBySlug(ctx *gin.Context, slug string) (db.GetProductBySlugRow, error)
	UpdateViewProduct(ctx *gin.Context, id int64, view int32, wg *sync.WaitGroup)
	UpdateUniqueView(ctx *gin.Context, productId int64, userId int, wg *sync.WaitGroup)
	DeleteProduct(ctx *gin.Context, id int64) error
	DeleteManyProduct(ctx *gin.Context, arrayId []int64) error
}
