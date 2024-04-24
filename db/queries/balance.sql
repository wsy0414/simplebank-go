-- name: GetBalanceByUser :many
SELECT * FROM balance
WHERE user_id = $1
ORDER BY currency;

-- name: GetSpecifyCurrencyBalanceByUser :one
SELECT * FROM balance
WHERE user_id = $1
AND currency = $2;

-- name: CreateBalance :one
insert into balance(user_id, currency, balance)
values($1, $2, $3)
RETURNING *;

-- name: AddBalance :one
update balance set
    balance = balance + $1
where
    user_id = $2
    and currency = $3
RETURNING *;

-- name: SubBalance :one
update balance set
    balance = balance - $1
where
    user_id = $2
    and currency = $3
RETURNING *;

-- name: InserOrUpdateBalance :one
insert into balance(user_id, currency, balance)
values($1, $2, $3)
on conflict on constraint uk_user_id_currency
do update set balance = excluded.balance
RETURNING *;