package IRepository

import (
	db "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/db/sqlc"
	"github.com/gin-gonic/gin"
)

type City interface {
	CreateCity(ctx *gin.Context, name string) (db.City, error)
	UpdateCity(ctx *gin.Context, id int64, name string) (db.City, error)
	GetCityById(ctx *gin.Context, id int64) (db.City, error)
	GetCityByName(ctx *gin.Context, name string) (db.City, error)
	DeleteCityById(ctx *gin.Context, id int64) error
	DeleteManyCityByIds(ctx *gin.Context, arrayId []int64) error
}
