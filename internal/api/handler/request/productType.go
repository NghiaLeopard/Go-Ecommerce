package IRequest

type CreateProductType struct {
	Name string `json:"name" binding:"required,min=4"`
	Slug string `json:"slug" binding:"required,min=1"`
}

type GetAllProductType struct {
	Limit  int32  `form:"limit"`
	Page   int32  `form:"page"`
	Search string `form:"search"`
	Order  string `form:"order"`
}

type GetProductType struct {
	ID int `uri:"id" binding:"required,min=1"`
}

type DeleteProductType struct {
	ID int `uri:"id" binding:"required,min=1"`
}

type DeleteManyProductType struct {
	ArrayId []int `json:"productTypeIds" binding:"required"`
}

type GetParamsUpdateProductType struct {
	ID int `uri:"id" binding:"required,min=1"`
}

type GetBodyUpdateProductType struct {
	Name string `json:"name" `
	Slug string `json:"slug"`
}
