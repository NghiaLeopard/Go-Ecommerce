package routers

import (
	IHandler "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/handler/interfaces"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/middleware"
	"github.com/gin-gonic/gin"
)

func CityRouter(api *gin.RouterGroup, middleware middleware.Middleware, cityHandler IHandler.ICityHandler) {
	city := api.Use(middleware.AuthMiddleware("CITY.CREATE", false, false))
	{
		city.POST("/city", cityHandler.CreateCity)
	}
}
