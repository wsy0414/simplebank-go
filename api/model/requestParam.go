package model

import "time"

type SignUpRequestParam struct {
	// 帳號
	Name string `json:"name" binding:"required" validate:"required"`
	// 密碼
	Password  string    `json:"password" binding:"required,passwordValidate"`
	Birthdate time.Time `json:"birthdate"`
	// 幣種
	Email string `json:"email" binding:"required,email"`
}

type LoginRequestParam struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required,passwordValidate"`
}

type CreateBalanceRequest struct {
	Currency string  `json:"currency" binding:"required,currencyValidate"`
	Amount   float64 `json:"amount"`
}

type BalanceRequest struct {
	Currency string  `json:"currency" binding:"required,currencyValidate"`
	Amount   float64 `json:"amount" binding:"required,gt=0"`
}

type DepositeRequestParam struct {
	BalanceRequest
}

type WithdrawRequestParam struct {
	BalanceRequest
}

type TransferRequestParam struct {
	ToUserId int `json:"toUserId" binding:"required"`
	BalanceRequest
}
