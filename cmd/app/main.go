package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/5822791760/hr/internal/configs"
	"github.com/5822791760/hr/internal/db/postgres"
	"github.com/5822791760/hr/internal/routes"
	"github.com/5822791760/hr/pkg/helpers"
	"github.com/go-chi/chi/v5"
)

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

	helpers.InitCoreDB(db)

	r := chi.NewRouter()

	err = routes.InitRoutes(r, db)
	if err != nil {
		panic(err)
	}

	const Port = 3000

	fmt.Printf("\n======================================\n\n")
	fmt.Printf("Listening to port %d", Port)
	fmt.Printf("\n\n======================================\n\n")

	http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", Port), r)
}
