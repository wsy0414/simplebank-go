package model

type SignUpRequestParam struct {
	// 帳號
	Name string `json:"name" binding:"required"`
	// 密碼
	Password string `json:"password" binding:"required"`
	// 幣種
	Email string `json:"email" binding:"required"`
}
