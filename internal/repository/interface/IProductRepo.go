package IRepository

import (
	IRequest "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/handler/request"
	db "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Product interface {
	CreateProductNotDiscount(ctx *gin.Context, req IRequest.CreateProduct) (db.Product, error)
	CreateProductDiscount(ctx *gin.Context, req IRequest.CreateProduct) (db.Product, error)
	// GetAllProduct(ctx *gin.Context, req IRequest.GetAllProduct) ([]db.Product, error)
	GetProductById(ctx *gin.Context, id int64) (db.GetProductByIdRow, error)
	// GetProductByName(ctx *gin.Context, name string) (db.Product, error)
	// UpdateProduct(ctx *gin.Context, id int64, name string, slug string) (db.Product, error)
	DeleteProduct(ctx *gin.Context, id int64) error
	DeleteManyProduct(ctx *gin.Context, arrayId []int64) error
}
