package IHandler

import "github.com/gin-gonic/gin"

type IAuthHandler interface {
	LoginUser(ctx *gin.Context)
	SignUpUser(ctx *gin.Context)
}