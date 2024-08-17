package middleware

import (
	"strings"

	"github.com/NghiaLeopard/Go-Ecommerce-Backend/pkg/response"
	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "authorization"
	authorizationType   = "Bearer"
	authorizationKey    = "authorization_payload"
)

func (c *middleware) authMiddleWare(permission string, isAuthMe bool, isPublic bool) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorization := ctx.GetHeader(authorizationHeader)

		if len(authorization) == 0 {
			response.ErrorResponse(ctx, "please provide authorization", 400)
		}

		fields := strings.Fields(authorization)

		if len(fields) < 2 {
			response.ErrorResponse(ctx, "invalid format header", 400)
		}

		if fields[0] != authorizationType {
			response.ErrorResponse(ctx, "invalid type header", 400)
		}

		payload, err := c.Token.VerifyTokenPaseto(fields[1])

		if err != nil {
			response.ErrorResponse(ctx, "Internal server", 500)
		}

		ctx.Set(authorizationKey, payload)
	}
}
