package IRepository

import (
	"time"

	"github.com/gin-gonic/gin"
)

type RedisToken interface {
	SetRefreshToken(ctx *gin.Context, userId int64, token string, expiration time.Duration) error
	DeleteRefreshToken(ctx *gin.Context, userId int64) error
	CheckRefreshToken(ctx *gin.Context, userId int64, token string) error
	BlackListToken(ctx *gin.Context, accessToken string) error
	CheckBlackListToken(ctx *gin.Context, accessToken string) error
	SetResetToken(ctx *gin.Context, userID int64, token string, expiration time.Duration) error
	CheckResetToken(ctx *gin.Context, userID int64, tokenID string) error
}
