package IHandler

import "github.com/gin-gonic/gin"

type Delivery interface {
	CreateDelivery(ctx *gin.Context)
	GetAllDelivery(ctx *gin.Context)
	GetDelivery(ctx *gin.Context)
	UpdateDelivery(ctx *gin.Context)
	DeleteDelivery(ctx *gin.Context)
	DeleteManyDelivery(ctx *gin.Context)
}
