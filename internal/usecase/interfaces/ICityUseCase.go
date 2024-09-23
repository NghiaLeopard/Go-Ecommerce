package IUseCase

import (
	IResponse "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/api/handler/response"
	db "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/db/sqlc"
	"github.com/gin-gonic/gin"
)

type City interface {
	CreateCityUseCase(ctx *gin.Context, name string) (IResponse.City, error, int)
	GetCityUseCase(ctx *gin.Context, id int) (IResponse.City, error, int)
	GetAllCityUseCase(ctx *gin.Context, page int, limit int, search string, order string) ([]db.City, error, int)
	UpdateCityUseCase(ctx *gin.Context, id int, name string) (IResponse.City, error, int)
	DeleteCityUseCase(ctx *gin.Context, id int) (error, int)
	DeleteManyCityUseCase(ctx *gin.Context, id []int) (error, int)
}
