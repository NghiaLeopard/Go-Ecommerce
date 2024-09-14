package IUseCase

import (
	IResponse "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/handler/response"
	"github.com/gin-gonic/gin"
)

type Role interface {
	CreateRole(ctx *gin.Context, name string) (IResponse.Role, error, int)
	GetRoleUseCase(ctx *gin.Context, id int) (IResponse.Role, error, int)
	UpdateRoleUseCase(ctx *gin.Context, id int, name string,permission []string) (IResponse.Role, error, int)
	DeleteRoleUseCase(ctx *gin.Context, id int) (error, int)
	DeleteManyRoleUseCase(ctx *gin.Context, id []int) (error, int)
}
