package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/jwtauth/v5"
	"github.com/google/uuid"
	"github.com/mickBoat00/TransactionAPI/models"
	"github.com/mickBoat00/TransactionAPI/sql/database"
	"github.com/mickBoat00/TransactionAPI/utils"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	bytePassword := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	return string(hash), err
}

func checkPassword(hashPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	return err == nil
}

func createJwtToken(user_claim map[string]interface{}) (string, error) {
	tokenAuth := jwtauth.New("HS256", []byte(os.Getenv("JWT_SECRET_KEY")), nil)
	_, tokenString, err := tokenAuth.Encode(user_claim)
	return tokenString, err
}

func decodeUserRequestBody(w http.ResponseWriter, r *http.Request, model models.UserRequestParams) models.UserRequestParams {
	decoder := json.NewDecoder(r.Body)

	params := model

	err := decoder.Decode(&params)

	utils.IfErrorRespondWithErrorJson(w, err, http.StatusInternalServerError, fmt.Sprintf("%s", err))

	return params
}

// CreateUser godoc
//
//	@Summary		Sign up
//	@Description	create a new user
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			User	body		models.UserRequestParams	true	"Add User"
//	@Success		200		{object}	models.UserResponseParams
//	@Failure		400	{object}	models.ErrorJsonParams
//	@Failure		500	{object}	models.ErrorJsonParams
//	@Router			/signup/ [post]
func (serverCfg *ServerConfig) CreateUser(w http.ResponseWriter, r *http.Request) {

	params := decodeUserRequestBody(w, r, models.UserRequestParams{})

	passwordHashed, err := hashPassword(params.Password)

	if utils.IfErrorRespondWithErrorJson(w, err, http.StatusInternalServerError, fmt.Sprintf("%s", err)) {
		return
	}

	user, err := serverCfg.DB.CreateUser(
		r.Context(),
		database.CreateUserParams{
			ID:        uuid.New(),
			Email:     params.Email,
			Password:  passwordHashed,
			Createdat: time.Now(),
			Updatedat: time.Now(),
		},
	)

	if utils.IfErrorRespondWithErrorJson(w, err, http.StatusInternalServerError, fmt.Sprintf("%s", err)) {
		return
	}

	token, err := createJwtToken(map[string]interface{}{"id": user.ID, "email": user.Email})

	if utils.IfErrorRespondWithErrorJson(w, err, http.StatusInternalServerError, fmt.Sprintf("%s", err)) {
		return
	}

	utils.RespondWithJson(w, http.StatusOK, models.UserResponseParams{
		Token: token,
	})
}

// LoginUser godoc
//
//	@Summary		Generate jwt token
//	@Description	Generate jwt token for user
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			User	body		models.UserRequestParams	true	"Login User"
//	@Success		200		{object}	models.UserResponseParams
//	@Failure		400	{object}	models.ErrorJsonParams
//	@Failure		500	{object}	models.ErrorJsonParams
//	@Router			/login/ [post]
func (serverCfg *ServerConfig) LoginUser(w http.ResponseWriter, r *http.Request) {

	params := decodeUserRequestBody(w, r, models.UserRequestParams{})

	user, err := serverCfg.DB.GetUserViaEmail(r.Context(), params.Email)

	if utils.IfErrorRespondWithErrorJson(w, err, http.StatusBadRequest, "Invalid credentails.") {
		return
	}

	validPassword := checkPassword(user.Password, params.Password)

	if !validPassword {
		if utils.IfErrorRespondWithErrorJson(w, "Invalid credentails.", http.StatusBadRequest, "Invalid credentails. password") {
			return
		}
	}

	log.Print("end?   ")

	token, err := createJwtToken(map[string]interface{}{"id": user.ID, "email": user.Email})

	if utils.IfErrorRespondWithErrorJson(w, err, http.StatusInternalServerError, fmt.Sprintf("%s", err)) {
		return
	}

	utils.RespondWithJson(w, http.StatusOK, models.UserResponseParams{
		Token: token,
	})

}
