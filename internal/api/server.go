package api

import (
	IHandler "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/handler/interfaces"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/middleware"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/routers"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/pkg/config"
	"github.com/gin-gonic/gin"
)

type ServerHTTP struct {
	Engine *gin.Engine
	config config.Config
}

func NewServerHTTP(config config.Config, middleware middleware.Middleware, authHandler IHandler.Auth, cityHandler IHandler.City) *ServerHTTP {
	engine := gin.Default()

	api := engine.Group("/api")

	routers.UserRouter(api, middleware, authHandler)
	routers.CityRouter(api, middleware, cityHandler)
	// routers.RoleRouter(api, middleware, roleHandler)

	return &ServerHTTP{Engine: engine, config: config}
}

func (s *ServerHTTP) Start() error {
	return s.Engine.Run(s.config.ServerAction)
}
