package IUseCase

import (
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/handler/response"
	"github.com/gin-gonic/gin"
)

type ICityUseCase interface {
	CreateCityUseCase(ctx *gin.Context, name string) (response.ICityResponse, error, int)
	GetCityUseCase(ctx *gin.Context, id int) (response.ICityResponse, error, int)
}
