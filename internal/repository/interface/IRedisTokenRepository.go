package IRepository

import (
	"time"

	"github.com/gin-gonic/gin"
)

type RedisToken interface {
	SetRefreshToken(ctx *gin.Context, userId int64, token string, expiration time.Duration) error
	DeleteRefreshToken(ctx *gin.Context, userId int64, token string) error
	CheckRefreshToken(ctx *gin.Context, userId int64, token string) error
}
