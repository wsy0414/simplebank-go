// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package sqlc

import (
	"database/sql"
	"time"
)

type Activity struct {
	ID         int32
	FromUserID int32
	ToUserID   int32
	// just postive
	Amount    string
	CreatedAt time.Time
}

type Balance struct {
	ID        int32
	UserID    int32
	Currency  string
	Balance   string
	CreatedAt time.Time
}

type Transfer struct {
	ID     int32
	UserID int32
	// postive or negative
	Amount    string
	CreatedAt time.Time
}

type User struct {
	ID int32
	// 帳號
	Name string
	// 密碼
	Password  string
	Email     string
	Birthdate sql.NullTime
	CreatedAt time.Time
}
