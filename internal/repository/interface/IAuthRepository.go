package IRepository

import (
	db "github.com/NghiaLeopard/Go-Ecommerce-Backend/db/sqlc"
	IRequest "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/handler/request"
	"github.com/gin-gonic/gin"
)

type Auth interface {
	CreateUser(ctx *gin.Context, arg db.CreateUserParams) (db.User, error)
	InitDefaultAdmin(ctx *gin.Context, arg db.InitDefaultAdminParams) (db.User, error)
	GetUserById(ctx *gin.Context, id int64) (db.GetUserByIdRow, error)
	GetUserByEmail(ctx *gin.Context, email string) (db.GetUserByEmailRow, error)
	UpdateAuthMe(ctx *gin.Context, req IRequest.UpdateAuthMe, id int64) (db.User, error)
	UpdatePasswordUser(ctx *gin.Context, arg db.UpdatePasswordUserParams) error
	DeleteUser(ctx *gin.Context, id int64) error
	FindUserById(ctx *gin.Context, id int64) error
}
