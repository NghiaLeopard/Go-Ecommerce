package IRequest

type CreateDelivery struct {
	Name  string `json:"name" binding:"required"`
	Price int32  `json:"price" binding:"required,min=1"`
}

type GetAllDelivery struct {
	Limit  int32  `form:"limit"`
	Page   int32  `form:"page"`
	Search string `form:"search"`
	Order  string `form:"order"`
}

type GetDelivery struct {
	ID int `uri:"id" binding:"required,min=1"`
}

type DeleteDelivery struct {
	ID int `uri:"id" binding:"required,min=1"`
}

type DeleteManyDelivery struct {
	ArrayId []int `json:"arrayId" binding:"required,min=1"`
}

type GetParamsUpdateDelivery struct {
	ID int `uri:"id" binding:"required,min=1"`
}

type GetBodyUpdateDelivery struct {
	Name string `json:"name" binding:"required,min=1"`
}
