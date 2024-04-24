package model

import "time"

type SignUpResponse struct {
	ID    int    `json:"id"`
	Token string `json:"token"`
}

type LoginResponse struct {
	ID    int    `json:"id"`
	Token string `json:"token"`
}

type GetUserInfoResponse struct {
	ID        int        `json:"id"`
	Name      string     `json:"name"`
	Birthdate CustomTime `json:"birthdate"`
	Email     string     `json:"email"`
	Balance   []Balance  `json:"balance"`
}

type BalanceResponse struct {
	Currency string `json:"currency"`
	Balance  string `json:"balance"`
}

type CreateBalanceResponse struct {
	BalanceResponse
}

type GetBalanceResponse struct {
	BalanceResponse
}

type Balance struct {
	Currency string `json:"currency"`
	Balance  string `json:"balance"`
}

type DepositeResponse struct {
	BalanceResponse
}

type WithdrawResponse struct {
	BalanceResponse
}

type TransferResponse struct {
	BalanceResponse
}

type ListActivityResponse struct {
	UserId   int       `json:"userId"`
	Amount   string    `json:"amount"`
	CreateAt time.Time `json:"createAt"`
}
