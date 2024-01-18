package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/joho/godotenv"
	"github.com/mickBoat00/TransactionAPI/handlers"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	PORT := os.Getenv("PORT")

	if PORT == "" {
		log.Fatal("Environment variable PORT is missing.")
	}

	router.Get("/", handlers.Home)
	router.Post("/users", handlers.CreateUser)

	log.Printf("Server is started on PORT:%v", PORT)
	http.ListenAndServe(fmt.Sprintf(":%s", PORT), router)
}
