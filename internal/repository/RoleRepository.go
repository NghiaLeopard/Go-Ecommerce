package repository

import (
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/global"
	db "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/db/sqlc"
	IRepository "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/repository/interface"
	"github.com/gin-gonic/gin"
)

type RoleRepository struct{}

func NewRoleRepository() IRepository.Role {
	return &RoleRepository{}
}

func (r *RoleRepository) CreateRole(ctx *gin.Context, name string) (role db.Role, err error) {

	role, err = global.DB.CreateRole(ctx, name)

	return
}

func (r *RoleRepository) GetRoleById(ctx *gin.Context, id int64) (db.Role, error) {
	role, err := global.DB.GetRoleById(ctx, id)

	return role, err
}

func (r *RoleRepository) GetRoleByName(ctx *gin.Context, name string) (db.Role, error) {
	role, err := global.DB.GetRoleByName(ctx, name)

	return role, err
}

func (r *RoleRepository) UpdateRole(ctx *gin.Context, id int64, name string, permission []string) (db.Role, error) {
	arg := db.UpdateRoleParams{
		ID:         id,
		Name:       name,
		Permission: permission,
	}
	role, err := global.DB.UpdateRole(ctx, arg)

	return role, err
}

func (r *RoleRepository) DeleteManyRole(ctx *gin.Context, arrayId []int64) error {
	err := global.DB.DeleteManyRolesByIds(ctx, arrayId)

	return err
}

func (r *RoleRepository) DeleteRole(ctx *gin.Context, id int64) error {
	err := global.DB.DeleteRoleById(ctx, id)

	return err
}
