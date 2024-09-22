package middleware

import (
	IRepository "github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/repository/interface"
	"github.com/gin-gonic/gin"
)

type Middleware interface {
	AuthMiddleware(permission string, isAuthMe bool, isPublic bool) gin.HandlerFunc
}

type middleware struct {
	Redis IRepository.RedisToken
}

func NewMiddleware(redis IRepository.RedisToken) Middleware {
	return &middleware{Redis: redis}
}
