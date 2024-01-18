package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/mickBoat00/TransactionAPI/models"
	"github.com/mickBoat00/TransactionAPI/utils"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	params := models.UserRequestParams{}

	err := decoder.Decode(&params)

	if err != nil {
		return
	}

	user := models.UserRequestParams{
		Email:    params.Email,
		Password: params.Password,
	}

	utils.RespondWithJson(w, 200, user)
}
