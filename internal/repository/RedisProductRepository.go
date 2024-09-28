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

func (r *RedisProductRepository) SetViewProduct(ctx *gin.Context, productId int64) error {
	key := fmt.Sprintf("%d:ViewProduct", productId)

	if err := r.Redis.Set(ctx, key, 0, 0).Err(); err != nil {
		global.Logger.Error("could not set view product to redis", zap.Error(err))
		return fmt.Errorf("could not set view product to redis for %v", err)
	}

	return nil
}

func (r *RedisProductRepository) IncViewProduct(ctx *gin.Context, productId int64) (int64, error) {
	key := fmt.Sprintf("%d:ViewProduct", productId)

	result := r.Redis.Incr(ctx, key)
	if result.Err() != nil {
		global.Logger.Error("could not set view product to redis", zap.Error(result.Err()))
		return 0, fmt.Errorf("could not set view product to redis %v", result.Err())
	}

	return result.Val(), nil
}

func (r *RedisProductRepository) SetProductLikedBy(ctx *gin.Context, productId int64, userID int64) error {
	key := fmt.Sprintf("%d:LikedBy", productId)

	if err := r.Redis.SAdd(ctx, key, userID).Err(); err != nil {
		global.Logger.Error("could not set likeBy product to redis", zap.Error(err))
		return fmt.Errorf("could not set likeBy product to redis for userID: %d: %v", userID, err)
	}

	return nil
}

func (r *RedisProductRepository) DeleteProductLikedBy(ctx *gin.Context, productId int64, userID int64) error {
	key := fmt.Sprintf("%d:LikedBy", productId)

	result := r.Redis.SRem(ctx, key, userID)

	if result.Err() != nil {
		global.Logger.Error("could not delete user likeBy product to redis", zap.Error(result.Err()))
		return fmt.Errorf("could not delete user likeBy product to redis : %d: %v", userID, result.Err())
	}

	if result.Val() < 1 {
		global.Logger.Error("likeBy product is not exist", zap.Error(result.Err()))
		return fmt.Errorf("likeBy product to redis for userID: %d does not exist", userID)
	}

	return nil
}

func (r *RedisProductRepository) SetProductUniqueView(ctx *gin.Context, productId int64, userID int64) error {
	key := fmt.Sprintf("%d:UniqueView", productId)

	if err := r.Redis.SAdd(ctx, key, userID).Err(); err != nil {
		global.Logger.Error("could not set UniqueView product to redis", zap.Error(err))
		return fmt.Errorf("could not set UniqueView product to redis for userID: %d: %v", userID, err)
	}

	return nil
}

func (r *RedisProductRepository) GetProductUniqueView(ctx *gin.Context, productId int64, userID int64) ([]string, error) {
	key := fmt.Sprintf("%d:UniqueView", productId)

	result := r.Redis.SMembers(ctx, key)

	if result.Err() != nil {
		global.Logger.Error("could not get user likeBy product to redis", zap.Error(result.Err()))
		return []string{}, fmt.Errorf("could not get user likeBy product to redis : %d: %v", userID, result.Err())
	}

	return result.Val(), nil
}

func (r *RedisProductRepository) GetTotalProductLikes(ctx *gin.Context, productId int64) (int64, error) {
	key := fmt.Sprintf("%d:LikedBy", productId)

	result := r.Redis.SCard(ctx, key)

	if result.Err() != nil {
		global.Logger.Error("could not get total like product to redis", zap.Error(result.Err()))
		return 0, fmt.Errorf("could not get total like product to redis %v", result.Err())
	}

	if result.Val() < 1 {
		return 0, nil
	}

	return result.Val(), nil
}

func (r *RedisProductRepository) SetUserView(ctx *gin.Context, productId int64, userID int64) error {
	key := fmt.Sprintf("%d:View", userID)

	if err := r.Redis.SAdd(ctx, key, productId).Err(); err != nil {
		global.Logger.Error("could not set UniqueView product to redis", zap.Error(err))
		return fmt.Errorf("could not set UniqueView product to redis for userID: %d: %v", userID, err)
	}

	return nil
}

func (r *RedisProductRepository) SetUserLikedBy(ctx *gin.Context, productId int64, userID int64) error {
	key := fmt.Sprintf("%d:LikedBy", userID)

	if err := r.Redis.SAdd(ctx, key, productId).Err(); err != nil {
		global.Logger.Error("could not set likeBy product to redis", zap.Error(err))
		return fmt.Errorf("could not set likeBy product to redis for userID: %d: %v", userID, err)
	}

	return nil
}

func (r *RedisProductRepository) DeleteUserLikedBy(ctx *gin.Context, productId int64, userID int64) error {
	key := fmt.Sprintf("%d:LikedBy", userID)

	result := r.Redis.SRem(ctx, key, productId)

	if result.Err() != nil {
		global.Logger.Error("could not delete user likeBy product to redis", zap.Error(result.Err()))
		return fmt.Errorf("could not delete user likeBy product to redis : %d: %v", userID, result.Err())
	}

	if result.Val() < 1 {
		global.Logger.Error("likeBy product is not exist", zap.Error(result.Err()))
		return fmt.Errorf("likeBy product to redis for userID: %d does not exist", userID)
	}

	return nil
}
