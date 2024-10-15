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

type DeliveryHandler struct {
	DeliveryUseCase IUseCase.Delivery
}

func NewDeliveryHandler(DeliveryUseCase IUseCase.Delivery) IHandler.Delivery {
	return &DeliveryHandler{DeliveryUseCase: DeliveryUseCase}
}

// CreateDelivery 		godoc
// @security 		BearerAuth
// @Summary 		Create Delivery
// @Description 	Create Delivery
// @Param 			tags body IRequest.CreateDelivery true "Create Delivery"
// @Produce 		application/json
// @Tags 			Delivery
// @Success 		200 {object} IResponse.Delivery{}
// @Router 			/api/delivery-type [post]
func (c *DeliveryHandler) CreateDelivery(ctx *gin.Context) {
	var req IRequest.CreateDelivery
	var _ *IResponse.Delivery
	if err := ctx.ShouldBindJSON(&req); err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		response.ErrorResponse(ctx, "Body is invalid or not exist", 400)
		return
	}

	Delivery, err, codeStatus := c.DeliveryUseCase.CreateDeliveryUseCase(ctx, req)

	if err != nil {
		response.ErrorResponse(ctx, err.Error(), codeStatus)
		return
	}

	global.Logger.Info("create Delivery", zap.String("Status", "Success"))
	response.SuccessResponse(ctx, "Create Delivery success", codeStatus, Delivery)
}

// GetDelivery 			godoc
// @security 		BearerAuth
// @Summary 		Get Delivery by id
// @Description 	Get Delivery by id
// @Param DeliveryId  	path int true "User ID"
// @Produce 		application/json
// @Tags 			Delivery
// @Success 		200 {object} IResponse.Delivery{}
// @Router 			/api/delivery-type/{DeliveryId} [get]
func (c *DeliveryHandler) GetDelivery(ctx *gin.Context) {
	var req IRequest.GetDelivery
	if err := ctx.ShouldBindUri(&req); err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		response.ErrorResponse(ctx, "Body is invalid or not exist", 400)
		return
	}

	Delivery, err, codeStatus := c.DeliveryUseCase.GetDeliveryUseCase(ctx, req.ID)

	if err != nil {
		response.ErrorResponse(ctx, err.Error(), codeStatus)
		return
	}

	global.Logger.Info("get Delivery", zap.String("Status", "Success"))
	response.SuccessResponse(ctx, "Get Delivery success", codeStatus, Delivery)
}

// GetAllDelivery 		godoc
// @security 		BearerAuth
// @Summary 		Get all Delivery
// @Description 	Get all Delivery
// @Param 			request query IRequest.GetAllDelivery true "get all Delivery"
// @Produce 		application/json
// @Tags 			Delivery
// @Success 		200 {array} []IResponse.Delivery{}
// @Router 			/api/delivery-type [get]
func (c *DeliveryHandler) GetAllDelivery(ctx *gin.Context) {
	var req IRequest.GetAllDelivery
	if err := ctx.ShouldBindQuery(&req); err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		response.ErrorResponse(ctx, "Body is invalid or not exist", 400)
		return
	}

	Delivery, err, codeStatus := c.DeliveryUseCase.GetAllDeliveryUseCase(ctx, req.Page, req.Limit, req.Search, req.Order)

	if err != nil {
		response.ErrorResponse(ctx, err.Error(), codeStatus)
		return
	}

	global.Logger.Info("get Delivery", zap.String("Status", "Success"))
	response.SuccessResponse(ctx, "Get Delivery success", codeStatus, Delivery)
}

// UpdateDelivery 		godoc
// @security 			BearerAuth
// @Summary 			Update Delivery
// @Description 		Update Delivery
// @Param DeliveryId 	path int true "Update Delivery"
// @Param 				tags body IRequest.GetBodyUpdateDelivery true "Update Delivery"
// @Produce 			application/json
// @Tags 				Delivery
// @Success 			200 {object} IResponse.Delivery{}
// @Router 				/api/delivery-type/{DeliveryId} [put]
func (c *DeliveryHandler) UpdateDelivery(ctx *gin.Context) {
	var params IRequest.GetParamsUpdateDelivery
	var body IRequest.GetBodyUpdateDelivery
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

	Delivery, err, codeStatus := c.DeliveryUseCase.UpdateDeliveryUseCase(ctx, params.ID, body)

	if err != nil {
		response.ErrorResponse(ctx, err.Error(), codeStatus)
		return
	}

	global.Logger.Info("get Delivery", zap.String("Status", "Success"))
	response.SuccessResponse(ctx, "Get Delivery success", codeStatus, Delivery)
}

// DeleteDelivery 		godoc
// @security 			BearerAuth
// @Summary 			Delete Delivery
// @Description 		Delete Delivery
// @Param DeliveryId 	path int true "Delete Delivery"
// @Produce 			application/json
// @Tags 				Delivery
// @Success 			200 {string} string [delete Delivery success]
// @Router 				/api/delivery-type/{DeliveryId} [delete]
func (c *DeliveryHandler) DeleteDelivery(ctx *gin.Context) {
	var req IRequest.DeleteDelivery
	if err := ctx.ShouldBindUri(&req); err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		response.ErrorResponse(ctx, "Body is invalid or not exist", 400)
		return
	}

	err, codeStatus := c.DeliveryUseCase.DeleteDeliveryUseCase(ctx, req.ID)

	if err != nil {
		response.ErrorResponse(ctx, err.Error(), codeStatus)
		return
	}

	global.Logger.Info("get Delivery", zap.String("Status", "Success"))
	response.SuccessResponse(ctx, "Delete Delivery success", codeStatus, map[string]int{"_id": 1})
}

// DeleteManyDelivery 		godoc
// @security 		BearerAuth
// @Summary 		Delete many Delivery
// @Description 	Delete many Delivery
// @Param 			tags body IRequest.DeleteManyDelivery true "DeleteMany Delivery"
// @Produce 		application/json
// @Tags 			Delivery
// @Success 		200 {string} string "Delete many Delivery success"
// @Router 			/api/delivery-type/delete-many [delete]
func (c *DeliveryHandler) DeleteManyDelivery(ctx *gin.Context) {
	var req IRequest.DeleteManyDelivery
	if err := ctx.ShouldBindJSON(&req); err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		response.ErrorResponse(ctx, "Body is invalid or not exist", 400)
		return
	}

	err, codeStatus := c.DeliveryUseCase.DeleteManyDeliveryUseCase(ctx, req.ArrayId)

	if err != nil {
		response.ErrorResponse(ctx, err.Error(), codeStatus)
		return
	}

	global.Logger.Info("get Delivery", zap.String("Status", "Success"))
	response.SuccessResponse(ctx, "Delete Delivery success", codeStatus, map[string]int{"_id": 1})
}
