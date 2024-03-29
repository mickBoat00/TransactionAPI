package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/mickBoat00/TransactionAPI/models"
	"github.com/mickBoat00/TransactionAPI/sql/database"
	"github.com/mickBoat00/TransactionAPI/utils"
)

func decodeCategoryRequestBody(w http.ResponseWriter, r *http.Request, model models.CategoryRequestParams) models.CategoryRequestParams {
	decoder := json.NewDecoder(r.Body)

	params := model

	err := decoder.Decode(&params)

	utils.IfErrorRespondWithErrorJson(w, err, http.StatusInternalServerError, fmt.Sprintf("%s", err))

	return params
}

func (serverCfg *ServerConfig) checkCategoriesPermissions(w http.ResponseWriter, r *http.Request, user_id uuid.UUID, category_uuid uuid.UUID) bool {

	userCategoriesIds, err := serverCfg.DB.GetUserCategoryIds(r.Context(), user_id)

	if utils.IfErrorRespondWithErrorJson(w, err, http.StatusInternalServerError, fmt.Sprintf("%s", err)) {
		return false
	}

	if !slices.Contains(userCategoriesIds, category_uuid) {
		if utils.IfErrorRespondWithErrorJson(w, "Permission Denied", http.StatusForbidden, "Permission Denied") {
			return false
		}
	}

	return true
}

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

	utils.IfErrorRespondWithErrorJson(w, err, http.StatusInternalServerError, fmt.Sprintf("%s", err))

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

	params := decodeCategoryRequestBody(w, r, models.CategoryRequestParams{})

	category, err := serverCfg.DB.CreateUserCategory(r.Context(), database.CreateUserCategoryParams{
		ID:        uuid.New(),
		Name:      params.Name,
		UserID:    user_id,
		Createdat: time.Now(),
		Updatedat: time.Now(),
	})

	utils.IfErrorRespondWithErrorJson(w, err, http.StatusInternalServerError, fmt.Sprintf("%s", err))

	utils.RespondWithJson(w, http.StatusOK, models.DatabaseCategoryToCategory(category))

}

// UpdateCategory godoc
//
//	@Summary		Update a category
//	@Description	Update by category UUID
//	@Tags			Category
//	@Accept			json
//	@Produce		json
//	@Param			uuid	path 	string	true	"Category UUID"
//	@Param			Category	body		models.CategoryRequestParams	true	"Add Category"
//	@Success		204	{object}	models.Category
//	@Failure		400	{object}	models.ErrorJsonParams
//	@Failure		500	{object}	models.ErrorJsonParams
//	@Security		ApiKeyAuth
//	@Router			/categories/{uuid}/ [put]
func (serverCfg *ServerConfig) UpdateCategory(w http.ResponseWriter, r *http.Request, user_id uuid.UUID) {

	category_id := chi.URLParam(r, "id")

	category_uuid, err := uuid.Parse(category_id)

	if utils.IfErrorRespondWithErrorJson(w, err, http.StatusInternalServerError, fmt.Sprintf("%s", err)) {
		return
	}

	if !serverCfg.checkCategoriesPermissions(w, r, user_id, category_uuid) {
		return
	}

	params := decodeCategoryRequestBody(w, r, models.CategoryRequestParams{})

	if params.Name == "" {
		utils.IfErrorRespondWithErrorJson(w, err, http.StatusInternalServerError, fmt.Sprintf("%s", err))
		return
	}

	category, err := serverCfg.DB.UpdateUserCategory(r.Context(), database.UpdateUserCategoryParams{
		Name:      params.Name,
		Updatedat: time.Now(),
		UserID:    user_id,
	})

	if err != nil {
		utils.IfErrorRespondWithErrorJson(w, err, http.StatusInternalServerError, fmt.Sprintf("%s", err))
		return
	}

	utils.RespondWithJson(w, http.StatusOK, models.DatabaseCategoryToCategory(category))

}

// DeleteCategory godoc
//
//	@Summary		Delete a category
//	@Description	Delete by category UUID
//	@Tags			Category
//	@Accept			json
//	@Produce		json
//	@Param			uuid	path 	string	true	"Category UUID"
//	@Success		204	{object}	models.Category
//	@Failure		400	{object}	models.ErrorJsonParams
//	@Failure		500	{object}	models.ErrorJsonParams
//	@Security		ApiKeyAuth
//	@Router			/categories/{uuid}/ [delete]
func (serverCfg *ServerConfig) DeleteCategory(w http.ResponseWriter, r *http.Request, user_id uuid.UUID) {

	category_id := chi.URLParam(r, "id")

	category_uuid, err := uuid.Parse(category_id)

	utils.IfErrorRespondWithErrorJson(w, err, http.StatusInternalServerError, fmt.Sprintf("%s", err))

	if !serverCfg.checkCategoriesPermissions(w, r, user_id, category_uuid) {
		return
	}

	db_err := serverCfg.DB.DeleteUserCategories(r.Context(), database.DeleteUserCategoriesParams{
		ID:     category_uuid,
		UserID: user_id,
	})

	utils.IfErrorRespondWithErrorJson(w, db_err, http.StatusBadRequest, fmt.Sprintf("%s", err))

	utils.RespondWithJson(w, http.StatusNoContent, struct{}{})

}
