package routers

import (
	IHandler "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/handler/interfaces"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/middleware"
	"github.com/gin-gonic/gin"
)

func ProductRouter(api *gin.RouterGroup, middleware middleware.Middleware, ProductHandler IHandler.Product) {
	apiProduct := api.Group("/products")
	{
		apiProduct.POST("", middleware.AuthMiddleware("MANAGE_PRODUCT.PRODUCT.CREATE", false, false), ProductHandler.CreateProduct)
		apiProduct.GET("/:id", middleware.AuthMiddleware("MANAGE_PRODUCT.PRODUCT.VIEW", true, false), ProductHandler.GetProduct)
		// apiProduct.GET("", middleware.AuthMiddleware("MANAGE_PRODUCT.PRODUCT.VIEW", true, false), ProductHandler.GetAllProduct)
		// apiProduct.PUT("/:id", middleware.AuthMiddleware("MANAGE_PRODUCT.PRODUCT.UPDATE", true, false), ProductHandler.UpdateProduct)
		apiProduct.DELETE("/:id", middleware.AuthMiddleware("MANAGE_PRODUCT.PRODUCT.DELETE", true, false), ProductHandler.DeleteProduct)
		apiProduct.DELETE("", middleware.AuthMiddleware("MANAGE_PRODUCT.PRODUCT.DELETE", true, false), ProductHandler.DeleteManyProduct)

		productPublic := apiProduct.Group("/public")
		{
			productPublic.GET("/:id", middleware.AuthMiddleware("1", false, true), ProductHandler.GetProduct)
			productPublic.GET("/slug/:slug", middleware.AuthMiddleware("1", false, true), ProductHandler.GetProductBySlug)
		}

	}

}
