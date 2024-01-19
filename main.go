package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth/v5"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/mickBoat00/TransactionAPI/handlers"
	"github.com/mickBoat00/TransactionAPI/sql/database"
)

func main() {
	router := chi.NewRouter()

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	PORT := os.Getenv("PORT")

	if PORT == "" {
		log.Fatal("Environment variable PORT is missing.")
	}

	databaseUrl := os.Getenv("DATABASE_URL")
	if databaseUrl == "" {
		log.Fatal("Cannot load database connection")
	}

	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")
	if jwtSecretKey == "" {
		log.Fatal("Cannot load jwt secret")
	}

	conn, err := sql.Open("postgres", databaseUrl)

	if err != nil {
		log.Fatalf("Cannot connect to database %s", err)
	}

	serverCfg := &handlers.ServerConfig{DB: database.New(conn)}

	router.Use(middleware.Logger)
	router.Use(jwtauth.Verifier(jwtauth.New("HS256", []byte(jwtSecretKey), nil)))

	router.Get("/", handlers.Home)
	router.Post("/users", serverCfg.CreateUser)

	log.Printf("Server is started on PORT:%v", PORT)
	http.ListenAndServe(fmt.Sprintf(":%s", PORT), router)
}
