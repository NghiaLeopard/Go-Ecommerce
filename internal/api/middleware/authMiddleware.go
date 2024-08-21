package middleware

import (
	"slices"
	"strings"

	"github.com/NghiaLeopard/Go-Ecommerce-Backend/pkg/config"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/pkg/response"
	"github.com/gin-gonic/gin"
)

const (
	AuthorizationHeader = "authorization"
	AuthorizationType   = "Bearer"
	AuthorizationKey    = "authorization_payload"
)

func (c *middleware) AuthMiddleware(permission string, isAuthMe bool, isPublic bool) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if isAuthMe {
			ctx.Next()
			return
		}

		if isPublic {
			ctx.Next()
			return
		}

		authorization := ctx.GetHeader(AuthorizationHeader)

		if len(authorization) == 0 {
			response.ErrorResponse(ctx, "please provide authorization", 401)
			return
		}

		fields := strings.Fields(authorization)

		if len(fields) < 2 {
			response.ErrorResponse(ctx, "invalid format header", 401)
			return
		}

		if fields[0] != AuthorizationType {
			response.ErrorResponse(ctx, "invalid type header", 401)
			return
		}

		payload, err := c.Token.VerifyTokenPaseto(fields[1])

		if err != nil {
			response.ErrorResponse(ctx, "Internal server", 401)
			return
		}

		if slices.Contains(payload.Permissions, permission) || slices.Contains(payload.Permissions, config.CONFIG_PERMISSIONS["ADMIN"].(string)) {
			ctx.Set(AuthorizationKey, payload)
			ctx.Next()
		} else {
			response.ErrorResponse(ctx, "unauthorized", 401)
			return
		}
	}
}
