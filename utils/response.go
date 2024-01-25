package utils

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/mickBoat00/TransactionAPI/models"
)

func IfErrorRespondWithErrorJson(w http.ResponseWriter, err interface{}, code int, errorMessage string) bool {

	errorFound := false

	if err != nil {
		if code > 499 {
			log.Println("Responding with error code ", code)
		}

		RespondWithJson(w, code, models.ErrorJsonParams{Error: errorMessage})
		errorFound = true

	}

	return errorFound

}

func RespondWithJson(w http.ResponseWriter, code int, payload interface{}) {

	data, err := json.Marshal(payload)

	if err != nil {
		log.Println("Failed to marshal JSON Response:", err)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}
