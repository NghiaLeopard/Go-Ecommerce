package routers

import (
	IHandler "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/handler/interfaces"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/middleware"
	"github.com/gin-gonic/gin"
)

func CityRouter(api *gin.RouterGroup, middleware middleware.Middleware, cityHandler IHandler.City) {
	apiCity := api.Group("/city")
	{
		apiCity.POST("", middleware.AuthMiddleware("CITY.CREATE", false, false), cityHandler.CreateCity)
		apiCity.GET("/:id", middleware.AuthMiddleware("1", true, false), cityHandler.GetCity)
		apiCity.PATCH("/:id", middleware.AuthMiddleware("1", true, false), cityHandler.UpdateCity)
		apiCity.DELETE("/:id", middleware.AuthMiddleware("1", true, false), cityHandler.DeleteCity)
		apiCity.DELETE("", middleware.AuthMiddleware("1", true, false), cityHandler.DeleteManyCity)
	}
}
