package IHandler

import "github.com/gin-gonic/gin"

type ICityHandler interface {
	CreateCity(ctx *gin.Context)
	GetCity(ctx *gin.Context)
}