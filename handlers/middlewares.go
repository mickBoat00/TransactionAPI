package handlers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/jwtauth/v5"
	"github.com/google/uuid"
	"github.com/mickBoat00/TransactionAPI/utils"
)

func AuthMiddleware(handler func(w http.ResponseWriter, r *http.Request, user_id uuid.UUID)) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		_, claims, err := jwtauth.FromContext(r.Context())

		utils.IfErrorRespondWithErrorJson(w, err, http.StatusForbidden, fmt.Sprintf("%s", err))

		user_id, err := uuid.Parse(claims["id"].(string))

		utils.IfErrorRespondWithErrorJson(w, err, http.StatusInternalServerError, fmt.Sprintf("%s", err))

		handler(w, r, user_id)

	})

}
