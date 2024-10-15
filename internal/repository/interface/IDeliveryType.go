package IRepository

import (
	db "github.com/NghiaLeopard/Go-Ecommerce-Backend/db/sqlc"
	IRequest "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/handler/request"
	"github.com/gin-gonic/gin"
)

type Delivery interface {
	CreateDelivery(ctx *gin.Context, req IRequest.CreateDelivery) (db.DeliveryType, error)
	UpdateDelivery(ctx *gin.Context, id int64, body IRequest.GetBodyUpdateDelivery) (db.DeliveryType, error)
	GetAllDelivery(ctx *gin.Context, page int32, limit int32, search string, order string) ([]db.ListDeliveryRow, error)
	GetDeliveryById(ctx *gin.Context, id int64) (db.DeliveryType, error)
	GetDeliveryByName(ctx *gin.Context, name string) (db.DeliveryType, error)
	DeleteDeliveryById(ctx *gin.Context, id int64) error
	DeleteManyDeliveryByIds(ctx *gin.Context, arrayId []int64) error
}
