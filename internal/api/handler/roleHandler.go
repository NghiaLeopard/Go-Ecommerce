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

type RoleHandler struct {
	RoleUseCase IUseCase.Role
}

// CreateRole 		godoc
// @security 		BearerAuth
// @Summary 		Create role
// @Description 	Create role
// @Param 			tags body IRequest.CreateRole true "Create Role"
// @Produce 		application/json
// @Tags 			Role
// @Success 		200 {object} IResponse.Role{}
// @Router 			/api/role [post]
func NewRoleHandler(roleUseCase IUseCase.Role) IHandler.Role {
	return &RoleHandler{RoleUseCase: roleUseCase}
}

func (r *RoleHandler) CreateRole(ctx *gin.Context) {
	var req IRequest.CreateRole
	if err := ctx.ShouldBindJSON(&req); err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		response.ErrorResponse(ctx, "Body is invalid or not exist", 400)
		return
	}

	Role, err, codeStatus := r.RoleUseCase.CreateRole(ctx, req.Name)

	if err != nil {
		response.ErrorResponse(ctx, err.Error(), codeStatus)
		return
	}

	global.Logger.Info("create role", zap.String("Status", "Success"))
	response.SuccessResponse(ctx, "Create role success", codeStatus, Role)
}

// GetRole 			godoc
// @security 		BearerAuth
// @Summary 		Get role by id
// @Description 	Get role by id
// @Param roleId  	path int true "User ID"
// @Produce 		application/json
// @Tags 			Role
// @Success 		200 {object} IResponse.Role{}
// @Router 			/api/role/{roleId} [get]
func (c *RoleHandler) GetRole(ctx *gin.Context) {
	var req IRequest.GetRole
	if err := ctx.ShouldBindUri(&req); err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		response.ErrorResponse(ctx, "Body is invalid or not exist", 400)
		return
	}

	role, err, codeStatus := c.RoleUseCase.GetRoleUseCase(ctx, req.ID)

	if err != nil {
		response.ErrorResponse(ctx, err.Error(), codeStatus)
		return
	}

	global.Logger.Info("get Role", zap.String("Status", "Success"))
	response.SuccessResponse(ctx, "Get Role success", codeStatus, role)

}

// GetAllRole 		godoc
// @security 		BearerAuth
// @Summary 		Get all role
// @Description 	Get all role
// @Param 			request query IRequest.GetAllRole true "get all Role"
// @Produce 		application/json
// @Tags 			Role
// @Success 		200 {array} []IResponse.Role{}
// @Router 			/api/role [get]
func (c *RoleHandler) GetAllRole(ctx *gin.Context) {
	var req IRequest.GetAllRole
	if err := ctx.ShouldBindQuery(&req); err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		response.ErrorResponse(ctx, "Body is invalid or not exist", 400)
		return
	}

	role, err, codeStatus := c.RoleUseCase.GetAllRoleUseCase(ctx, req)

	if err != nil {
		response.ErrorResponse(ctx, err.Error(), codeStatus)
		return
	}

	global.Logger.Info("get Role", zap.String("Status", "Success"))
	response.SuccessResponse(ctx, "Get Role success", codeStatus, role)

}

// UpdateRole 		godoc
// @security 		BearerAuth
// @Summary 		Update role
// @Description 	Update role
// @Param roleId 	path int true "Update Role"
// @Param 			tags body IRequest.GetBodyUpdateRole true "Update Role"
// @Produce 		application/json
// @Tags 			Role
// @Success 		200 {object} IResponse.Role{}
// @Router 			/api/role/{roleId} [put]
func (c *RoleHandler) UpdateRole(ctx *gin.Context) {
	var params IRequest.GetParamsUpdateRole
	var body IRequest.GetBodyUpdateRole
	if err := ctx.ShouldBindUri(&params); err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		response.ErrorResponse(ctx, "Body is invalid or not exist", 400)
		return
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		response.ErrorResponse(ctx, "Body is invalid or not exist", 400)
		return
	}

	Role, err, codeStatus := c.RoleUseCase.UpdateRoleUseCase(ctx, params.ID, body.Name, body.Permission)

	if err != nil {
		response.ErrorResponse(ctx, err.Error(), codeStatus)
		return
	}

	global.Logger.Info("get Role", zap.String("Status", "Success"))
	response.SuccessResponse(ctx, "Get Role success", codeStatus, Role)
}

// DeleteRole 		godoc
// @security 		BearerAuth
// @Summary 		Delete Role
// @Description 	Delete Role
// @Param RoleId 	path int true "Delete Role"
// @Produce 		application/json
// @Tags 			Role
// @Success 		200 {string} string [delete Role success]
// @Router 			/api/role/{roleId} [delete]
func (c *RoleHandler) DeleteRole(ctx *gin.Context) {
	var req IRequest.DeleteRole
	if err := ctx.ShouldBindUri(&req); err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		response.ErrorResponse(ctx, "Body is invalid or not exist", 400)
		return
	}

	err, codeStatus := c.RoleUseCase.DeleteRoleUseCase(ctx, req.ID)

	if err != nil {
		response.ErrorResponse(ctx, err.Error(), codeStatus)
		return
	}

	global.Logger.Info("get Role", zap.String("Status", "Success"))
	response.SuccessResponse(ctx, "Delete Role success", codeStatus, "")
}

// DeleteManyRole 	godoc
// @security 		BearerAuth
// @Summary 		Delete many role
// @Description 	Delete many role
// @Param 			tags body IRequest.DeleteManyRole true "DeleteMany Role"
// @Produce 		application/json
// @Tags 			Role
// @Success 		200 {string} string "Delete many Role success"
// @Router 			/api/role/delete-many [delete]
func (c *RoleHandler) DeleteManyRole(ctx *gin.Context) {
	var req IRequest.DeleteManyRole
	if err := ctx.ShouldBindJSON(&req); err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		response.ErrorResponse(ctx, "Body is invalid or not exist", 400)
		return
	}

	err, codeStatus := c.RoleUseCase.DeleteManyRoleUseCase(ctx, req.ArrayId)

	if err != nil {
		response.ErrorResponse(ctx, err.Error(), codeStatus)
		return
	}

	global.Logger.Info("get Role", zap.String("Status", "Success"))
	response.SuccessResponse(ctx, "Delete Role success", codeStatus, "")
}
