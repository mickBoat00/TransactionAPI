package handlers

import (
	"encoding/json"
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

func createJwtToken(user_claim map[string]interface{}) (string, error) {
	tokenAuth := jwtauth.New("HS256", []byte(os.Getenv("JWT_SECRET_KEY")), nil)
	_, tokenString, err := tokenAuth.Encode(user_claim)
	return tokenString, err
}

// CreateUser godoc
//
//	@Summary		Create a user
//	@Description	create by json User
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			User	body		models.UserRequestParams	true	"Add User"
//	@Success		200		{object}	models.UserResponseParams
//	@Router			/users/ [post]
func (serverCfg *ServerConfig) CreateUser(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	params := models.UserRequestParams{}

	err := decoder.Decode(&params)

	if err != nil {
		log.Print("Server Error", err)
		return
	}

	passwordHashed, err := hashPassword(params.Password)

	if err != nil {
		log.Print("Server Error", err)
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

	if err != nil {
		log.Print("Server Error", err)
		return
	}

	token, err := createJwtToken(map[string]interface{}{"id": user.ID, "email": user.Email})

	if err != nil {
		log.Print("Server Error", err)
		return
	}

	utils.RespondWithJson(w, 200, models.UserResponseParams{
		Token: token,
	})
}
