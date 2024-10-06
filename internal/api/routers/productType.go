package routers

import (
	IHandler "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/handler/interfaces"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/middleware"
	"github.com/gin-gonic/gin"
)

func ProductTypeRouter(api *gin.RouterGroup, middleware middleware.Middleware, ProductTypeHandler IHandler.ProductType) {
	apiProductType := api.Group("/product-types")
	{
		apiProductType.POST("", middleware.AuthMiddleware("MANAGE_PRODUCT.PRODUCT_TYPE.CREATE", false, false), ProductTypeHandler.CreateProductType)
		apiProductType.GET("", ProductTypeHandler.GetAllProductType)
		apiProductType.GET("/:id", ProductTypeHandler.GetProductType)
		apiProductType.PUT("/:id", middleware.AuthMiddleware("MANAGE_PRODUCT.PRODUCT_TYPE.UPDATE", true, false), ProductTypeHandler.UpdateProductType)
		apiProductType.DELETE("/:id", middleware.AuthMiddleware("MANAGE_PRODUCT.PRODUCT_TYPE.DELETE", true, false), ProductTypeHandler.DeleteProductType)
		apiProductType.DELETE("", middleware.AuthMiddleware("MANAGE_PRODUCT.PRODUCT_TYPE.DELETE", true, false), ProductTypeHandler.DeleteManyProductType)
	}
}
