-- name: CreateTransferLog :one
insert into transfer_log(from_user_id, to_user_id, amount)
values($1, $2, $3)
RETURNING *;