-- name: GetBalanceByUser :many
SELECT * FROM balance
WHERE user_id = $1
ORDER BY currency;