package response

import (
	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
	Status    string      `json:"status"`
	TypeError string      `json:"typeError"`
}

func SuccessResponse(ctx *gin.Context, message string, code int, data interface{}) {
	var rsp = ResponseData{
		Message:   message,
		Data:      data,
		Status:    "Success",
		TypeError: httpResponse[code].Type,
	}

	ctx.JSON(httpResponse[code].Status, rsp)
}

func ErrorResponse(ctx *gin.Context, message string, code int) {
	var rsp = ResponseData{
		Message:   message,
		Data:      "",
		Status:    "Error",
		TypeError: httpResponse[code].Type,
	}

	ctx.JSON(httpResponse[code].Status, rsp)
}
