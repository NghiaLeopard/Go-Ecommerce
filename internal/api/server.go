package api

import (
	_ "github.com/NghiaLeopard/Go-Ecommerce-Backend/docs"
	IHandler "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/handler/interfaces"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/middleware"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/routers"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/pkg/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type ServerHTTP struct {
	Engine *gin.Engine
	config config.Config
}

func NewServerHTTP(config config.Config, middleware middleware.Middleware, authHandler IHandler.Auth, userHandler IHandler.User, productHandler IHandler.Product, cityHandler IHandler.City, DeliveryHandler IHandler.Delivery, PaymentHandler IHandler.Payment, roleHandler IHandler.Role, productTypeHandler IHandler.ProductType) *ServerHTTP {
	engine := gin.Default()

	engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// engine.Use(cors.Default())

	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := engine.Group("/api")

	routers.AuthRouter(api, middleware, authHandler)
	routers.UserRouter(api, middleware, userHandler)
	routers.CityRouter(api, middleware, cityHandler)
	routers.DeliveryRouter(api, middleware, DeliveryHandler)
	routers.PaymentRouter(api, middleware, PaymentHandler)
	routers.RoleRouter(api, middleware, roleHandler)
	routers.ProductRouter(api, middleware, productHandler)
	routers.ProductTypeRouter(api, middleware, productTypeHandler)

	return &ServerHTTP{Engine: engine, config: config}
}

func (s *ServerHTTP) Start() error {
	return s.Engine.Run(s.config.ServerAction)
}
