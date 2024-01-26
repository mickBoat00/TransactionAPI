// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: transactions.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createTransaction = `-- name: CreateTransaction :one
INSERT INTO transactions (id, currency_id, category_id, amount, date, user_id, createdAt, updatedAt)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING id, currency_id, category_id, amount, date, user_id, createdat, updatedat
`

type CreateTransactionParams struct {
	ID         uuid.UUID
	CurrencyID uuid.UUID
	CategoryID uuid.UUID
	Amount     string
	Date       time.Time
	UserID     uuid.UUID
	Createdat  time.Time
	Updatedat  time.Time
}

func (q *Queries) CreateTransaction(ctx context.Context, arg CreateTransactionParams) (Transaction, error) {
	row := q.db.QueryRowContext(ctx, createTransaction,
		arg.ID,
		arg.CurrencyID,
		arg.CategoryID,
		arg.Amount,
		arg.Date,
		arg.UserID,
		arg.Createdat,
		arg.Updatedat,
	)
	var i Transaction
	err := row.Scan(
		&i.ID,
		&i.CurrencyID,
		&i.CategoryID,
		&i.Amount,
		&i.Date,
		&i.UserID,
		&i.Createdat,
		&i.Updatedat,
	)
	return i, err
}

const getUserTransactions = `-- name: GetUserTransactions :many
SELECT id, currency_id, category_id, amount, date, user_id, createdat, updatedat FROM transactions WHERE user_id = $1
`

func (q *Queries) GetUserTransactions(ctx context.Context, userID uuid.UUID) ([]Transaction, error) {
	rows, err := q.db.QueryContext(ctx, getUserTransactions, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Transaction
	for rows.Next() {
		var i Transaction
		if err := rows.Scan(
			&i.ID,
			&i.CurrencyID,
			&i.CategoryID,
			&i.Amount,
			&i.Date,
			&i.UserID,
			&i.Createdat,
			&i.Updatedat,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
