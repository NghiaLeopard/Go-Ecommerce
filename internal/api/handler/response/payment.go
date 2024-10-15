package IResponse

import (
	"time"

	db "github.com/NghiaLeopard/Go-Ecommerce-Backend/db/sqlc"
)

type Payment struct {
	Id       int64     `json:"_id"`
	Name     string    `json:"name"`
	Type     string    `json:"type"`
	CreateAt time.Time `json:"create_at"`
}

type GetAllPayment struct {
	PaymentTypes []db.ListPaymentRow `json:"cities"`
	TotalCount   int64               `json:"totalCount"`
	TotalPage    int64               `json:"totalPage"`
}
