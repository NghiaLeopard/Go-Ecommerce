package usecase

import (
	"database/sql"
	"fmt"

	"github.com/NghiaLeopard/Go-Ecommerce-Backend/global"
	IResponse "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/handler/response"
	IRepository "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/repository/interface"
	IUseCase "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/usecase/interfaces"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type RoleUseCase struct {
	RoleRepo IRepository.Role
}

func NewRoleUseCase(roleRepo IRepository.Role) IUseCase.Role {
	return &RoleUseCase{
		RoleRepo: roleRepo,
	}
}

func (c *RoleUseCase) CreateRole(ctx *gin.Context, name string) (IResponse.Role, error, int) {
	_, err := c.RoleRepo.GetRoleByName(ctx, name)

	if err != sql.ErrNoRows {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return IResponse.Role{}, fmt.Errorf("role is  exist"), 409
	}

	role, err := global.DB.CreateRole(ctx, name)

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return IResponse.Role{}, err, 401
	}

	global.Logger.Info("Create ", zap.String("Status", "Error"))
	return IResponse.Role{
		Id:         role.ID,
		Name:       role.Name,
		Permission: role.Permission,
	}, nil, 201
}
