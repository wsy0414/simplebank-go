package controller

import (
	"net/http"
	"simplebank/api/middleware"
	"simplebank/api/model"
	"simplebank/api/service"
	"simplebank/customError"

	"github.com/gin-gonic/gin"
)

type BalanceController interface {
	CreateBalance(*gin.Context)
	GetBalance(*gin.Context)
	ListBalance(*gin.Context)
	Deposite(*gin.Context)
	Withdraw(*gin.Context)
	Transfer(*gin.Context)
	GetActivity(*gin.Context)
}

type balanceController struct {
	balanceService service.BalanceService
}

func NewBalanceController(router *gin.Engine, service service.BalanceService) {
	controller := &balanceController{
		balanceService: service,
	}

	routerGroup := router.Group("/balance", middleware.CheckToken(), middleware.HandleError())

	routerGroup.POST("/", controller.CreateBalance)
	routerGroup.GET("/:currency", controller.GetBalance)
	routerGroup.GET("/list", controller.ListBalance)
	routerGroup.POST("/deposite", controller.Deposite)
	routerGroup.POST("/withdraw", controller.Withdraw)
	routerGroup.POST("/transfer", controller.Transfer)
	routerGroup.GET("/activity", controller.GetActivity)
}

// CreateBalance godoc
//
//	@Summary		CreateBalance
//	@Description	CreateBalance
//	@Tags			balance
//	@Accept			json
//	@Produce		json
//	@Param			authorization	header		string						true	"jwtToken"
//	@Param			payload			body		model.CreateBalanceRequest	true	"CreateBalanceRequest"
//	@Success		200				{object}	model.CreateBalanceResponse
//	@Router			/balance [post]
func (controller balanceController) CreateBalance(ctx *gin.Context) {
	var param model.CreateBalanceRequest
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.Error(customError.NewBadRequestError(err))
		return
	}

	response, err := controller.balanceService.CreateBalance(ctx, param)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// GetBalance godoc
//
//	@Summary		GetBalance
//	@Description	get specify balance by user
//	@Tags			balance
//	@Accept			json
//	@Produce		json
//	@Param			authorization	header		string	true	"jwtToken"
//	@Param			currency		path		string	true	"currency(TWD||USD||JPN)"
//	@Success		200				{object}	model.GetBalanceResponse
//	@Router			/balance/{currency} [get]
func (controller balanceController) GetBalance(ctx *gin.Context) {
	currency := ctx.Param("currency")

	response, err := controller.balanceService.GetBalance(ctx, currency)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// ListBalance godoc
//
//	@Summary		ListBalance
//	@Description	get all balance by user
//	@Tags			balance
//	@Accept			json
//	@Produce		json
//	@Param			authorization	header		string	true	"jwtToken"
//	@Success		200				{object}	[]model.GetBalanceResponse
//	@Router			/balance/list [get]
func (controller balanceController) ListBalance(ctx *gin.Context) {
	response, err := controller.balanceService.ListBalance(ctx)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// Deposite godoc
//
//	@Summary		Deposite
//	@Description	deposite amount into specify currency balance
//	@Tags			balance
//	@Accept			json
//	@Produce		json
//	@Param			authorization	header		string						true	"jwtToken"
//	@Param			payload			body		model.DepositeRequestParam	true	"DepositeRequestParam"
//	@Success		200				{object}	model.DepositeResponse
//	@Router			/balance/deposite [post]
func (controller balanceController) Deposite(ctx *gin.Context) {
	var param model.DepositeRequestParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.Error(customError.NewBadRequestError(err))
		return
	}

	response, err := controller.balanceService.Deposite(ctx, param)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// Withdraw godoc
//
//	@Summary		Withdraw
//	@Description	withdraw amount from specify currency balance
//	@Tags			balance
//	@Accept			json
//	@Produce		json
//	@Param			authorization	header		string						true	"jwtToken"
//	@Param			payload			body		model.WithdrawRequestParam	true	"WithdrawRequestParam"
//	@Success		200				{object}	model.WithdrawResponse
//	@Router			/balance/withdraw [post]
func (controller balanceController) Withdraw(ctx *gin.Context) {
	var param model.WithdrawRequestParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.Error(customError.NewBadRequestError(err))
		return
	}

	response, err := controller.balanceService.Withdraw(ctx, param)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// Transfer godoc
//
//	@Summary		Transfer
//	@Description	transfer amount from specify currency balance to other user's  balance
//	@Tags			balance
//	@Accept			json
//	@Produce		json
//	@Param			authorization	header		string						true	"jwtToken"
//	@Param			payload			body		model.TransferRequestParam	true	"TransferRequestParam"
//	@Success		200				{object}	model.TransferResponse
//	@Router			/balance/transfer [post]
func (controller balanceController) Transfer(ctx *gin.Context) {
	var param model.TransferRequestParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.Error(customError.NewBadRequestError(err))
		return
	}

	response, err := controller.balanceService.Transfer(ctx, param)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// GetActivity godoc
//
//	@Summary		GetActivity
//	@Description	get activity log by user
//	@Tags			balance
//	@Accept			json
//	@Produce		json
//	@Param			authorization	header		string	true	"jwtToken"
//	@Success		200				{object}	[]model.ListActivityResponse
//	@Router			/balance/activity [get]
func (controller balanceController) GetActivity(ctx *gin.Context) {
	response, err := controller.balanceService.GetActivity(ctx)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}
