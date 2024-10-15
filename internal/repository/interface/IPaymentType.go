package IRepository

import (
	db "github.com/NghiaLeopard/Go-Ecommerce-Backend/db/sqlc"
	IRequest "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/handler/request"
	"github.com/gin-gonic/gin"
)

type Payment interface {
	CreatePayment(ctx *gin.Context, req IRequest.CreatePayment) (db.PaymentType, error)
	UpdatePayment(ctx *gin.Context, id int64, body IRequest.GetBodyUpdatePayment) (db.PaymentType, error)
	GetAllPayment(ctx *gin.Context, page int32, limit int32, search string, order string) ([]db.ListPaymentRow, error)
	GetPaymentById(ctx *gin.Context, id int64) (db.PaymentType, error)
	GetPaymentByName(ctx *gin.Context, name string) (db.PaymentType, error)
	DeletePaymentById(ctx *gin.Context, id int64) error
	DeleteManyPaymentByIds(ctx *gin.Context, arrayId []int64) error
}
