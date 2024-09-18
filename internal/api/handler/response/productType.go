package IResponse

import "time"

type ProductType struct {
	Id       int64     `json:"_id"`
	Name     string    `json:"name"`
	Slug     string    `json:"slug"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
}
