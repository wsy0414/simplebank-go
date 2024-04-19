package controller

import (
	"net/http"
	"simplebank/api/middleware"
	"simplebank/api/model"
	"simplebank/api/service"

	"github.com/gin-gonic/gin"
)

type ActivityController interface {
	Deposite(*gin.Context)
	Withdraw(*gin.Context)
	Transfer(*gin.Context)
	ListActivities(*gin.Context)
	ListTransfers(*gin.Context)
}

type activityController struct {
	acrivityService service.ActivityService
}

func NewActivityController(activityService service.ActivityService, router *gin.Engine) {
	controller := activityController{
		acrivityService: activityService,
	}

	router.POST("/deposite", middleware.CheckToken(), controller.Deposite)
	router.POST("/withdraw", middleware.CheckToken(), controller.Withdraw)
	router.POST("/transfer", middleware.CheckToken(), controller.Transfer)
	router.POST("/activity/list", middleware.CheckToken(), controller.ListActivities)
	router.POST("/transfer/list", middleware.CheckToken(), controller.ListTransfers)
}

func (a activityController) Deposite(ctx *gin.Context) {
	var param model.DepositeRequestParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err.Error()))
		return
	}

	response, err := a.acrivityService.Deposite(ctx, ctx.GetInt("userId"), param)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (a activityController) Withdraw(ctx *gin.Context) {}

func (a activityController) Transfer(ctx *gin.Context) {}

func (a activityController) ListActivities(ctx *gin.Context) {}

func (a activityController) ListTransfers(ctx *gin.Context) {}
