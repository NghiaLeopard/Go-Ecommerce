package handler

import (
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/global"
	IHandler "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/handler/interfaces"
	IRequest "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/handler/request"
	IUseCase "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/usecase/interfaces"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/pkg/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func NewRoleHandler(roleUseCase IUseCase.Role) IHandler.Role {
	return &Role{RoleUseCase: roleUseCase}
}

type Role struct {
	RoleUseCase IUseCase.Role
}

func (r *Role) CreateRole(ctx *gin.Context) {
	var req IRequest.CreateRole
	if err := ctx.ShouldBindJSON(&req); err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		response.ErrorResponse(ctx, "Body is invalid or not exist", 400)
		return
	}

	city, err, codeStatus := r.RoleUseCase.CreateRole(ctx, req.Name)

	if err != nil {
		response.ErrorResponse(ctx, err.Error(), codeStatus)
		return
	}

	global.Logger.Error("create role", zap.String("Status", "Error"))
	response.SuccessResponse(ctx, "Create role success", codeStatus, city)
}

func (r *Role) DeleteManyRole(ctx *gin.Context) {

}

func (r *Role) DeleteRole(ctx *gin.Context) {

}

func (r *Role) GetRole(ctx *gin.Context) {

}

func (r *Role) UpdateRole(ctx *gin.Context) {

}
