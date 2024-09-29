package IRepository

import (
	"github.com/gin-gonic/gin"
)

type RedisProduct interface {
	SetProductUniqueView(ctx *gin.Context, productId int64, userID int) error
	CheckProductUniqueView(ctx *gin.Context, productId int64, userID int) (bool, error)
	SetLikeProduct(ctx *gin.Context, productId int64, userID int) error
	CheckLikeProduct(ctx *gin.Context, productId int64, userID int) (bool, error)
	DeleteLikeProduct(ctx *gin.Context, productId int64, userID int) error
}
