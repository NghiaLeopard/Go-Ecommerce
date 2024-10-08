package IUseCase

import (
	IResponse "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/handler/response"
	"github.com/gin-gonic/gin"
)

type City interface {
	CreateCityUseCase(ctx *gin.Context, name string) (IResponse.City, error, int)
	GetCityUseCase(ctx *gin.Context, id int) (IResponse.City, error, int)
	GetAllCityUseCase(ctx *gin.Context, page int32, limit int32, search string, order string) (IResponse.GetAllCity, error, int)
	UpdateCityUseCase(ctx *gin.Context, id int, name string) (IResponse.City, error, int)
	DeleteCityUseCase(ctx *gin.Context, id int) (error, int)
	DeleteManyCityUseCase(ctx *gin.Context, id []int) (error, int)
}
