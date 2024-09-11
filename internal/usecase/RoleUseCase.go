package usecase

import (
	IRepository "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/repository/interface"
	IUseCase "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/usecase/interfaces"
	"github.com/gin-gonic/gin"
)

type RoleUseCase struct {
	RoleCity IRepository.Role
}

func NewRoleUseCase(roleCity IRepository.Role) IUseCase.Role {
	return &RoleUseCase{
		RoleCity: roleCity,
	}
}

func (u *RoleUseCase) CreateCity(ctx *gin.Context,name string,permission []string)() {

}
