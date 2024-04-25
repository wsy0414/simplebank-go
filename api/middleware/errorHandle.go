package middleware

import (
	"net/http"
	"simplebank/customError"

	"github.com/gin-gonic/gin"
)

func HandleError() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
		if len(ctx.Errors) == 0 {
			return
		}
		err := ctx.Errors[0]

		if customError, ok := err.Err.(*customError.CustomError); ok {
			ctx.JSON(customError.StatusCode, errorResponse(customError.Error()))
		} else {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err.Error()))
		}
	}
}

type ErrorResponse struct {
	Msg string `json:"msg"`
}

func errorResponse(msg string) ErrorResponse {
	return ErrorResponse{
		Msg: msg,
	}
}
