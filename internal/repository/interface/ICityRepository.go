package IRepository

import (
	db "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/db/sqlc"
	"github.com/gin-gonic/gin"
)

type ICity interface {
	CreateCity(ctx *gin.Context, name string) (db.City, error)
}
