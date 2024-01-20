package utils

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/mickBoat00/TransactionAPI/models"
)

func RespondWithError(w http.ResponseWriter, code int, errorMessage string) {

	if code > 499 {
		log.Println("Responding with error code 5XX.")
	}

	RespondWithJson(w, code, models.ErrorJsonParams{Error: errorMessage})

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
