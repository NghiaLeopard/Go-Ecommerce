package repository

import (
	"fmt"

	"github.com/NghiaLeopard/Go-Ecommerce-Backend/global"
	IRepository "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/repository/interface"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type RedisProductRepository struct {
	Redis *redis.Client
}

func NewRedisProductRepository(redis *redis.Client) IRepository.RedisProduct {
	return &RedisProductRepository{
		Redis: redis,
	}
}

func (r *RedisProductRepository) SetProductUniqueView(ctx *gin.Context, productId int64, userID int) error {
	key := fmt.Sprintf("%d:UniqueView", productId)

	if err := r.Redis.SAdd(ctx, key, userID).Err(); err != nil {
		global.Logger.Error("could not set UniqueView product to redis", zap.Error(err))
		return fmt.Errorf("could not set UniqueView product to redis for userID: %d: %v", userID, err)
	}

	return nil
}

func (r *RedisProductRepository) CheckProductUniqueView(ctx *gin.Context, productId int64, userID int) (bool, error) {
	key := fmt.Sprintf("%d:UniqueView", productId)

	result := r.Redis.SIsMember(ctx, key, userID)

	if result.Err() != nil {
		global.Logger.Error("could not get UniqueView product to redis", zap.Error(result.Err()))
		return true, fmt.Errorf("could not set UniqueView product to redis for userID: %d: %v", userID, result.Err())
	}

	return result.Val(), nil
}
