-- name: CreateActivity :one
insert into activity_log(user_id, amount)
values($1, $2)
RETURNING *;

-- name: GetActivity :many
select *
from activity_log
where user_id = $1;