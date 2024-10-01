package IResponse

import (
	"time"

	db "github.com/NghiaLeopard/Go-Ecommerce-Backend/db/sqlc"
)

type City struct {
	Id       int64     `json:"_id"`
	Name     string    `json:"name"`
	CreateAt time.Time `json:"create_at"`
}

type GetAllCity struct {
	Cities     []db.ListCityRow `json:"cities"`
	TotalCount int64            `json:"totalCount"`
	TotalPage  int64            `json:"totalPage"`
}
