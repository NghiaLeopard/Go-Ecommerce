package IRequest

type CreatePayment struct {
	Name string `json:"name" binding:"required"`
	Type string `json:"type" binding:"required,min=1"`
}

type GetAllPayment struct {
	Limit  int32  `form:"limit"`
	Page   int32  `form:"page"`
	Search string `form:"search"`
	Order  string `form:"order"`
}

type GetPayment struct {
	ID int `uri:"id" binding:"required,min=1"`
}

type DeletePayment struct {
	ID int `uri:"id" binding:"required,min=1"`
}

type DeleteManyPayment struct {
	ArrayId []int `json:"paymentTypeIds" binding:"required,min=1"`
}

type GetParamsUpdatePayment struct {
	ID int `uri:"id" binding:"required,min=1"`
}

type GetBodyUpdatePayment struct {
	Name string `json:"name" binding:"required,min=1"`
	Type string `json:"type" binding:"required,min=1"`
}
