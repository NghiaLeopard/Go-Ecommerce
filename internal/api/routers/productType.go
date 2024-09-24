package routers

import (
	IHandler "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/handler/interfaces"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/middleware"
	"github.com/gin-gonic/gin"
)

func ProductTypeRouter(api *gin.RouterGroup, middleware middleware.Middleware, ProductTypeHandler IHandler.ProductType) {
	apiProductType := api.Group("/product-types")
	{
		apiProductType.POST("", middleware.AuthMiddleware("MANAGE_PRODUCT.PRODUCT.CREATE", false, false), ProductTypeHandler.CreateProductType)
		apiProductType.GET("", middleware.AuthMiddleware("1", true, false), ProductTypeHandler.GetAllProductType)
		apiProductType.GET("/:id", middleware.AuthMiddleware("1", true, false), ProductTypeHandler.GetProductType)
		apiProductType.PUT("/:id", middleware.AuthMiddleware("1", true, false), ProductTypeHandler.UpdateProductType)
		apiProductType.DELETE("/:id", middleware.AuthMiddleware("1", true, false), ProductTypeHandler.DeleteProductType)
		apiProductType.DELETE("", middleware.AuthMiddleware("1", true, false), ProductTypeHandler.DeleteManyProductType)
	}
}
