package IResponse

import (
	"time"

	db "github.com/NghiaLeopard/Go-Ecommerce-Backend/db/sqlc"
)

type ProductType struct {
	Id       int64     `json:"_id"`
	Name     string    `json:"name"`
	Slug     string    `json:"slug"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
}

type GetAllProductType struct {
	ProductTypes []db.ListProductTypeRow `json:"productTypes"`
	TotalCount   int64                   `json:"totalCount"`
	TotalPage    int64                   `json:"totalPage"`
}
