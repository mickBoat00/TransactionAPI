-- name: GetAllUsers :many
SELECT * FROM users;

-- name: CreateUser :one
INSERT INTO users (id, email, password, createdAt, updatedAt)
VALUES ($1, $2, $3, $4, $5) 
RETURNING *;

-- name: GetUserViaEmail :one
SELECT * FROM users WHERE email = $1;