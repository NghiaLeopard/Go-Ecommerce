package routers

import (
	IHandler "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/handler/interfaces"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/middleware"
	"github.com/gin-gonic/gin"
)

func UserRouter(api *gin.RouterGroup, middleware middleware.Middleware, authHandler IHandler.IAuthHandler) {
	auth := api.Group("/auth")
	{
		auth.POST("/register", authHandler.SignUpUser)
		auth.POST("/login", authHandler.LoginUser)
		auth.POST("/logout", authHandler.LogoutUser)

		authMe := auth.Use(middleware.AuthMiddleware("1", true, false))
		{
			authMe.PATCH("/change-password", authHandler.ChangePasswordUser)
			authMe.POST("/forgot-password", authHandler.ForgotPasswordUser)
			authMe.POST("/reset-password",authHandler.ResetPasswordUser)

		}
	}

}
