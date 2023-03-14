package utils

import "github.com/gin-gonic/gin"

type Responses struct {
	StatusCode int         `json:"statusCode"`
	Method     string      `json:"method"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

type ErrResponse struct {
	StatusCode int         `json:"statusCode"`
	Method     string      `json:"method"`
	Error      interface{} `json:"error"`
}

func APIResponse(ctx *gin.Context, message string, statusCode int, method string, data interface{}) {
	jsonResponse := Responses{
		StatusCode: statusCode,
		Method:     method,
		Message:    message,
		Data:       data,
	}

	if statusCode > 400 {
		ctx.JSON(statusCode, jsonResponse)
		defer ctx.AbortWithStatus(statusCode)
	} else {
		ctx.JSON(statusCode, jsonResponse)
	}
}

func ValidatorErrorResponse(ctx *gin.Context, message string, statusCode int, method string, Error interface{}) {
	errorResponse := ErrResponse{
		StatusCode: statusCode,
		Method:     method,
		Error:      Error,
	}

	ctx.JSON(statusCode, errorResponse)
	defer ctx.AbortWithStatus(statusCode)
}
