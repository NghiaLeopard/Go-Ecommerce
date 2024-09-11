package IHandler

import "github.com/gin-gonic/gin"

type Role interface {
	CreateRole(ctx *gin.Context)
	GetRole(ctx *gin.Context)
	UpdateRole(ctx *gin.Context)
	DeleteRole(ctx *gin.Context)
	DeleteManyRole(ctx *gin.Context)
}
