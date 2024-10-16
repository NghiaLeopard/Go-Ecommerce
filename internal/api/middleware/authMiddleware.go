package middleware

import (
	"slices"
	"strings"

	"github.com/NghiaLeopard/Go-Ecommerce-Backend/global"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/constant"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/pkg/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (c *middleware) AuthMiddleware(permission string, isAuthMe bool, isPublic bool) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorization := ctx.GetHeader(constant.AuthorizationHeader)

		if authorization != "" {
			if len(authorization) == 0 {
				global.Logger.Error("please provide authorization", zap.String("Status", "Error"))
				response.ErrorResponse(ctx, "please provide authorization", 401)
				ctx.Abort()
				return
			}

			fields := strings.Fields(authorization)

			if len(fields) < 2 {
				global.Logger.Error("invalid format header", zap.String("Status", "Error"))
				response.ErrorResponse(ctx, "invalid format header", 401)
			}

			if fields[0] != constant.AuthorizationType {
				global.Logger.Error("invalid type header", zap.String("Status", "Error"))
				response.ErrorResponse(ctx, "invalid type header", 401)
				ctx.Abort()
				return
			}

			err := c.Redis.CheckBlackListToken(ctx, fields[1])

			if err != nil {
				global.Logger.Error("Token in blacklist", zap.String("Status", "Error"))
				response.ErrorResponse(ctx, "Token in blacklist", 401)
				ctx.Abort()
				return
			}

			payload, err := global.Token.VerifyToken(fields[1])

			if err != nil {
				global.Logger.Error("Verify token invalid", zap.String("Status", "Error"))
				response.ErrorResponse(ctx, err.Error(), 401)
				ctx.Abort()
				return
			}

			if slices.Contains(payload.Permissions, permission) || slices.Contains(payload.Permissions, constant.CONFIG_PERMISSIONS["ADMIN"].(string)) || isAuthMe {
				ctx.Set(constant.AuthorizationKey, payload)
				ctx.Next()
				return
			} else {
				global.Logger.Error("unauthorized", zap.String("Status", "Error"))
				response.ErrorResponse(ctx, "unauthorized", 401)
				ctx.Abort()
				return
			}
		} else if isPublic {
			ctx.Next()
			return
		} else {
			global.Logger.Error("unauthorized", zap.String("Status", "Error"))
			response.ErrorResponse(ctx, "unauthorized", 401)
			ctx.Abort()
			return
		}
	}
}
