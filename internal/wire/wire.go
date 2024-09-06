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

		// use case
		usecase.NewAuthUseCase,

		// handler
		handler.NewAuthHandler,

		api.NewServerHTTP,
	)

	return &api.ServerHTTP{}, nil
}
