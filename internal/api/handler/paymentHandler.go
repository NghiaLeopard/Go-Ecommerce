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

type PaymentHandler struct {
	PaymentUseCase IUseCase.Payment
}

func NewPaymentHandler(PaymentUseCase IUseCase.Payment) IHandler.Payment {
	return &PaymentHandler{PaymentUseCase: PaymentUseCase}
}

// CreatePayment 		godoc
// @security 		BearerAuth
// @Summary 		Create Payment
// @Description 	Create Payment
// @Param 			tags body IRequest.CreatePayment true "Create Payment"
// @Produce 		application/json
// @Tags 			Payment
// @Success 		200 {object} IResponse.Payment{}
// @Router 			/api/Payment-type [post]
func (c *PaymentHandler) CreatePayment(ctx *gin.Context) {
	var req IRequest.CreatePayment
	var _ *IResponse.Payment
	if err := ctx.ShouldBindJSON(&req); err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		response.ErrorResponse(ctx, "Body is invalid or not exist", 400)
		return
	}

	Payment, err, codeStatus := c.PaymentUseCase.CreatePaymentUseCase(ctx, req)

	if err != nil {
		response.ErrorResponse(ctx, err.Error(), codeStatus)
		return
	}

	global.Logger.Info("create Payment", zap.String("Status", "Success"))
	response.SuccessResponse(ctx, "Create Payment success", codeStatus, Payment)
}

// GetPayment 			godoc
// @security 		BearerAuth
// @Summary 		Get Payment by id
// @Description 	Get Payment by id
// @Param PaymentId  	path int true "User ID"
// @Produce 		application/json
// @Tags 			Payment
// @Success 		200 {object} IResponse.Payment{}
// @Router 			/api/Payment-type/{PaymentId} [get]
func (c *PaymentHandler) GetPayment(ctx *gin.Context) {
	var req IRequest.GetPayment
	if err := ctx.ShouldBindUri(&req); err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		response.ErrorResponse(ctx, "Body is invalid or not exist", 400)
		return
	}

	Payment, err, codeStatus := c.PaymentUseCase.GetPaymentUseCase(ctx, req.ID)

	if err != nil {
		response.ErrorResponse(ctx, err.Error(), codeStatus)
		return
	}

	global.Logger.Info("get Payment", zap.String("Status", "Success"))
	response.SuccessResponse(ctx, "Get Payment success", codeStatus, Payment)
}

// GetAllPayment 		godoc
// @security 		BearerAuth
// @Summary 		Get all Payment
// @Description 	Get all Payment
// @Param 			request query IRequest.GetAllPayment true "get all Payment"
// @Produce 		application/json
// @Tags 			Payment
// @Success 		200 {array} []IResponse.Payment{}
// @Router 			/api/Payment-type [get]
func (c *PaymentHandler) GetAllPayment(ctx *gin.Context) {
	var req IRequest.GetAllPayment
	if err := ctx.ShouldBindQuery(&req); err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		response.ErrorResponse(ctx, "Body is invalid or not exist", 400)
		return
	}

	Payment, err, codeStatus := c.PaymentUseCase.GetAllPaymentUseCase(ctx, req.Page, req.Limit, req.Search, req.Order)

	if err != nil {
		response.ErrorResponse(ctx, err.Error(), codeStatus)
		return
	}

	global.Logger.Info("get Payment", zap.String("Status", "Success"))
	response.SuccessResponse(ctx, "Get Payment success", codeStatus, Payment)
}

// UpdatePayment 		godoc
// @security 			BearerAuth
// @Summary 			Update Payment
// @Description 		Update Payment
// @Param PaymentId 	path int true "Update Payment"
// @Param 				tags body IRequest.GetBodyUpdatePayment true "Update Payment"
// @Produce 			application/json
// @Tags 				Payment
// @Success 			200 {object} IResponse.Payment{}
// @Router 				/api/Payment-type/{PaymentId} [put]
func (c *PaymentHandler) UpdatePayment(ctx *gin.Context) {
	var params IRequest.GetParamsUpdatePayment
	var body IRequest.GetBodyUpdatePayment
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

	Payment, err, codeStatus := c.PaymentUseCase.UpdatePaymentUseCase(ctx, params.ID, body)

	if err != nil {
		response.ErrorResponse(ctx, err.Error(), codeStatus)
		return
	}

	global.Logger.Info("get Payment", zap.String("Status", "Success"))
	response.SuccessResponse(ctx, "Get Payment success", codeStatus, Payment)
}

// DeletePayment 		godoc
// @security 			BearerAuth
// @Summary 			Delete Payment
// @Description 		Delete Payment
// @Param PaymentId 	path int true "Delete Payment"
// @Produce 			application/json
// @Tags 				Payment
// @Success 			200 {string} string [delete Payment success]
// @Router 				/api/Payment-type/{PaymentId} [delete]
func (c *PaymentHandler) DeletePayment(ctx *gin.Context) {
	var req IRequest.DeletePayment
	if err := ctx.ShouldBindUri(&req); err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		response.ErrorResponse(ctx, "Body is invalid or not exist", 400)
		return
	}

	err, codeStatus := c.PaymentUseCase.DeletePaymentUseCase(ctx, req.ID)

	if err != nil {
		response.ErrorResponse(ctx, err.Error(), codeStatus)
		return
	}

	global.Logger.Info("get Payment", zap.String("Status", "Success"))
	response.SuccessResponse(ctx, "Delete Payment success", codeStatus, map[string]int{"_id": 1})
}

// DeleteManyPayment 		godoc
// @security 		BearerAuth
// @Summary 		Delete many Payment
// @Description 	Delete many Payment
// @Param 			tags body IRequest.DeleteManyPayment true "DeleteMany Payment"
// @Produce 		application/json
// @Tags 			Payment
// @Success 		200 {string} string "Delete many Payment success"
// @Router 			/api/Payment-type/delete-many [delete]
func (c *PaymentHandler) DeleteManyPayment(ctx *gin.Context) {
	var req IRequest.DeleteManyPayment
	if err := ctx.ShouldBindJSON(&req); err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		response.ErrorResponse(ctx, "Body is invalid or not exist", 400)
		return
	}

	err, codeStatus := c.PaymentUseCase.DeleteManyPaymentUseCase(ctx, req.ArrayId)

	if err != nil {
		response.ErrorResponse(ctx, err.Error(), codeStatus)
		return
	}

	global.Logger.Info("get Payment", zap.String("Status", "Success"))
	response.SuccessResponse(ctx, "Delete Payment success", codeStatus, map[string]int{"_id": 1})
}
