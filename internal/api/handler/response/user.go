package IResponse

import (
	db "github.com/NghiaLeopard/Go-Ecommerce-Backend/db/sqlc"
)

type GetAllUser struct {
	Users      []db.ListUserAdminRow `json:"users"`
	TotalCount int64                 `json:"totalCount"`
	TotalPage  int64                 `json:"totalPage"`
}
