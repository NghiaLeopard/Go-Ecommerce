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
		apiProduct.GET("/liked/me", middleware.AuthMiddleware("1", true, false), ProductHandler.GetAllProductMeLiked)
		apiProduct.GET("/viewed/me", middleware.AuthMiddleware("1", true, false), ProductHandler.GetAllProductMeViewed)
		apiProduct.GET("/:id", middleware.AuthMiddleware("1", true, false), ProductHandler.GetProduct)
		// apiProduct.GET("", middleware.AuthMiddleware("MANAGE_PRODUCT.PRODUCT.VIEW", true, false), ProductHandler.GetAllProduct)
		// apiProduct.PUT("/:id", middleware.AuthMiddleware("MANAGE_PRODUCT.PRODUCT.UPDATE", true, false), ProductHandler.UpdateProduct)
		apiProduct.DELETE("/:id", middleware.AuthMiddleware("1", true, false), ProductHandler.DeleteProduct)
		apiProduct.DELETE("/delete-many", middleware.AuthMiddleware("MANAGE_PRODUCT.PRODUCT.DELETE", true, false), ProductHandler.DeleteManyProduct)
		apiProduct.POST("/like", middleware.AuthMiddleware("MANAGE_PRODUCT.PRODUCT.VIEW", true, false), ProductHandler.LikeProduct)
		apiProduct.POST("/unlike", middleware.AuthMiddleware("1", true, false), ProductHandler.UnLikeProduct)
		// apiProduct.POST("/viewed/me", middleware.AuthMiddleware("MANAGE_PRODUCT.PRODUCT.VIEW", true, false), ProductHandler.GetProductMeViewed)

		productPublic := apiProduct.Group("/public")
		{
			productPublic.GET("/:productId", middleware.AuthMiddleware("1", true, true), ProductHandler.GetProductPublicById)
			productPublic.GET("/slug/:productSlug", middleware.AuthMiddleware("1", true, true), ProductHandler.GetProductBySlug)
		}

	}

}
