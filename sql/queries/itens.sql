-- name: CreateIten :exec
INSERT INTO itens (id, name, user_id) VALUES(?,?,?);

-- name: GetIten :one
SELECT * FROM itens WHERE id = ?;

-- name: GetItens :many
SELECT * FROM itens;