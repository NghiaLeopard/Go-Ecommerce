package IRequest

type CreateUser struct {
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required,min=6"`
	Address     string `json:"address" binding:"required"`
	Avatar      string `json:"avatar" binding:"required"`
	City        int64  `json:"city" binding:"required" format:"int64"`
	Role        int64  `json:"role" binding:"required" format:"int64"`
	FirstName   string `json:"firstName" binding:"required"`
	LastName    string `json:"lastName" binding:"required"`
	MiddleName  string `json:"middleName" binding:"required"`
	PhoneNumber string `json:"phoneNumber" binding:"required" format:"string"`
}

type GetAllUser struct {
	Limit  int32  `form:"limit"`
	Page   int32  `form:"page"`
	Search string `form:"search"`
	Order  string `form:"order"`
}

type GetUser struct {
	ID int `uri:"id" binding:"required,min=1"`
}

type DeleteUser struct {
	ID int `uri:"id" binding:"required,min=1"`
}

type DeleteManyUser struct {
	ArrayId []int `json:"arrayId" binding:"required,min=1"`
}

type GetParamsUpdateUser struct {
	ID int `uri:"id" binding:"required,min=1"`
}

type GetBodyUpdateUser struct {
	Name string `json:"name" binding:"required,min=1"`
}
