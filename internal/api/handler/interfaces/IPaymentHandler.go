package IHandler

import "github.com/gin-gonic/gin"

type Payment interface {
	CreatePayment(ctx *gin.Context)
	GetAllPayment(ctx *gin.Context)
	GetPayment(ctx *gin.Context)
	UpdatePayment(ctx *gin.Context)
	DeletePayment(ctx *gin.Context)
	DeleteManyPayment(ctx *gin.Context)
}
