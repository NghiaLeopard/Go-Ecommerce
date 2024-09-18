//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/handler"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/middleware"
	db "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/db/sqlc"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/repository"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/usecase"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/pkg/config"
	"github.com/google/wire"
)

func InitServer(sqlcDB *db.Queries, config config.Config) (*api.ServerHTTP, error) {
	wire.Build(
		// middleware
		middleware.NewMiddleware,

		// repository
		repository.NewAuthRepository,
		repository.NewCityRepository,
		repository.NewRoleRepository,
		repository.NewProductTypeRepository,

		// use case
		usecase.NewAuthUseCase,
		usecase.NewCityUseCase,
		usecase.NewRoleUseCase,
		usecase.NewProductTypeUseCase,

		// handler
		handler.NewAuthHandler,
		handler.NewCityHandler,
		handler.NewRoleHandler,
		handler.NewProductTypeHandler,

		api.NewServerHTTP,
	)

	return &api.ServerHTTP{}, nil
}
