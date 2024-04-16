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
