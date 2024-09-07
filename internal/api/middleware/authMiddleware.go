package middleware

import (
	"slices"
	"strings"

	"github.com/NghiaLeopard/Go-Ecommerce-Backend/global"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/constant"
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

		if isPublic {
			ctx.Next()
		}

		authorization := ctx.GetHeader(AuthorizationHeader)

		if len(authorization) == 0 {
			response.ErrorResponse(ctx, "please provide authorization", 401)
		}

		fields := strings.Fields(authorization)

		if len(fields) < 2 {
			response.ErrorResponse(ctx, "invalid format header", 401)
		}

		if fields[0] != AuthorizationType {
			response.ErrorResponse(ctx, "invalid type header", 401)

		}

		payload, err := global.Token.VerifyTokenPaseto(fields[1])

		if err != nil {
			response.ErrorResponse(ctx, "Verify token invalid", 401)
		}

		if slices.Contains(payload.Permissions, permission) || slices.Contains(payload.Permissions, constant.CONFIG_PERMISSIONS["ADMIN"].(string)) || isAuthMe {
			ctx.Set(AuthorizationKey, payload)
			ctx.Next()
			ctx.Abort()
		} else {
			response.ErrorResponse(ctx, "unauthorized", 401)
			ctx.Abort()
		}
	}
}
