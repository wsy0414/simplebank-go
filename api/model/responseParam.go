package model

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

type Balance struct {
	Currency string `json:"currency"`
	Balance  string `json:"balance"`
}

type DepositeResponse struct {
	Currency string `json:"currency"`
	Balance  string `json:"balance"`
}
