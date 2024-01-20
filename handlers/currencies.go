package handlers

import (
	"fmt"
	"net/http"

	"github.com/mickBoat00/TransactionAPI/models"
	"github.com/mickBoat00/TransactionAPI/utils"
)

// ListCurrencies godoc
//
//	@Summary		List currencies
//	@Description	get currencies
//	@Tags			Currency
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		models.Currency
//	@Failure		400	{object}	models.ErrorJsonParams
//	@Failure		500	{object}	models.ErrorJsonParams
//	@Router			/currencies/ [get]
func (serverCfg *ServerConfig) ListCurrencies(w http.ResponseWriter, r *http.Request) {

	currencies, err := serverCfg.DB.GetAllCurrencies(r.Context())

	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("%s", err))
		return
	}

	utils.RespondWithJson(w, http.StatusOK, models.DatabaseCurrenciesToCurrencies(currencies))

}
