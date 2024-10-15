package IUseCase

import (
	IRequest "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/handler/request"
	IResponse "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/handler/response"
	"github.com/gin-gonic/gin"
)

type Payment interface {
	CreatePaymentUseCase(ctx *gin.Context, req IRequest.CreatePayment) (IResponse.Payment, error, int)
	GetPaymentUseCase(ctx *gin.Context, id int) (IResponse.Payment, error, int)
	GetAllPaymentUseCase(ctx *gin.Context, page int32, limit int32, search string, order string) (IResponse.GetAllPayment, error, int)
	UpdatePaymentUseCase(ctx *gin.Context, id int, body IRequest.GetBodyUpdatePayment) (IResponse.Payment, error, int)
	DeletePaymentUseCase(ctx *gin.Context, id int) (error, int)
	DeleteManyPaymentUseCase(ctx *gin.Context, id []int) (error, int)
}
