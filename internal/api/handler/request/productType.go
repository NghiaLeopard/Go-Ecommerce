package IRequest

type CreateProductType struct {
	Name string `json:"name" binding:"required,min=4"`
	Slug string `json:"slug" binding:"required,min=1"`
}

type GetProductType struct {
	ID int `uri:"id" binding:"required,min=1"`
}

type DeleteProductType struct {
	ID int `uri:"id" binding:"required,min=1"`
}

type DeleteManyProductType struct {
	ArrayId []int `json:"arrayId" binding:"required"`
}

type GetParamsUpdateProductType struct {
	ID int `uri:"id" binding:"required,min=1"`
}

type GetBodyUpdateProductType struct {
	Name string `json:"name" `
	Slug string `json:"slug"`
}
