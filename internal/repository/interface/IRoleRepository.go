package IRepository

import (
	db "github.com/NghiaLeopard/Go-Ecommerce-Backend/db/sqlc"
	IRequest "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/handler/request"
	"github.com/gin-gonic/gin"
)

type Role interface {
	CreateRole(ctx *gin.Context, name string) (db.Role, error)
	GetRoleById(ctx *gin.Context, id int64) (db.GetRoleByIdRow, error)
	GetRoleByName(ctx *gin.Context, name string) (db.Role, error)
	GetAllRole(ctx *gin.Context, req IRequest.GetAllRole) ([]db.ListRoleRow, error)
	UpdateRole(ctx *gin.Context, id int64, name string, permission []string) (db.Role, error)
	DeleteRole(ctx *gin.Context, id int64) error
	DeleteManyRole(ctx *gin.Context, arrayId []int64) error
}
