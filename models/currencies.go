package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/mickBoat00/TransactionAPI/sql/database"
)

type Currency struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Code      string    `json:"code"`
	Createdat time.Time `json:"created_at"`
	Updatedat time.Time `json:"updated_at"`
}

func databaseCurrencyToCurrency(currency database.Currency) Currency {
	return Currency{
		ID:        currency.ID,
		Name:      currency.Name,
		Code:      currency.Code,
		Createdat: currency.Createdat,
		Updatedat: currency.Updatedat,
	}
}

func DatabaseCurrenciesToCurrencies(currencies []database.Currency) []Currency {
	currencySlice := make([]Currency, 0)

	for _, currency := range currencies {
		currencySlice = append(currencySlice, databaseCurrencyToCurrency(currency))
	}

	return currencySlice
}
