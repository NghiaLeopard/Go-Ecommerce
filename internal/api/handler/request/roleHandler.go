package IRequest

type CreateRole struct {
	Name string `json:"name" binding:"required,min=6"`
}

type GetAllRole struct {
	Limit  int32  `form:"limit" binding:"required,min=1"`
	Page   int32  `form:"page" binding:"required,min=1"`
	Search string `form:"search"`
	Order  string `form:"order"`
}

type GetRole struct {
	ID int `uri:"id" binding:"required,min=1"`
}

type DeleteRole struct {
	ID int `uri:"id" binding:"required,min=1"`
}

type DeleteManyRole struct {
	ArrayId []int `json:"arrayId" binding:"required"`
}

type GetParamsUpdateRole struct {
	ID int `uri:"id" binding:"required,min=1"`
}

type GetBodyUpdateRole struct {
	Name       string   `json:"name" `
	Permission []string `json:"permission"`
}
