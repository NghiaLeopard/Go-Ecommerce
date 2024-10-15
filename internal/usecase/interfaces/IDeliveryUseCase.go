package IUseCase

import (
	IRequest "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/handler/request"
	IResponse "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/handler/response"
	"github.com/gin-gonic/gin"
)

type Delivery interface {
	CreateDeliveryUseCase(ctx *gin.Context, req IRequest.CreateDelivery) (IResponse.Delivery, error, int)
	GetDeliveryUseCase(ctx *gin.Context, id int) (IResponse.Delivery, error, int)
	GetAllDeliveryUseCase(ctx *gin.Context, page int32, limit int32, search string, order string) (IResponse.GetAllDelivery, error, int)
	UpdateDeliveryUseCase(ctx *gin.Context, id int, name string) (IResponse.Delivery, error, int)
	DeleteDeliveryUseCase(ctx *gin.Context, id int) (error, int)
	DeleteManyDeliveryUseCase(ctx *gin.Context, id []int) (error, int)
}
