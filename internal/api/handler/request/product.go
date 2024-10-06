package IRequest

import "time"

type CreateProduct struct {
	Name            string    `json:"name" binding:"required"`
	Slug            string    `json:"slug" binding:"required"`
	Description     string    `json:"description" binding:"required"`
	Discount        int32     `json:"discount"`
	DiscountEndDate time.Time `json:"discountEndDate" example:"[{\"value\": \"null or time\"}]"`
	DiscountStart   time.Time `json:"discountStart" example:"[{\"value\": \"null or time\"}]"`
	Image           string    `json:"image" binding:"required"`
	Location        int32     `json:"location" binding:"required"`
	Type            int32     `json:"type" binding:"required"`
	Status          int32     `json:"status" binding:"required"`
	Price           int32     `json:"price" binding:"required"`
	CountInStock    int32     `json:"countInStock" binding:"required"`
}

type UpdateProduct struct {
	Name            string    `json:"name" binding:"required"`
	Slug            string    `json:"slug" binding:"required"`
	Description     string    `json:"description" binding:"required"`
	Discount        int32     `json:"discount"`
	DiscountEndDate time.Time `json:"discountEndDate" example:"[{\"value\": \"null or time\"}]"`
	DiscountStart   time.Time `json:"discountStart" example:"[{\"value\": \"null or time\"}]"`
	Image           string    `json:"image" binding:"required"`
	Location        int32     `json:"location" binding:"required"`
	Type            int32     `json:"type" binding:"required"`
	Status          int32     `json:"status" binding:"required"`
	Price           int32     `json:"price" binding:"required"`
	CountInStock    int32     `json:"countInStock" binding:"required"`
}

type GetAllProduct struct {
	Limit  int32  `form:"limit" binding:"required,min=1"`
	Page   int32  `form:"page" binding:"required,min=1"`
	Search string `form:"search"`
	Order  string `form:"order"`
}

type GetAllProductLiked struct {
	Limit  int32  `form:"limit" binding:"required,min=1"`
	Page   int32  `form:"page" binding:"required,min=1"`
	Search string `form:"search"`
}

type GetAllProductAdmin struct {
	Limit       int32  `form:"limit" binding:"required,min=1"`
	Page        int32  `form:"page" binding:"required,min=1"`
	Search      string `form:"search"`
	ProductType int32  `form:"productType"`
	Status      int32  `form:"status"`
	Order  		string `form:"order"`
}

type GetAllProductPublic struct {
	Limit       int32  `form:"limit" binding:"required,min=1"`
	Page        int32  `form:"page" binding:"required,min=1"`
	Search      string `form:"search"`
	ProductType int32  `form:"productType"`
	Status      int32  `form:"status"`
}

type GetAllProductViewed struct {
	Limit  int32  `form:"limit" binding:"required,min=1"`
	Page   int32  `form:"page" binding:"required,min=1"`
	Search string `form:"search"`
}

type GetAllProductRelated struct {
	Limit int32  `form:"limit" binding:"required"`
	Page  int32  `form:"page" binding:"required"`
	Slug  string `form:"slug" binding:"required"`
}

type GetProduct struct {
	ID int64 `uri:"id" binding:"required"`
}

type GetProductPublicById struct {
	ID int64 `uri:"productId" binding:"required"`
}

type UpdateProductUrl struct {
	ID int64 `uri:"productId" binding:"required"`
}

type GetProductBySlug struct {
	Slug string `uri:"productSlug" binding:"required,min=1"`
}

// like
type LikeProduct struct {
	ID int64 `json:"productId" binding:"required"`
}

type UnLikeProduct struct {
	ID int64 `json:"productId" binding:"required"`
}

type GetParamsIsViewed struct {
	IsViewed bool `form:"isViewed" binding:"required"`
}

type GetParamsUpdateProduct struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

type GetBodyUpdateProduct struct {
	Name string `json:"name" `
	Slug string `json:"slug"`
}

// Delete
type DeleteProduct struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

type DeleteManyProduct struct {
	ArrayId []int64 `json:"arrayId" binding:"required"`
}
