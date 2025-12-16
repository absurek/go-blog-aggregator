-- name: CreateUser :one
INSERT INTO users (id, name, created_at, updated_at)
VALUES ($1, $2, DEFAULT, DEFAULT)
RETURNING *;

-- name: GetUser :one
SELECT id, name, created_at, updated_at
FROM users
WHERE name = $1;

-- name: GetUsers :many
SELECT id, name, created_at, updated_at
FROM users;

-- name: DeleteAllUsers :exec
DELETE FROM users;
