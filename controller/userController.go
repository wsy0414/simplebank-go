package controller

import (
	"net/http"
	"simplebank/model"
	"simplebank/service"

	"github.com/gin-gonic/gin"
)

func SignUp(ctx *gin.Context) {
	var param model.SignUpRequestParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err.Error()))
		return
	}

	response, err := service.SignUp(&param)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, response)
}

type ErrorResponse struct {
	Msg string `json:"msg"`
}

func errorResponse(msg string) ErrorResponse {
	return ErrorResponse{
		Msg: msg,
	}
}
