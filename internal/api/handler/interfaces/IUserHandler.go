package IHandler

import "github.com/gin-gonic/gin"

type User interface {
	CreateUser(ctx *gin.Context)
	GetAllUser(ctx *gin.Context)
	GetUser(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
	DeleteManyUser(ctx *gin.Context)
}
