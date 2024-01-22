package handlers

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/mickBoat00/TransactionAPI/models"
	"github.com/mickBoat00/TransactionAPI/utils"
)

// ListCategories godoc
//
//	@Summary		List categories
//	@Description	get user categories
//	@Tags			Category
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		models.Category
//	@Failure		400	{object}	models.ErrorJsonParams
//	@Failure		500	{object}	models.ErrorJsonParams
//	@Router			/categories/ [get]
func (serverCfg *ServerConfig) ListCategories(w http.ResponseWriter, r *http.Request) {

	categories, err := serverCfg.DB.GetUserCategories(r.Context(), uuid.New())

	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("%s", err))
		return
	}

	utils.RespondWithJson(w, http.StatusOK, models.DatabaseCategoriesToCategories(categories))

}
