package IRequest

type CreateCity struct {
	Name string `json:"name" binding:"required,min=6"`
}

type GetCity struct {
	ID int `uri:"id" binding:"required,min=1"`
}

type DeleteCity struct {
	ID int `uri:"id" binding:"required,min=1"`
}

type DeleteManyCity struct {
	ArrayId []int `json:"arrayId" binding:"required,min=1"`
}

type GetParamsUpdateCity struct {
	ID int `uri:"id" binding:"required,min=1"`
}

type GetBodyUpdateCity struct {
	Name string `json:"name" binding:"required,min=1"`
}
