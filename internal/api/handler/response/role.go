package IResponse

import db "github.com/NghiaLeopard/Go-Ecommerce-Backend/db/sqlc"

type Role struct {
	Id         int64    `json:"_id"`
	Name       string   `json:"name"`
	Permission []string `json:"permissions"`
}

type GetAllRole struct {
	Roles      []db.ListRoleRow `json:"roles"`
	TotalCount int64            `json:"totalCount"`
	TotalPage  int64            `json:"totalPage"`
}
