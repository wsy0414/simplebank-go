-- name: GetAccount :one
SELECT * FROM accounts
WHERE id = $1 LIMIT 1;

-- name: ListAccount :many
SELECT * FROM accounts
LIMIT $1
OFFSET $2;

-- name: CreateAccount :one
INSERT INTO accounts(
    name, password, balance, currency
) VALUES(
    $1, $2, $3, $4
)
RETURNING *;

-- name: UpdateAccount :exec
UPDATE accounts SET balance = $1
WHERE id = $2;