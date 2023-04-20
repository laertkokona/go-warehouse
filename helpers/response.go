package helpers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type JSONResult struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type JSONSuccessResult struct {
	Code    int         `json:"code" example:"200"`
	Message string      `json:"message" example:"Success"`
	Data    interface{} `json:"data,omitempty" example:"{name: 'John', age: 30}"`
}

type JSONSuccessResultNoData struct {
	Code    int    `json:"code" example:"200"`
	Message string `json:"message" example:"Success"`
}

type JSONBadRequestResult struct {
	Code    int         `json:"code" example:"400"`
	Message string      `json:"message" example:"Wrong request"`
	Data    interface{} `json:"data,omitempty"`
}

type JSONNotFoundResult struct {
	Code    int         `json:"code" example:"404"`
	Message string      `json:"message" example:"Not found"`
	Data    interface{} `json:"data,omitempty"`
}

type JSONUnauthorizedResult struct {
	Code    int         `json:"code" example:"401"`
	Message string      `json:"message" example:"Unauthorized"`
	Data    interface{} `json:"data,omitempty"`
}

type JSONForbiddenResult struct {
	Code    int         `json:"code" example:"403"`
	Message string      `json:"message" example:"Forbidden"`
	Data    interface{} `json:"data,omitempty"`
}

type JSONInternalServerErrorResult struct {
	Code    int         `json:"code" example:"500"`
	Message string      `json:"message" example:"Internal server error"`
	Data    interface{} `json:"data,omitempty"`
}

func SuccessResponse(ctx *gin.Context, data interface{}) {
	if data == "" {
		data = "Success"
	}
	ctx.JSON(http.StatusOK, JSONSuccessResult{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    data,
	})
}

func SuccessResponseNoData(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, JSONSuccessResultNoData{
		Code:    http.StatusOK,
		Message: "Success",
	})
}

func FailedResponse(ctx *gin.Context, respCode int, message string, data interface{}) {
	switch respCode {
	case http.StatusBadRequest:
		ctx.JSON(respCode, JSONBadRequestResult{
			Code:    http.StatusBadRequest,
			Message: message,
			Data:    data,
		})
		return
	case http.StatusNotFound:
		ctx.JSON(respCode, JSONNotFoundResult{
			Code:    http.StatusNotFound,
			Message: message,
			Data:    data,
		})
		return
	case http.StatusUnauthorized:
		ctx.JSON(respCode, JSONUnauthorizedResult{
			Code:    http.StatusUnauthorized,
			Message: message,
			Data:    data,
		})
		return
	case http.StatusForbidden:
		ctx.JSON(respCode, JSONForbiddenResult{
			Code:    http.StatusForbidden,
			Message: message,
			Data:    data,
		})
		return
	case http.StatusInternalServerError:
		ctx.JSON(respCode, JSONInternalServerErrorResult{
			Code:    http.StatusInternalServerError,
			Message: message,
			Data:    data,
		})
		return
	}
}
