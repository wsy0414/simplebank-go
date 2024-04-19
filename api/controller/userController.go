package controller

import (
	"net/http"
	"simplebank/api/middleware"
	"simplebank/api/model"
	"simplebank/api/service"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	SignUp(*gin.Context)
	Login(*gin.Context)
	GetUserInfo(*gin.Context)
}

type userController struct {
	userService service.UserService
}

func NewUserController(router *gin.Engine, userService service.UserService) {
	controller := &userController{
		userService: userService,
	}
	router.POST("/signup", controller.SignUp)
	router.POST("/login", controller.Login)
	router.GET("/user", middleware.CheckToken(), controller.GetUserInfo)
}

// Signup godoc
//
//	@Summary		SignUp
//	@Description	SignUp a Guest Account
//	@Tags			accounts
//	@Accept			json
//	@Produce		json
//	@Param			payload	body		model.SignUpRequestParam	true	"Account ID"
//	@Success		200		{object}	model.SignUpResponse
//	@Router			/signup [post]
func (uc userController) SignUp(ctx *gin.Context) {
	var param model.SignUpRequestParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err.Error()))
		return
	}

	response, err := uc.userService.SignUp(ctx, &param)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (uc userController) Login(ctx *gin.Context) {
	var param model.LoginRequestParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err.Error()))
		return
	}

	response, err := uc.userService.Login(ctx, &param)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (uc userController) GetUserInfo(ctx *gin.Context) {
	userId := ctx.GetInt("userId")
	response, err := uc.userService.GetUserInfo(ctx, int(userId))
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
