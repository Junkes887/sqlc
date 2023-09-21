-- name: CreateUser :exec
INSERT INTO users (id, name) VALUES(?,?);

-- name: GetUser :one
SELECT * FROM users WHERE id = ?;

-- name: GetUsers :many
SELECT * FROM users;