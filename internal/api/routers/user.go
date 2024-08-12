package routers

import (
	IHandler "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/handler/interfaces"
	"github.com/gin-gonic/gin"
)

func UserRouter(api *gin.RouterGroup,authHandler IHandler.IAuthHandler) {
	auth := api.Group("/auth") 
	{
		auth.POST("/register",authHandler.SignUpUser)
		auth.POST("/login",authHandler.LoginUser)
	}
}