package IRequest

type CreateCity struct {
	Name string `json:"name" binding:"required"`
}

type GetCity struct {
	ID int `uri:"id" binding:"required"`
}
