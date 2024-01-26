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

//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization

func main() {

	requiredEnvironmentVariables := [3]string{"PORT", "DATABASE_URL", "JWT_SECRET_KEY"}

	envs := getRequiredEnvironmentVariables(".env", requiredEnvironmentVariables)

	conn, err := sql.Open("postgres", envs["DATABASE_URL"])

	errorMessage(err)

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
	router.Use(jwtauth.Verifier(jwtauth.New("HS256", []byte(envs["JWT_SECRET_KEY"]), nil)))

	v1Router := chi.NewRouter()

	v1Router.Get("/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8000/docs/swagger.json"),
	))
	v1Router.Post("/signup/", serverCfg.CreateUser)
	v1Router.Post("/login/", serverCfg.LoginUser)

	v1Router.Get("/currencies/", serverCfg.ListCurrencies)

	v1Router.Get("/categories/", handlers.AuthMiddleware(serverCfg.ListCategories))
	v1Router.Post("/categories/", handlers.AuthMiddleware(serverCfg.CreateCategory))
	v1Router.Put("/categories/{id}/", handlers.AuthMiddleware(serverCfg.UpdateCategory))
	v1Router.Delete("/categories/{id}/", handlers.AuthMiddleware(serverCfg.DeleteCategory))

	v1Router.Get("/transactions/", handlers.AuthMiddleware(serverCfg.ListTransactions))
	v1Router.Post("/transactions/", handlers.AuthMiddleware(serverCfg.CreateTransaction))

	workDir, err := os.Getwd()
	errorMessage(err)

	filesDir := http.Dir(filepath.Join(workDir, "docs"))

	handlers.FileServer(router, "/docs", filesDir)

	router.Mount("/api/v1", v1Router)

	log.Printf("Server is started on PORT:%v", envs["PORT"])
	http.ListenAndServe(fmt.Sprintf(":%s", envs["PORT"]), router)
}

func getRequiredEnvironmentVariables(envs_file_name string, envs [3]string) map[string]string {

	environments := make(map[string]string)

	err := godotenv.Load(envs_file_name)

	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	for _, env := range envs {
		value := os.Getenv(env)
		if value == "" {
			log.Fatal("Cannot load environment variable ", env)
		}
		environments[env] = value

	}

	return environments

}

func errorMessage(err interface{}) {
	switch err.(type) {
	case string:
		log.Fatalf("Error: %s", err)
	}
}
