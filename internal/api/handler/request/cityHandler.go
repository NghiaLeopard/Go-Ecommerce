package IRequest

type CreateCity struct {
	Name string `json:"name" binding:"required"`
}
