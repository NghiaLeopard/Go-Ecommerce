//go:build wireinject
// +build wireinject

package wire

import (
	db "github.com/NghiaLeopard/Go-Ecommerce-Backend/db/sqlc"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/handler"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/middleware"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/repository"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/usecase"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/pkg/config"
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
)

func InitServer(sqlcDB *db.Queries, config config.Config, redis *redis.Client) (*api.ServerHTTP, error) {
	wire.Build(
		// middleware
		repository.NewRedisTokenRepository,
		repository.NewRedisProductRepository,
		middleware.NewMiddleware,

		// repository
		repository.NewAuthRepository,
		repository.NewCityRepository,
		repository.NewUserRepository,
		repository.NewDeliveryRepository,
		repository.NewPaymentRepository,
		repository.NewRoleRepository,
		repository.NewProductTypeRepository,
		repository.NewProductRepository,

		// use case
		usecase.NewAuthUseCase,
		usecase.NewCityUseCase,
		usecase.NewUserUseCase,
		usecase.NewDeliveryUseCase,
		usecase.NewPaymentUseCase,
		usecase.NewRoleUseCase,
		usecase.NewProductTypeUseCase,
		usecase.NewProductUseCase,

		// handler
		handler.NewAuthHandler,
		handler.NewDeliveryHandler,
		handler.NewPaymentHandler,
		handler.NewCityHandler,
		handler.NewUserHandler,
		handler.NewRoleHandler,
		handler.NewProductTypeHandler,
		handler.NewProductHandler,

		api.NewServerHTTP,
	)

	return &api.ServerHTTP{}, nil
}
