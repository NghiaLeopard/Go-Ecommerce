package routers

import (
	IHandler "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/handler/interfaces"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/middleware"
	"github.com/gin-gonic/gin"
)

func PaymentRouter(api *gin.RouterGroup, middleware middleware.Middleware, PaymentHandler IHandler.Payment) {
	apiPayment := api.Group("/payment-type")
	{
		apiPayment.POST("", middleware.AuthMiddleware("SETTING.PAYMENT_TYPE.CREATE", false, false), PaymentHandler.CreatePayment)
		apiPayment.GET("", PaymentHandler.GetAllPayment)
		apiPayment.GET("/:id", PaymentHandler.GetPayment)
		apiPayment.PUT("/:id", middleware.AuthMiddleware("SETTING.PAYMENT_TYPE.UPDATE", false, false), PaymentHandler.UpdatePayment)
		apiPayment.DELETE("/:id", middleware.AuthMiddleware("SETTING.PAYMENT_TYPE.DELETE", false, false), PaymentHandler.DeletePayment)
		apiPayment.DELETE("/delete-many", middleware.AuthMiddleware("SETTING.PAYMENT_TYPE.DELETE", false, false), PaymentHandler.DeleteManyPayment)
	}
}
