package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth/v5"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/mickBoat00/TransactionAPI/handlers"
	"github.com/mickBoat00/TransactionAPI/sql/database"

	httpSwagger "github.com/swaggo/http-swagger/v2"
)

//	@title			Transaction API
//	@version		1.0
//	@description	This is a sample transactions server.
//	@termsOfService	http://swagger.io/terms/

//	@host		localhost:8000
//	@BasePath	/api/v1

//	@securityDefinitions.basic	BasicAuth

func main() {

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

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	router.Use(middleware.Logger)
	router.Use(jwtauth.Verifier(jwtauth.New("HS256", []byte(jwtSecretKey), nil)))

	v1Router := chi.NewRouter()

	v1Router.Get("/", handlers.Home)
	v1Router.Post("/users/", serverCfg.CreateUser)
	v1Router.Get("/docs/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8000/docs/swagger.yaml"), //The url pointing to API definition
	))

	workDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("%s", err)
	}

	filesDir := http.Dir(filepath.Join(workDir, "docs"))

	handlers.FileServer(router, "/docs", filesDir)

	router.Mount("/api/v1", v1Router)

	log.Printf("Server is started on PORT:%v", PORT)
	http.ListenAndServe(fmt.Sprintf(":%s", PORT), router)
}
