package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/5822791760/hr/internal/configs"
	"github.com/5822791760/hr/internal/db/postgres"
	"github.com/go-chi/chi/v5"

	_ "github.com/5822791760/hr/docs"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

//	@title			Swagger Example API
//	@version		1.0
//	@description	This is a sample server Petstore server.
//	@termsOfService	http://swagger.io/terms/

//	@host		localhost:3000
//	@BasePath	/api/v1

// @securityDefinitions.apikey	Bearer
// @in							header
// @name						Authorization
func main() {
	err := configs.LoadConfig()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	db, err := postgres.ConnectDB(ctx, configs.GetDBConnectionString())
	if err != nil {
		panic(err)
	}
	defer db.Close()

	r := chi.NewRouter()

	r.Get("/api/documentation/*", httpSwagger.Handler(
		httpSwagger.URL("/api/documentation/doc.json"),
		httpSwagger.UIConfig(map[string]string{
			"persistAuthorization": "true",
		}),
	))

	err = configs.InitRoutes(r, db)
	if err != nil {
		panic(err)
	}

	const Port = 3000

	fmt.Printf("\n======================================\n\n")
	fmt.Printf("Listening to port %d", Port)
	fmt.Printf("\n\n======================================\n\n")

	http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", Port), r)
}
