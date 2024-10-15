package IResponse

import (
	"time"

	db "github.com/NghiaLeopard/Go-Ecommerce-Backend/db/sqlc"
)

type Delivery struct {
	Id       int64     `json:"_id"`
	Name     string    `json:"name"`
	Price    int32     `json:"price"`
	CreateAt time.Time `json:"create_at"`
}

type GetAllDelivery struct {
	Cities     []db.ListDeliveryRow `json:"cities"`
	TotalCount int64                `json:"totalCount"`
	TotalPage  int64                `json:"totalPage"`
}
