package IRequest

type CreateCity struct {
	Name string `json:"name" binding:"required,min=6"`
}

type GetCity struct {
	ID int `uri:"id" binding:"required,min=1"`
}

type GetParamsUpdateCity struct {
	ID int `uri:"id" binding:"required,min=1"`
}

type GetBodyUpdateCity struct {
	Name string `uri:"name" binding:"required,min=1"`
}
