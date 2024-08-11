package interfaces

import "github.com/gin-gonic/gin"

type IAuthHandler interface {
	LoginUser(ctx *gin.Context)
}