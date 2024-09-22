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

func (r *RedisTokenRepository) SetRefreshToken(ctx *gin.Context, userID int64, token string, expiration time.Duration) error {
	key := fmt.Sprintf("%d:refreshToken", userID)

	if err := r.Redis.Set(ctx, key, token, expiration).Err(); err != nil {
		global.Logger.Error("could not set refresh token to redis", zap.Error(err))
		return fmt.Errorf("could not set refresh token to redis for userID: %d: %v", userID, err)
	}

	return nil
}

func (r *RedisTokenRepository) DeleteRefreshToken(ctx *gin.Context, userID int64) error {
	key := fmt.Sprintf("%d:refreshToken", userID)

	result := r.Redis.Del(ctx, key)

	if result.Err() != nil {
		global.Logger.Error("could not delete refresh token to redis", zap.Error(result.Err()))
		return fmt.Errorf("could not delete refresh token to redis for userID: %d: %v", userID, result.Err())
	}

	if result.Val() < 1 {
		global.Logger.Error("refresh token is not exist", zap.Error(result.Err()))
		return fmt.Errorf("refresh token to redis for userID/tokenID: %d does not exist", userID)
	}

	return nil
}

func (r *RedisTokenRepository) CheckRefreshToken(ctx *gin.Context, userID int64, tokenID string) error {
	key := fmt.Sprintf("%d:refreshToken", userID)

	result := r.Redis.Get(ctx, key)

	if result.Err() != nil {
		global.Logger.Error("could not find refresh token to redis", zap.Error(result.Err()))
		return fmt.Errorf("could not find refresh token to redis for userID: %d: %v", userID, result.Err())
	}

	if result.Val() != tokenID {
		global.Logger.Error("refresh token is invalid", zap.Error(result.Err()))
		return fmt.Errorf("refresh token is invalid")
	}

	return nil
}

func (r *RedisTokenRepository) SetResetToken(ctx *gin.Context, userID int64, token string, expiration time.Duration) error {
	key := fmt.Sprintf("%d:resetToken", userID)

	if err := r.Redis.Set(ctx, key, token, expiration).Err(); err != nil {
		global.Logger.Error("could not set reset token to redis", zap.Error(err))
		return fmt.Errorf("could not set reset token to redis for userID: %d: %v", userID, err)
	}

	return nil
}

func (r *RedisTokenRepository) CheckResetToken(ctx *gin.Context, userID int64, tokenID string) error {
	key := fmt.Sprintf("%d:resetToken", userID)

	result := r.Redis.Get(ctx, key)

	if result.Err() != nil {
		global.Logger.Error("token is expire", zap.Error(result.Err()))
		return fmt.Errorf("token is expire")
	}

	if result.Val() != tokenID {
		global.Logger.Error("reset token is invalid", zap.Error(result.Err()))
		return fmt.Errorf("reset token is invalid")
	}

	return nil
}

func (r *RedisTokenRepository) BlackListToken(ctx *gin.Context, accessToken string) error {
	if err := r.Redis.SAdd(ctx, "blacklist", accessToken).Err(); err != nil {
		global.Logger.Error("could not set access token to blacklist redis", zap.Error(err))
		return fmt.Errorf("could not find refresh token to redis for userID: %d", err)
	}
	return nil
}

func (r *RedisTokenRepository) CheckBlackListToken(ctx *gin.Context, accessToken string) error {
	result := r.Redis.SIsMember(ctx, "blacklist", accessToken)

	if result.Err() != nil {
		global.Logger.Error("could not check blacklist redis", zap.Error(result.Err()))
		return fmt.Errorf("could not check blacklist redis for userID: %d", result.Err())
	}

	if result.Val() {
		global.Logger.Error("access token have blacklist", zap.Error(result.Err()))
		return fmt.Errorf("access token have blacklist")
	}

	fmt.Println(result.Val())

	return nil
}
