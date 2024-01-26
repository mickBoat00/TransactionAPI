-- name: GetUserTransactions :many
SELECT * FROM transactions WHERE user_id = $1;

-- name: CreateTransaction :one
INSERT INTO transactions (id, currency_id, category_id, amount, date, user_id, createdAt, updatedAt)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;