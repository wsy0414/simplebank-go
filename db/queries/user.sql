-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: CreateUser :one
insert into users(
    name,
    password,
    email,
    birthdate
) values($1, $2, $3, $4)
RETURNING *;

-- name: GetUserByName :one
SELECT * FROM users
WHERE name = $1;