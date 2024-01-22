-- name: CreateUserCategory :one
INSERT INTO categories (id, name, user_id, createdAt, updatedAt)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetUserCategories :many
SELECT * FROM categories WHERE user_id = $1;