package middleware

import (
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/pkg/token"
	"github.com/gin-gonic/gin"
)

type Middleware interface {
	AuthMiddleware(permission string, isAuthMe bool, isPublic bool) gin.HandlerFunc
}

type middleware struct {
	Token token.Maker
}

func NewMiddleware(token token.Maker) Middleware {
	return &middleware{Token: token}
}
