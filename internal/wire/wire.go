//go:build wireinject
// +build wireinject

package wire

import (
	"database/sql"

	"github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/handler"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/middleware"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/usecase"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/pkg/config"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/pkg/token"
	"github.com/google/wire"
)

func InitApi(db *sql.DB, config config.Config, token token.Maker) (*api.ServerHTTP, error) {
	wire.Build(
		middleware.NewMiddleware,
		// use case
		usecase.NewAuthUseCase,

		// handler
		handler.NewAuthHandler,

		api.NewServerHTTP,
	)

	return &api.ServerHTTP{}, nil
}
