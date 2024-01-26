package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/mickBoat00/TransactionAPI/models"
	"github.com/mickBoat00/TransactionAPI/sql/database"
	"github.com/mickBoat00/TransactionAPI/utils"
)

func decodeTransactionRequestBody(w http.ResponseWriter, r *http.Request, model models.TransactionRequestParams) (models.TransactionRequestParams, error) {
	decoder := json.NewDecoder(r.Body)

	params := model

	err := decoder.Decode(&params)

	return params, err
}

// ListTransactions godoc
//
//	@Summary		List transaction
//	@Description	get user transactions
//	@Tags			Transaction
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		models.TransactionResponseParams
//	@Failure		400	{object}	models.ErrorJsonParams
//	@Failure		500	{object}	models.ErrorJsonParams
//	@Security		ApiKeyAuth
//	@Router			/transactions/ [get]
func (serverCfg *ServerConfig) ListTransactions(w http.ResponseWriter, r *http.Request, user_id uuid.UUID) {

	transactions, err := serverCfg.DB.GetUserTransactions(r.Context(), user_id)

	if utils.IfErrorRespondWithErrorJson(w, err, http.StatusInternalServerError, fmt.Sprintf("%s", err)) {
		return
	}

	utils.RespondWithJson(w, http.StatusOK, models.DatabaseTransactionsToTransactions(transactions))

}

// CreateCategory godoc
//
//	@Summary		Create a transaction
//	@Description	create by json Transaction
//	@Tags			Transaction
//	@Accept			json
//	@Produce		json
//	@Param			Transaction	body		models.TransactionRequestParams	true	"Add Transaction"
//	@Success		200		{object}	models.TransactionResponseParams
//	@Failure		400	{object}	models.ErrorJsonParams
//	@Failure		500	{object}	models.ErrorJsonParams
//	@Security		ApiKeyAuth
//	@Router			/transactions/ [post]
func (serverCfg *ServerConfig) CreateTransaction(w http.ResponseWriter, r *http.Request, user_id uuid.UUID) {

	params, err := decodeTransactionRequestBody(w, r, models.TransactionRequestParams{})

	if utils.IfErrorRespondWithErrorJson(w, err, http.StatusInternalServerError, fmt.Sprintf("%s", err)) {
		return
	}

	transaction, err := serverCfg.DB.CreateTransaction(r.Context(), database.CreateTransactionParams{
		ID:         uuid.New(),
		CurrencyID: params.CurrencyID,
		CategoryID: params.CategoryID,
		Amount:     fmt.Sprintf("%.2f", params.Amount),
		Date:       params.Date,
		UserID:     user_id,
		Createdat:  time.Now().UTC(),
		Updatedat:  time.Now().UTC(),
	})

	if utils.IfErrorRespondWithErrorJson(w, err, http.StatusInternalServerError, fmt.Sprintf("%s", err)) {
		return
	}

	utils.RespondWithJson(w, http.StatusOK, models.DatabaseTransactionToTransaction(transaction))

}
