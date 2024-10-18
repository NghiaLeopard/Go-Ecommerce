package handler

import (
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/global"
	IHandler "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/handler/interfaces"
	IRequest "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/handler/request"
	IResponse "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/handler/response"
	IUseCase "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/usecase/interfaces"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/pkg/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserHandler struct {
	UserUseCase IUseCase.User
}

func NewUserHandler(UserUseCase IUseCase.User) IHandler.User {
	return &UserHandler{UserUseCase: UserUseCase}
}

// CreateUser 		godoc
// @security 		BearerAuth
// @Summary 		Create User
// @Description 	Create User
// @Param 			tags body IRequest.CreateUser true "Create User"
// @Produce 		application/json
// @Tags 			User
// @Success 		200 {object} IResponse.User{}
// @Router 			/api/User [post]
func (c *UserHandler) CreateUser(ctx *gin.Context) {
	var req IRequest.CreateUser
	var _ *IResponse.User
	if err := ctx.ShouldBindJSON(&req); err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		response.ErrorResponse(ctx, "Body is invalid or not exist", 400)
		return
	}

	err, codeStatus := c.UserUseCase.CreateUserUseCase(ctx, req)

	if err != nil {
		response.ErrorResponse(ctx, err.Error(), codeStatus)
		return
	}

	global.Logger.Info("create User", zap.String("Status", "Success"))
	response.SuccessResponse(ctx, "Create User success", codeStatus, map[string]int{"_id": 1})
}

// GetUser 			godoc
// @security 		BearerAuth
// @Summary 		Get User by id
// @Description 	Get User by id
// @Param UserId  	path int true "User ID"
// @Produce 		application/json
// @Tags 			User
// @Success 		200 {object} IResponse.User{}
// @Router 			/api/User/{UserId} [get]
func (c *UserHandler) GetUser(ctx *gin.Context) {
	var req IRequest.GetUser
	if err := ctx.ShouldBindUri(&req); err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		response.ErrorResponse(ctx, "Body is invalid or not exist", 400)
		return
	}

	User, err, codeStatus := c.UserUseCase.GetUserUseCase(ctx, req.ID)

	if err != nil {
		response.ErrorResponse(ctx, err.Error(), codeStatus)
		return
	}

	global.Logger.Info("get User", zap.String("Status", "Success"))
	response.SuccessResponse(ctx, "Get User success", codeStatus, User)
}

// GetAllUser 		godoc
// @security 		BearerAuth
// @Summary 		Get all User
// @Description 	Get all User
// @Param 			request query IRequest.GetAllUser true "get all User"
// @Produce 		application/json
// @Tags 			User
// @Success 		200 {array} []IResponse.User{}
// @Router 			/api/User [get]
func (c *UserHandler) GetAllUser(ctx *gin.Context) {
	var req IRequest.GetAllUser
	if err := ctx.ShouldBindQuery(&req); err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		response.ErrorResponse(ctx, "Body is invalid or not exist", 400)
		return
	}

	User, err, codeStatus := c.UserUseCase.GetAllUserUseCase(ctx, req.Page, req.Limit, req.Search, req.Order)

	if err != nil {
		response.ErrorResponse(ctx, err.Error(), codeStatus)
		return
	}

	global.Logger.Info("get User", zap.String("Status", "Success"))
	response.SuccessResponse(ctx, "Get User success", codeStatus, User)
}

// UpdateUser 		godoc
// @security 		BearerAuth
// @Summary 		Update User
// @Description 	Update User
// @Param UserId 	path int true "Update User"
// @Param 			tags body IRequest.GetBodyUpdateUser true "Update User"
// @Produce 		application/json
// @Tags 			User
// @Success 		200 {object} IResponse.User{}
// @Router 			/api/User/{UserId} [put]
func (c *UserHandler) UpdateUser(ctx *gin.Context) {
	var params IRequest.GetParamsUpdateUser
	var body IRequest.GetBodyUpdateUser
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

	err, codeStatus := c.UserUseCase.UpdateUserUseCase(ctx, params.ID, body.Name)

	if err != nil {
		response.ErrorResponse(ctx, err.Error(), codeStatus)
		return
	}

	global.Logger.Info("get User", zap.String("Status", "Success"))
	response.SuccessResponse(ctx, "Get User success", codeStatus, map[string]int{"_id": 1})
}

// DeleteUser 		godoc
// @security 		BearerAuth
// @Summary 		Delete User
// @Description 	Delete User
// @Param UserId 	path int true "Delete User"
// @Produce 		application/json
// @Tags 			User
// @Success 		200 {string} string [delete User success]
// @Router 			/api/User/{UserId} [delete]
func (c *UserHandler) DeleteUser(ctx *gin.Context) {
	var req IRequest.DeleteUser
	if err := ctx.ShouldBindUri(&req); err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		response.ErrorResponse(ctx, "Body is invalid or not exist", 400)
		return
	}

	err, codeStatus := c.UserUseCase.DeleteUserUseCase(ctx, req.ID)

	if err != nil {
		response.ErrorResponse(ctx, err.Error(), codeStatus)
		return
	}

	global.Logger.Info("get User", zap.String("Status", "Success"))
	response.SuccessResponse(ctx, "Delete User success", codeStatus, map[string]int{"_id": 1})
}

// DeleteManyUser 		godoc
// @security 		BearerAuth
// @Summary 		Delete many User
// @Description 	Delete many User
// @Param 			tags body IRequest.DeleteManyUser true "DeleteMany User"
// @Produce 		application/json
// @Tags 			User
// @Success 		200 {string} string "Delete many User success"
// @Router 			/api/User/delete-many [delete]
func (c *UserHandler) DeleteManyUser(ctx *gin.Context) {
	var req IRequest.DeleteManyUser
	if err := ctx.ShouldBindJSON(&req); err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		response.ErrorResponse(ctx, "Body is invalid or not exist", 400)
		return
	}

	err, codeStatus := c.UserUseCase.DeleteManyUserUseCase(ctx, req.ArrayId)

	if err != nil {
		response.ErrorResponse(ctx, err.Error(), codeStatus)
		return
	}

	global.Logger.Info("get User", zap.String("Status", "Success"))
	response.SuccessResponse(ctx, "Delete User success", codeStatus, map[string]int{"_id": 1})
}
