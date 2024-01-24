-- name: CreateUserCategory :one
INSERT INTO categories (id, name, user_id, createdAt, updatedAt)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetUserCategories :many
SELECT * FROM categories WHERE user_id = $1;

-- name: GetUserCategoryIds :many
SELECT id FROM categories WHERE user_id = $1;

-- name: UpdateUserCategory :one
UPDATE categories SET name = $1 , updatedAt = $2
WHERE user_id = $3
RETURNING *;

-- name: DeleteUserCategories :exec
DELETE FROM categories WHERE id = $1 AND user_id = $2;