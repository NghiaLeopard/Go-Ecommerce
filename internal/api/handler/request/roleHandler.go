package IRequest

type CreateRole struct {
	Name string `json:"name" binding:"required,min=6"`
}

type GetRole struct {
	ID int `uri:"id" binding:"required,min=1"`
}

type DeleteRole struct {
	ID int `uri:"id" binding:"required,min=1"`
}

type DeleteManyRole struct {
	ArrayId []int `json:"arrayId" binding:"required,min=1"`
}

type GetParamsUpdateRole struct {
	ID int `uri:"id" binding:"required,min=1"`
}

type GetBodyUpdateRole struct {
	Name string `uri:"name" binding:"required,min=1"`
}
