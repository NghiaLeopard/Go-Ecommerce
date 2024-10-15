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
		apiCity.GET("", cityHandler.GetAllCity)
		apiCity.GET("/:id", cityHandler.GetCity)
		apiCity.PUT("/:id", middleware.AuthMiddleware("CITY.UPDATE", false, false), cityHandler.UpdateCity)
		apiCity.DELETE("/:id", middleware.AuthMiddleware("CITY.DELETE", false, false), cityHandler.DeleteCity)
		apiCity.DELETE("/delete-many", middleware.AuthMiddleware("CITY.DELETE", false, false), cityHandler.DeleteManyCity)
	}
}
