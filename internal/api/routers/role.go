package routers

import (
	IHandler "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/handler/interfaces"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/middleware"
	"github.com/gin-gonic/gin"
)

func RoleRouter(api *gin.RouterGroup, middleware middleware.Middleware, RoleHandler IHandler.Role) {
	apiRole := api.Group("/role")
	{
		apiRole.POST("", middleware.AuthMiddleware("ROLE.CREATE", false, false), RoleHandler.CreateRole)
		apiRole.GET("", middleware.AuthMiddleware("1", true, false), RoleHandler.GetAlRole)
		apiRole.GET("/:id", middleware.AuthMiddleware("1", true, false), RoleHandler.GetRole)
		apiRole.PATCH("/:id", middleware.AuthMiddleware("1", true, false), RoleHandler.UpdateRole)
		apiRole.DELETE("/:id", middleware.AuthMiddleware("1", true, false), RoleHandler.DeleteRole)
		apiRole.DELETE("", middleware.AuthMiddleware("1", true, false), RoleHandler.DeleteManyRole)
	}
}
