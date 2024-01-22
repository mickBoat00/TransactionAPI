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
//	@Security		ApiKeyAuth
//	@Router			/categories/ [get]
func (serverCfg *ServerConfig) ListCategories(w http.ResponseWriter, r *http.Request, user_id uuid.UUID) {

	categories, err := serverCfg.DB.GetUserCategories(r.Context(), user_id)

	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("%s", err))
		return
	}

	utils.RespondWithJson(w, http.StatusOK, models.DatabaseCategoriesToCategories(categories))

}

// CreateCategory godoc
//
//	@Summary		Create a category
//	@Description	create by json Category
//	@Tags			Category
//	@Accept			json
//	@Produce		json
//	@Param			Category	body		models.CategoryRequestParams	true	"Add Category"
//	@Success		200		{object}	models.Category
//	@Failure		400	{object}	models.ErrorJsonParams
//	@Failure		500	{object}	models.ErrorJsonParams
//	@Security		ApiKeyAuth
//	@Router			/categories/ [post]
func (serverCfg *ServerConfig) CreateCategory(w http.ResponseWriter, r *http.Request, user_id uuid.UUID) {

	decoder := json.NewDecoder(r.Body)

	params := models.CategoryRequestParams{}

	err := decoder.Decode(&params)

	if err != nil || params.Name == "" {
		utils.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("%s", err))
		return
	}

	category, err := serverCfg.DB.CreateUserCategory(r.Context(), database.CreateUserCategoryParams{
		ID:        uuid.New(),
		Name:      params.Name,
		UserID:    user_id,
		Createdat: time.Now(),
		Updatedat: time.Now(),
	})

	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("%s", err))
		return
	}

	utils.RespondWithJson(w, http.StatusOK, models.DatabaseCategoryToCategory(category))

}
