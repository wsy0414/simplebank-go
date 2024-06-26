// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: balance.sql

package sqlc

import (
	"context"
)

const addBalance = `-- name: AddBalance :one
update balance set
    balance = balance + $1
where
    user_id = $2
    and currency = $3
RETURNING id, user_id, currency, balance, created_at
`

type AddBalanceParams struct {
	Balance  string
	UserID   int32
	Currency string
}

func (q *Queries) AddBalance(ctx context.Context, arg AddBalanceParams) (Balance, error) {
	row := q.db.QueryRowContext(ctx, addBalance, arg.Balance, arg.UserID, arg.Currency)
	var i Balance
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Currency,
		&i.Balance,
		&i.CreatedAt,
	)
	return i, err
}

const createBalance = `-- name: CreateBalance :one
insert into balance(user_id, currency, balance)
values($1, $2, $3)
RETURNING id, user_id, currency, balance, created_at
`

type CreateBalanceParams struct {
	UserID   int32
	Currency string
	Balance  string
}

func (q *Queries) CreateBalance(ctx context.Context, arg CreateBalanceParams) (Balance, error) {
	row := q.db.QueryRowContext(ctx, createBalance, arg.UserID, arg.Currency, arg.Balance)
	var i Balance
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Currency,
		&i.Balance,
		&i.CreatedAt,
	)
	return i, err
}

const getBalanceByUser = `-- name: GetBalanceByUser :many
SELECT id, user_id, currency, balance, created_at FROM balance
WHERE user_id = $1
ORDER BY currency
`

func (q *Queries) GetBalanceByUser(ctx context.Context, userID int32) ([]Balance, error) {
	rows, err := q.db.QueryContext(ctx, getBalanceByUser, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Balance
	for rows.Next() {
		var i Balance
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Currency,
			&i.Balance,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getSpecifyCurrencyBalanceByUser = `-- name: GetSpecifyCurrencyBalanceByUser :one
SELECT id, user_id, currency, balance, created_at FROM balance
WHERE user_id = $1
AND currency = $2
`

type GetSpecifyCurrencyBalanceByUserParams struct {
	UserID   int32
	Currency string
}

func (q *Queries) GetSpecifyCurrencyBalanceByUser(ctx context.Context, arg GetSpecifyCurrencyBalanceByUserParams) (Balance, error) {
	row := q.db.QueryRowContext(ctx, getSpecifyCurrencyBalanceByUser, arg.UserID, arg.Currency)
	var i Balance
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Currency,
		&i.Balance,
		&i.CreatedAt,
	)
	return i, err
}

const inserOrUpdateBalance = `-- name: InserOrUpdateBalance :one
insert into balance(user_id, currency, balance)
values($1, $2, $3)
on conflict on constraint uk_user_id_currency
do update set balance = excluded.balance
RETURNING id, user_id, currency, balance, created_at
`

type InserOrUpdateBalanceParams struct {
	UserID   int32
	Currency string
	Balance  string
}

func (q *Queries) InserOrUpdateBalance(ctx context.Context, arg InserOrUpdateBalanceParams) (Balance, error) {
	row := q.db.QueryRowContext(ctx, inserOrUpdateBalance, arg.UserID, arg.Currency, arg.Balance)
	var i Balance
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Currency,
		&i.Balance,
		&i.CreatedAt,
	)
	return i, err
}

const subBalance = `-- name: SubBalance :one
update balance set
    balance = balance - $1
where
    user_id = $2
    and currency = $3
RETURNING id, user_id, currency, balance, created_at
`

type SubBalanceParams struct {
	Balance  string
	UserID   int32
	Currency string
}

func (q *Queries) SubBalance(ctx context.Context, arg SubBalanceParams) (Balance, error) {
	row := q.db.QueryRowContext(ctx, subBalance, arg.Balance, arg.UserID, arg.Currency)
	var i Balance
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Currency,
		&i.Balance,
		&i.CreatedAt,
	)
	return i, err
}
