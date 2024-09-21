package repository

import (
	"fmt"
	"time"

	"github.com/NghiaLeopard/Go-Ecommerce-Backend/global"
	IRepository "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/repository/interface"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type RedisTokenRepository struct {
	Redis *redis.Client
}

func NewRedisTokenRepository(redis *redis.Client) IRepository.RedisToken {
	return &RedisTokenRepository{
		Redis: redis,
	}
}

func (r *RedisTokenRepository) SetRefreshToken(ctx *gin.Context, userID int64, tokenID string, expiration time.Duration) error {
	key := fmt.Sprintf("%d:refreshToken:%s", userID, tokenID)
	fmt.Println(key)

	if err := r.Redis.Set(ctx, key, 0, expiration).Err(); err != nil {
		global.Logger.Error("could not set refresh token to redis", zap.Error(err))
		return fmt.Errorf("could not set refresh token to redis for userID: %d: %v", userID, err)
	}

	return nil
}

func (r *RedisTokenRepository) DeleteRefreshToken(ctx *gin.Context, userID int64, tokenID string) error {
	key := fmt.Sprintf("%d:refreshToken:%s", userID, tokenID)

	result := r.Redis.Del(ctx, key)

	if result.Err() != nil {
		global.Logger.Error("could not delete refresh token to redis", zap.Error(result.Err()))
		return fmt.Errorf("could not delete refresh token to redis for userID: %d: %v", userID, result.Err())
	}

	if result.Val() < 1 {
		global.Logger.Error("refresh token is not exist", zap.Error(result.Err()))
		return fmt.Errorf("refresh token to redis for userID/tokenID: %d/%s does not exist", userID, tokenID)
	}

	return nil
}

func (r *RedisTokenRepository) CheckRefreshToken(ctx *gin.Context, userID int64, tokenID string) error {
	key := fmt.Sprintf("%d:refreshToken:%s", userID, tokenID)

	result := r.Redis.Exists(ctx, key)

	if result.Err() != nil {
		global.Logger.Error("could not find refresh token to redis", zap.Error(result.Err()))
		return fmt.Errorf("could not find refresh token to redis for userID: %d: %v", userID, result.Err())
	}

	if result.Val() < 1 {
		global.Logger.Error("refresh token is not exist", zap.Error(result.Err()))
		return fmt.Errorf("refresh token to redis for userID/tokenID: %d/%s does not exist", userID, tokenID)
	}

	return nil
}
