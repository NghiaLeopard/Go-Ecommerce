package api

import (
	_ "github.com/NghiaLeopard/Go-Ecommerce-Backend/docs"
	IHandler "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/handler/interfaces"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/middleware"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/routers"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/pkg/config"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type ServerHTTP struct {
	Engine *gin.Engine
	config config.Config
}

func NewServerHTTP(config config.Config, middleware middleware.Middleware, authHandler IHandler.Auth, cityHandler IHandler.City, roleHandler IHandler.Role) *ServerHTTP {
	engine := gin.Default()

	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := engine.Group("/api")

	routers.UserRouter(api, middleware, authHandler)
	routers.CityRouter(api, middleware, cityHandler)
	routers.RoleRouter(api, middleware, roleHandler)

	return &ServerHTTP{Engine: engine, config: config}
}

func (s *ServerHTTP) Start() error {
	return s.Engine.Run(s.config.ServerAction)
}
