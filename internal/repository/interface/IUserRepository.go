package IRepository

import (
	db "github.com/NghiaLeopard/Go-Ecommerce-Backend/db/sqlc"
	IRequest "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/handler/request"
	"github.com/gin-gonic/gin"
)

type User interface {
	CreateUserAdmin(ctx *gin.Context, req IRequest.CreateUser) (db.User, error)
	UpdateUser(ctx *gin.Context, id int64, body IRequest.GetBodyUpdateUser) (db.User, error)
	GetAllUser(ctx *gin.Context, page int32, limit int32, search string, order string) ([]db.ListUserAdminRow, error)
	GetUserById(ctx *gin.Context, id int64) (db.GetUserAdminByIdRow, error)
	GetUserByEmail(ctx *gin.Context, email string) (db.GetUserByEmailRow, error)
	DeleteUserById(ctx *gin.Context, id int64) error
	DeleteManyUserByIds(ctx *gin.Context, arrayId []int64) error
}
