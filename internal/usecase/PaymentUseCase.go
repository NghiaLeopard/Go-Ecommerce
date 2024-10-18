package usecase

import (
	"fmt"

	"github.com/NghiaLeopard/Go-Ecommerce-Backend/global"
	IRequest "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/handler/request"
	IResponse "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/handler/response"
	IRepository "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/repository/interface"
	IUseCase "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/usecase/interfaces"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/pkg/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type PaymentUseCase struct {
	PaymentRepo IRepository.Payment
}

func NewPaymentUseCase(PaymentRepo IRepository.Payment) IUseCase.Payment {
	return &PaymentUseCase{PaymentRepo: PaymentRepo}
}

func (c *PaymentUseCase) CreatePaymentUseCase(ctx *gin.Context, req IRequest.CreatePayment) (IResponse.Payment, error, int) {
	_, err := c.PaymentRepo.GetPaymentByName(ctx, req.Name)

	if err == nil {
		global.Logger.Error("Payment is  exist", zap.String("Status", "Error"))
		return IResponse.Payment{}, fmt.Errorf("Payment is  exist"), 409
	}

	Payment, err := c.PaymentRepo.CreatePayment(ctx, req)

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return IResponse.Payment{}, err, 401
	}

	return IResponse.Payment{
		Id:   Payment.ID,
		Name: Payment.Name,
		Type: Payment.Type,
	}, nil, 201
}

func (c *PaymentUseCase) GetPaymentUseCase(ctx *gin.Context, id int) (IResponse.Payment, error, int) {

	Payment, err := c.PaymentRepo.GetPaymentById(ctx, int64(id))

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return IResponse.Payment{}, fmt.Errorf("get Payment is not exist"), 401
	}

	return IResponse.Payment{
		Id:       Payment.ID,
		Name:     Payment.Name,
		Type:     Payment.Type,
		CreateAt: Payment.CreateAt,
	}, nil, 200
}

func (c *PaymentUseCase) GetAllPaymentUseCase(ctx *gin.Context, page int32, limit int32, search string, order string) (IResponse.GetAllPayment, error, int) {
	Payment, err := c.PaymentRepo.GetAllPayment(ctx, page, limit, search, order)

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return IResponse.GetAllPayment{}, fmt.Errorf("get Payment is not exist"), 401
	}

	if len(Payment) == 0 {
		return IResponse.GetAllPayment{
			PaymentTypes: Payment,
			TotalCount:   0,
			TotalPage:    0,
		}, nil, 200
	}

	totalPage := utils.PageCount(int64(limit), Payment[0].TotalCount)

	return IResponse.GetAllPayment{
		PaymentTypes: Payment,
		TotalCount:   Payment[0].TotalCount,
		TotalPage:    totalPage,
	}, nil, 200
}

func (c *PaymentUseCase) UpdatePaymentUseCase(ctx *gin.Context, id int, body IRequest.GetBodyUpdatePayment) (IResponse.Payment, error, int) {
	idInt64 := int64(id)

	_, err := global.DB.GetPaymentById(ctx, idInt64)

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return IResponse.Payment{}, fmt.Errorf("Payment is not exist"), 401
	}

	Payment, err := c.PaymentRepo.UpdatePayment(ctx, idInt64, body)

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return IResponse.Payment{}, fmt.Errorf("update Payment is fail"), 401
	}

	res := IResponse.Payment{
		Id:   Payment.ID,
		Name: Payment.Name,
		Type: Payment.Type,
	}

	return res, nil, 200
}

func (c *PaymentUseCase) DeletePaymentUseCase(ctx *gin.Context, id int) (error, int) {
	err := global.DB.DeletePaymentById(ctx, int64(id))

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return fmt.Errorf("get Payment is not exist"), 401
	}

	return nil, 200
}

func (c *PaymentUseCase) DeleteManyPaymentUseCase(ctx *gin.Context, arrayId []int) (error, int) {
	if len(arrayId) == 0 {
		global.Logger.Error("ArrayID is empty", zap.String("Status", "Error"))
		return fmt.Errorf("ArrayID is empty"), 401
	}

	arrayInt64 := make([]int64, len(arrayId))

	for i, v := range arrayId {
		arrayInt64[i] = int64(v)
	}

	err := global.DB.DeleteManyPaymentByIds(ctx, arrayInt64)

	if err != nil {
		global.Logger.Error(err.Error(), zap.String("Status", "Error"))
		return fmt.Errorf("get Payment is not exist"), 401
	}

	return nil, 200
}
