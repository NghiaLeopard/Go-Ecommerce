package repository

import (
	db "github.com/NghiaLeopard/Go-Ecommerce-Backend/db/sqlc"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/global"
	IRequest "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/handler/request"
	IRepository "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/repository/interface"
	"github.com/gin-gonic/gin"
)

type DeliveryRepository struct{}

func NewDeliveryRepository() IRepository.Delivery {
	return &DeliveryRepository{}
}

func (c *DeliveryRepository) CreateDelivery(ctx *gin.Context, req IRequest.CreateDelivery) (Delivery db.DeliveryType, err error) {
	arg := db.CreateDeliveryParams{
		Name:  req.Name,
		Price: req.Price,
	}
	Delivery, err = global.DB.CreateDelivery(ctx, arg)

	return
}

func (c *DeliveryRepository) UpdateDelivery(ctx *gin.Context, id int64, body IRequest.GetBodyUpdateDelivery) (Delivery db.DeliveryType, err error) {
	arg := db.UpdateDeliveryParams{
		ID:    id,
		Name:  body.Name,
		Price: body.Price,
	}

	Delivery, err = global.DB.UpdateDelivery(ctx, arg)

	return
}

func (c *DeliveryRepository) GetDeliveryById(ctx *gin.Context, id int64) (Delivery db.DeliveryType, err error) {
	Delivery, err = global.DB.GetDeliveryById(ctx, id)

	return
}

func (c *DeliveryRepository) GetAllDelivery(ctx *gin.Context, page int32, limit int32, search string, order string) (Delivery []db.ListDeliveryRow, err error) {

	offset := limit * (page - 1)

	arg := db.ListDeliveryParams{
		LimitOpt:  limit,
		OffsetOpt: offset,
		OrderBy:   order,
		Search:    search,
	}
	Delivery, err = global.DB.ListDelivery(ctx, arg)

	return
}

func (c *DeliveryRepository) GetDeliveryByName(ctx *gin.Context, name string) (Delivery db.DeliveryType, err error) {
	Delivery, err = global.DB.GetDeliveryByName(ctx, name)

	return
}

func (c *DeliveryRepository) DeleteDeliveryById(ctx *gin.Context, id int64) (err error) {
	err = global.DB.DeleteDeliveryById(ctx, id)

	return
}

func (c *DeliveryRepository) DeleteManyDeliveryByIds(ctx *gin.Context, arrayID []int64) (err error) {

	err = global.DB.DeleteManyDeliveryByIds(ctx, arrayID)

	return
}
