package IUseCase

import (
	db "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Role interface {
	CreateCity(ctx *gin.Context,name string,permission []string) (db.Role,error,int)
}
