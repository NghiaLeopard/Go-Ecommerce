package routers

import (
	IHandler "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/handler/interfaces"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/middleware"
	"github.com/gin-gonic/gin"
)

func DeliveryRouter(api *gin.RouterGroup, middleware middleware.Middleware, DeliveryHandler IHandler.Delivery) {
	apiDelivery := api.Group("/delivery-type")
	{
		apiDelivery.POST("", middleware.AuthMiddleware("SETTING.DELIVERY_TYPE.CREATE", false, false), DeliveryHandler.CreateDelivery)
		apiDelivery.GET("", DeliveryHandler.GetAllDelivery)
		apiDelivery.GET("/:id", DeliveryHandler.GetDelivery)
		apiDelivery.PUT("/:id", middleware.AuthMiddleware("SETTING.DELIVERY_TYPE.UPDATE", false, false), DeliveryHandler.UpdateDelivery)
		apiDelivery.DELETE("/:id", middleware.AuthMiddleware("SETTING.DELIVERY_TYPE.DELETE", false, false), DeliveryHandler.DeleteDelivery)
		apiDelivery.DELETE("/delete-many", middleware.AuthMiddleware("SETTING.DELIVERY_TYPE.DELETE", false, false), DeliveryHandler.DeleteManyDelivery)
	}
}
