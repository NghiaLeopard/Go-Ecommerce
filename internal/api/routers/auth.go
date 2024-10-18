package routers

import (
	IHandler "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/handler/interfaces"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/middleware"
	"github.com/gin-gonic/gin"
)

func AuthRouter(api *gin.RouterGroup, middleware middleware.Middleware, authHandler IHandler.Auth) {
	auth := api.Group("/auth")
	{
		auth.POST("/register", authHandler.SignUpUser)
		auth.POST("/login", authHandler.LoginUser)
		auth.POST("/logout", authHandler.LogoutUser)
		auth.POST("/refresh-token", authHandler.RefreshToken)

		authMePublic := auth.Use(middleware.AuthMiddleware("1", true, true))
		{
			authMePublic.POST("/forgot-password", authHandler.ForgotPasswordUser)
			authMePublic.POST("/reset-password", authHandler.ResetPasswordUser)
		}

		authMe := auth.Use(middleware.AuthMiddleware("1", true, false))
		{
			authMe.PATCH("/change-password", authHandler.ChangePasswordUser)
			authMe.GET("/me", authHandler.GetAuthMe)
			authMe.PUT("/me", authHandler.UpdateAuthMe)
		}
	}
}
