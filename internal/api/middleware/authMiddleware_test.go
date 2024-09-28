package middleware

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/NghiaLeopard/Go-Ecommerce-Backend/global"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/constant"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/internal/repository"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/pkg/response"
	"github.com/NghiaLeopard/Go-Ecommerce-Backend/pkg/token"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func AddAuthorization(t *testing.T, request *http.Request, makerToken token.Maker, authorizationHeader string, authorizationType string, id int, permission []string, duration time.Duration) {
	token, payload, err := makerToken.CreateTokenPaseto(id, permission, duration)

	require.NoError(t, err)
	require.NotEmpty(t, payload)
	require.NotZero(t, token)

	requestHeader := fmt.Sprintf("%s %s", authorizationType, token)

	request.Header.Set(authorizationHeader, requestHeader)
}

func TestAuthMiddleware(t *testing.T) {
	testCase := []struct {
		name          string
		caseGet       func(router *gin.Engine, middleware Middleware)
		setupAuth     func(t *testing.T, request *http.Request, makerToken token.Maker, authorizationHeader string, authorizationType string, id int, permission []string, duration time.Duration)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			caseGet: func(router *gin.Engine, middleware Middleware) {
				router.GET("/auth", middleware.AuthMiddleware(constant.CONFIG_PERMISSIONS["ADMIN"].(string), false, false), func(ctx *gin.Context) {
					response.SuccessResponse(ctx, "Success", 200, nil)
				})
			},
			setupAuth: func(t *testing.T, request *http.Request, makerToken token.Maker, authorizationHeader string, authorizationType string, id int, permission []string, duration time.Duration) {
				AddAuthorization(t, request, makerToken, authorizationHeader, authorizationType, id, permission, duration)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name: "is auth me",
			caseGet: func(router *gin.Engine, middleware Middleware) {
				router.GET("/auth", middleware.AuthMiddleware("1", true, false), func(ctx *gin.Context) {
					response.SuccessResponse(ctx, "Success", 200, nil)
				})
			},
			setupAuth: func(t *testing.T, request *http.Request, makerToken token.Maker, authorizationHeader string, authorizationType string, id int, permission []string, duration time.Duration) {
				AddAuthorization(t, request, makerToken, authorizationHeader, authorizationType, id, permission, duration)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name: "isPublic",
			caseGet: func(router *gin.Engine, middleware Middleware) {
				router.GET("/auth", middleware.AuthMiddleware("1", false, true), func(ctx *gin.Context) {
					response.SuccessResponse(ctx, "Success", 200, nil)
				})
			},
			setupAuth: func(t *testing.T, request *http.Request, makerToken token.Maker, authorizationHeader string, authorizationType string, id int, permission []string, duration time.Duration) {
				AddAuthorization(t, request, makerToken, authorizationHeader, authorizationType, id, []string{"Test"}, duration)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name: "unauthorized",
			caseGet: func(router *gin.Engine, middleware Middleware) {
				router.GET("/auth", middleware.AuthMiddleware("1", false, false), func(ctx *gin.Context) {
					response.SuccessResponse(ctx, "Success", 200, nil)
				})
			},
			setupAuth: func(t *testing.T, request *http.Request, makerToken token.Maker, authorizationHeader string, authorizationType string, id int, permission []string, duration time.Duration) {
				AddAuthorization(t, request, makerToken, authorizationHeader, authorizationType, id, []string{"Test"}, duration)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name: "provide authorization",
			caseGet: func(router *gin.Engine, middleware Middleware) {
				router.GET("/auth", middleware.AuthMiddleware("", false, true), func(ctx *gin.Context) {
					response.SuccessResponse(ctx, "Success", 200, nil)
				})
			},
			setupAuth: func(t *testing.T, request *http.Request, makerToken token.Maker, authorizationHeader string, authorizationType string, id int, permission []string, duration time.Duration) {
				AddAuthorization(t, request, makerToken, "", authorizationType, id, permission, duration)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name: "invalid format",
			caseGet: func(router *gin.Engine, middleware Middleware) {
				router.GET("/auth", middleware.AuthMiddleware("", false, true), func(ctx *gin.Context) {
					response.SuccessResponse(ctx, "Success", 200, nil)
				})
			},
			setupAuth: func(t *testing.T, request *http.Request, makerToken token.Maker, authorizationHeader string, authorizationType string, id int, permission []string, duration time.Duration) {
				AddAuthorization(t, request, makerToken, authorizationHeader, "", id, permission, duration)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name: "invalid type",
			caseGet: func(router *gin.Engine, middleware Middleware) {
				router.GET("/auth", middleware.AuthMiddleware("", false, true), func(ctx *gin.Context) {
					response.SuccessResponse(ctx, "Success", 200, nil)
				})
			},
			setupAuth: func(t *testing.T, request *http.Request, makerToken token.Maker, authorizationHeader string, authorizationType string, id int, permission []string, duration time.Duration) {
				AddAuthorization(t, request, makerToken, authorizationHeader, "1", id, permission, duration)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name: "token is expired",
			caseGet: func(router *gin.Engine, middleware Middleware) {
				router.GET("/auth", middleware.AuthMiddleware("", false, true), func(ctx *gin.Context) {
					response.SuccessResponse(ctx, "Success", 200, nil)
				})
			},
			setupAuth: func(t *testing.T, request *http.Request, makerToken token.Maker, authorizationHeader string, authorizationType string, id int, permission []string, duration time.Duration) {
				AddAuthorization(t, request, makerToken, authorizationHeader, authorizationType, id, permission, -duration)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
	}

	for i := range testCase {
		tc := testCase[i]
		t.Run(tc.name, func(t *testing.T) {
			router := gin.Default()

			redisRepo := repository.NewRedisTokenRepository(global.Rdb)

			middleware := NewMiddleware(redisRepo)

			recorder := httptest.NewRecorder()

			request, err := http.NewRequest(http.MethodGet, "/auth", nil)
			tc.caseGet(router, middleware)

			tc.setupAuth(t, request, global.Token, AuthorizationHeader, AuthorizationType, 1, []string{constant.CONFIG_PERMISSIONS["ADMIN"].(string)}, 10*time.Hour)

			require.NoError(t, err)

			router.ServeHTTP(recorder, request)

			tc.checkResponse(t, recorder)
		})
	}

}
