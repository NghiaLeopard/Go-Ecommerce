package middleware

import (
	"slices"
	"strings"

	"github.com/NghiaLeopard/Go-Ecommerce-Backend/pkg/config"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/pkg/response"
	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "authorization"
	authorizationType   = "Bearer"
	authorizationKey    = "authorization_payload"
)

func (c *middleware) AuthMiddleware(permission string, isAuthMe bool, isPublic bool) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorization := ctx.GetHeader(authorizationHeader)

		if len(authorization) == 0 {
			response.ErrorResponse(ctx, "please provide authorization", 401)
		}

		fields := strings.Fields(authorization)

		if len(fields) < 2 {
			response.ErrorResponse(ctx, "invalid format header", 401)
		}

		if fields[0] != authorizationType {
			response.ErrorResponse(ctx, "invalid type header", 401)
		}

		payload, err := c.Token.VerifyTokenPaseto(fields[1])

		if err != nil {
			response.ErrorResponse(ctx, "Internal server", 401)
		}

		if slices.Contains(payload.Permissions, permission) || slices.Contains(payload.Permissions, config.CONFIG_PERMISSIONS["ADMIN"].(string)) || isAuthMe {
			ctx.Set(authorizationKey, payload)
			ctx.Next()
		} else if isPublic {
			ctx.Next()
		}

		response.ErrorResponse(ctx, "unauthorized", 401)
	}
}
