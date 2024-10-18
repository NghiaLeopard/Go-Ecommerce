package routers

import (
	IHandler "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/handler/interfaces"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/middleware"
	"github.com/gin-gonic/gin"
)

func UserRouter(api *gin.RouterGroup, middleware middleware.Middleware, UserHandler IHandler.User) {
	apiUser := api.Group("/users")
	{
		apiUser.POST("", middleware.AuthMiddleware("SYSTEM.USER.CREATE", false, false), UserHandler.CreateUser)
		apiUser.GET("", UserHandler.GetAllUser)
		apiUser.GET("/:id", UserHandler.GetUser)
		apiUser.PUT("/:id", middleware.AuthMiddleware("SYSTEM.USER.UPDATE", false, false), UserHandler.UpdateUser)
		apiUser.DELETE("/:id", middleware.AuthMiddleware("SYSTEM.USER.DELETE", false, false), UserHandler.DeleteUser)
		apiUser.DELETE("/delete-many", middleware.AuthMiddleware("SYSTEM.USER.DELETE", false, false), UserHandler.DeleteManyUser)
	}
}
