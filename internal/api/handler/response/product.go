package IResponse

import (
	"encoding/json"
	"time"

	db "github.com/NghiaLeopard/Go-Ecommerce-Backend/db/sqlc"
)

type Product struct {
	Id                int64     `json:"_id"`
	Name              string    `json:"name"`
	Slug              string    `json:"slug"`
	Image             string    `json:"image"`
	Price             int32     `json:"price"`
	CountInStock      int32     `json:"countInStock"`
	Description       string    `json:"description"`
	Discount          int32     `json:"discount"`
	DiscountStartDate time.Time `json:"discountStartDate"`
	DiscountEndDate   time.Time `json:"discountEndDate"`
	Type              int32     `json:"type"`
	Location          int32     `json:"location"`
	LikedBy           []int32   `json:"likedBy"`
	TotalLikes        int64     `json:"totalLikes"`
	Status            int32     `json:"status"`
	Views             int32     `json:"views"`
	UniqueViews       []int32   `json:"uniqueViews"`
	CreateAt          time.Time `json:"createdAt"`
}

type GetProduct struct {
	Id                int64           `json:"_id"`
	Name              string          `json:"name"`
	Slug              string          `json:"slug"`
	Image             string          `json:"image"`
	Price             int32           `json:"price"`
	CountInStock      int32           `json:"countInStock"`
	Description       string          `json:"description"`
	Discount          int32           `json:"discount"`
	DiscountStartDate time.Time       `json:"discountStartDate"`
	DiscountEndDate   time.Time       `json:"discountEndDate"`
	Type              int32           `json:"type"`
	Location          int32           `json:"location"`
	LikedBy           json.RawMessage `json:"likedBy"`
	TotalLikes        int64           `json:"totalLikes"`
	Status            int32           `json:"status"`
	Views             int32           `json:"views"`
	UniqueViews       json.RawMessage `json:"uniqueViews" `
	CreateAt          time.Time       `json:"createdAt"`
}

type UpdateProduct struct {
	Id                int64     `json:"_id"`
	Name              string    `json:"name"`
	Slug              string    `json:"slug"`
	Image             string    `json:"image"`
	Price             int32     `json:"price"`
	CountInStock      int32     `json:"countInStock"`
	Description       string    `json:"description"`
	Discount          int32     `json:"discount"`
	DiscountStartDate time.Time `json:"discountStartDate"`
	DiscountEndDate   time.Time `json:"discountEndDate"`
	Type              int32     `json:"type"`
	Location          int32     `json:"location"`
	Status            int32     `json:"status"`
	Views             int32     `json:"views"`
	CreateAt          time.Time `json:"createdAt"`
}

type GetAllMeLiked struct {
	Products   []db.GetAllProductLikeRow `json:"products"`
	TotalCount int64                     `json:"totalCount"`
	TotalPage  int64                     `json:"totalPage"`
}

type GetAllMeViewed struct {
	Products   []db.GetAllProductViewRow `json:"products"`
	TotalCount int64                     `json:"totalCount"`
	TotalPage  int64                     `json:"totalPage"`
}

type GetAllProductRelated struct {
	Products   []db.GetAllProductRelatedRow `json:"products"`
	TotalCount int64                        `json:"totalCount"`
	TotalPage  int64                        `json:"totalPage"`
}
