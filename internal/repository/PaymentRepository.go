package repository

import (
	db "github.com/NghiaLeopard/Go-Ecommerce-Backend/db/sqlc"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/global"
	IRequest "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/handler/request"
	IRepository "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/repository/interface"
	"github.com/gin-gonic/gin"
)

type PaymentRepository struct{}

func NewPaymentRepository() IRepository.Payment {
	return &PaymentRepository{}
}

func (c *PaymentRepository) CreatePayment(ctx *gin.Context, req IRequest.CreatePayment) (Payment db.PaymentType, err error) {
	arg := db.CreatePaymentParams{
		Name: req.Name,
		Type: req.Type,
	}
	Payment, err = global.DB.CreatePayment(ctx, arg)

	return
}

func (c *PaymentRepository) UpdatePayment(ctx *gin.Context, id int64, body IRequest.GetBodyUpdatePayment) (Payment db.PaymentType, err error) {
	arg := db.UpdatePaymentParams{
		ID:   id,
		Name: body.Name,
		Type: body.Type,
	}

	Payment, err = global.DB.UpdatePayment(ctx, arg)

	return
}

func (c *PaymentRepository) GetPaymentById(ctx *gin.Context, id int64) (Payment db.PaymentType, err error) {
	Payment, err = global.DB.GetPaymentById(ctx, id)

	return
}

func (c *PaymentRepository) GetAllPayment(ctx *gin.Context, page int32, limit int32, search string, order string) (Payment []db.ListPaymentRow, err error) {

	offset := limit * (page - 1)

	arg := db.ListPaymentParams{
		LimitOpt:  limit,
		OffsetOpt: offset,
		OrderBy:   order,
		Search:    search,
	}
	Payment, err = global.DB.ListPayment(ctx, arg)

	return
}

func (c *PaymentRepository) GetPaymentByName(ctx *gin.Context, name string) (Payment db.PaymentType, err error) {
	Payment, err = global.DB.GetPaymentByName(ctx, name)

	return
}

func (c *PaymentRepository) DeletePaymentById(ctx *gin.Context, id int64) (err error) {
	err = global.DB.DeletePaymentById(ctx, id)

	return
}

func (c *PaymentRepository) DeleteManyPaymentByIds(ctx *gin.Context, arrayID []int64) (err error) {

	err = global.DB.DeleteManyPaymentByIds(ctx, arrayID)

	return
}
