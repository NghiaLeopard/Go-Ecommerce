package IRepository

import (
	"github.com/gin-gonic/gin"
)

type RedisProduct interface {
	SetProductUniqueView(ctx *gin.Context, productId int64, userID int64) error
	DeleteProductLikedBy(ctx *gin.Context, productId int64, userID int64) error
	SetProductLikedBy(ctx *gin.Context, productId int64, userID int64) error
	IncViewProduct(ctx *gin.Context, productId int64) (int64, error)
	SetViewProduct(ctx *gin.Context, productId int64) error
	GetTotalProductLikes(ctx *gin.Context, productId int64) (int64, error)
}
