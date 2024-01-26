package models

import (
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/mickBoat00/TransactionAPI/sql/database"
)

type TransactionRequestParams struct {
	CurrencyID uuid.UUID `json:"currency_id"`
	CategoryID uuid.UUID `json:"category_id"`
	Amount     float64   `json:"amount"`
	Date       time.Time `json:"date"`
}

type TransactionResponseParams struct {
	ID         uuid.UUID `json:"id"`
	CurrencyID uuid.UUID `json:"currency_id"`
	CategoryID uuid.UUID `json:"category_id"`
	Amount     float64   `json:"amount"`
	Date       time.Time `json:"date"`
	UserID     uuid.UUID `json:"user_id"`
	Createdat  time.Time `json:"created_at"`
	Updatedat  time.Time `json:"updated_at"`
}

func DatabaseTransactionToTransaction(transaction database.Transaction) TransactionResponseParams {
	amountInDecimal, _ := strconv.ParseFloat(transaction.Amount, 64)

	return TransactionResponseParams{
		ID:         transaction.ID,
		CurrencyID: transaction.CurrencyID,
		CategoryID: transaction.CategoryID,
		Amount:     amountInDecimal,
		Date:       transaction.Date,
		UserID:     transaction.UserID,
		Createdat:  transaction.Createdat,
		Updatedat:  transaction.Updatedat,
	}
}

func DatabaseTransactionsToTransactions(transactions []database.Transaction) []TransactionResponseParams {
	transactionSlice := make([]TransactionResponseParams, 0)

	for _, transaction := range transactions {
		transactionSlice = append(transactionSlice, DatabaseTransactionToTransaction(transaction))
	}

	return transactionSlice
}
