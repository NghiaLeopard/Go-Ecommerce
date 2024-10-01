package IRepository

import (
	"sync"

	db "github.com/NghiaLeopard/Go-Ecommerce-Backend/db/sqlc"
	IRequest "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/handler/request"
	"github.com/gin-gonic/gin"
)

type Product interface {
	CreateProductNotDiscount(ctx *gin.Context, req IRequest.CreateProduct) (db.Product, error)
	CreateProductDiscount(ctx *gin.Context, req IRequest.CreateProduct) (db.Product, error)
	// GetAllProduct(ctx *gin.Context, req IRequest.GetAllProduct) ([]db.Product, error)
	GetAllProductMeLiked(ctx *gin.Context, req IRequest.GetAllProductLiked, userId int) ([]db.GetAllProductLikeRow, error)
	GetProductById(ctx *gin.Context, id int64) (db.GetProductByIdRow, error)
	GetProductBySlug(ctx *gin.Context, slug string) (db.GetProductBySlugRow, error)
	GetProductPublicById(ctx *gin.Context, productId int64) (db.GetProductPublicByIdRow, error)
	DeleteProduct(ctx *gin.Context, userId int64) error
	DeleteManyProduct(ctx *gin.Context, arrayId []int64) error

	// View product
	UpdateViewProduct(ctx *gin.Context, id int64, view int32, wg *sync.WaitGroup)
	UpdateUniqueView(ctx *gin.Context, productId int64, userId int, wg *sync.WaitGroup)

	// Liked product
	UpdateLikeProduct(ctx *gin.Context, productId int64, userId int) error
	DeleteLikeProduct(ctx *gin.Context, userId int) error
}
