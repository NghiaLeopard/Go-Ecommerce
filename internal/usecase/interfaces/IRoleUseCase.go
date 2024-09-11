package IUseCase

import (
	IResponse "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/handler/response"
	"github.com/gin-gonic/gin"
)

type Role interface {
	CreateRole(ctx *gin.Context, name string) (IResponse.Role, error, int)
}
