package IUseCase

import (
	db "github.com/NghiaLeopard/Go-Ecommerce-Backend/db/sqlc"
	IRequest "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/handler/request"
	IResponse "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/handler/response"
	"github.com/gin-gonic/gin"
)

type User interface {
	CreateUserUseCase(ctx *gin.Context, req IRequest.CreateUser) (error, int)
	GetUserUseCase(ctx *gin.Context, id int) (db.GetUserAdminByIdRow, error, int)
	GetAllUserUseCase(ctx *gin.Context, page int32, limit int32, search string, order string) (IResponse.GetAllUser, error, int)
	UpdateUserUseCase(ctx *gin.Context, id int, body IRequest.GetBodyUpdateUser) (error, int)
	DeleteUserUseCase(ctx *gin.Context, id int) (error, int)
	DeleteManyUserUseCase(ctx *gin.Context, id []int) (error, int)
}
