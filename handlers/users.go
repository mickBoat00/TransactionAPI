package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/mickBoat00/TransactionAPI/models"
	"github.com/mickBoat00/TransactionAPI/utils"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	bytePassword := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	return string(hash), err
}

func CreateUser(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	params := models.UserRequestParams{}

	err := decoder.Decode(&params)

	if err != nil {
		return
	}

	passwordHashed, err := hashPassword(params.Password)

	if err != nil {
		return
	}

	user := models.UserRequestParams{
		Email:    params.Email,
		Password: passwordHashed,
	}

	utils.RespondWithJson(w, 200, user)
}
