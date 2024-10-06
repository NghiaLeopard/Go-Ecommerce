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

	// Get all product me action
	GetAllProductMeLiked(ctx *gin.Context, req IRequest.GetAllProductLiked, userId int) ([]db.GetAllProductLikeRow, error)
	GetAllProductMeViewed(ctx *gin.Context, req IRequest.GetAllProductViewed, userId int) ([]db.GetAllProductViewRow, error)

	// Get all
	GetAllProductAdmin(ctx *gin.Context, req IRequest.GetAllProductAdmin) ([]db.GetAllProductAdminRow, error)
	GetAllProductPublic(ctx *gin.Context, req IRequest.GetAllProductPublic) ([]db.GetAllProductPublicRow, error)

	// All product related
	GetAllProductRelated(ctx *gin.Context, req IRequest.GetAllProductRelated, id int64, city int32) ([]db.GetAllProductRelatedRow, error)

	// Get product
	GetProductById(ctx *gin.Context, id int64) (db.GetProductByIdRow, error)
	GetProductBySlug(ctx *gin.Context, slug string) (db.GetProductBySlugRow, error)
	GetProductPublicById(ctx *gin.Context, productId int64) (db.GetProductPublicByIdRow, error)
	GetProductTypeBySlug(ctx *gin.Context, slug string) (db.GetProductTypeBySlugRow, error)

	// Update
	UpdateProduct(ctx *gin.Context, id int64, body IRequest.UpdateProduct) (db.Product, error)

	// Delete
	DeleteProduct(ctx *gin.Context, userId int64) error
	DeleteManyProduct(ctx *gin.Context, arrayId []int64) error

	// View product
	UpdateViewProduct(ctx *gin.Context, id int64, view int32, wg *sync.WaitGroup)
	UpdateUniqueView(ctx *gin.Context, productId int64, userId int, wg *sync.WaitGroup)

	// Liked product
	UpdateLikeProduct(ctx *gin.Context, productId int64, userId int) error
	DeleteLikeProduct(ctx *gin.Context, userId int) error

	CheckProduct(ctx *gin.Context, id int64) error
}
