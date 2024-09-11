package IHandler

import "github.com/gin-gonic/gin"

type Role interface {
	CreateCity(ctx *gin.Context)
	GetCity(ctx *gin.Context)
	UpdateCity(ctx *gin.Context)
	DeleteCity(ctx *gin.Context)
	DeleteManyCity(ctx *gin.Context)
}
