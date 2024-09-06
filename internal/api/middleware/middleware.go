package middleware

import (
	"github.com/gin-gonic/gin"
)

type Middleware interface {
	AuthMiddleware(permission string, isAuthMe bool, isPublic bool) gin.HandlerFunc
}

type middleware struct {
}

func NewMiddleware() Middleware {
	return &middleware{}
}
