package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

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

	err = routes.InitRoutes(r)
	if err != nil {
		panic(err)
	}

	ListenAndServe(r, 3000)
}

func ListenAndServe(r *chi.Mux, port int) {
	con := make(chan struct{})
	srv := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%d", port),
		Handler: r,
	}

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGTERM)
		<-sigint

		if err := srv.Shutdown(context.Background()); err != nil {
			log.Fatalf("HTTP server Shutdown: %v", err)
		}

		fmt.Printf("\n\n")
		fmt.Printf("Gracefully Shutting down Server....")
		fmt.Printf("\n\n")

		close(con)
	}()

	fmt.Printf("\n======================================\n\n")
	fmt.Printf("Listening to port %d", port)
	fmt.Printf("\n\n======================================\n\n")
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		// Error starting or closing listener:
		fmt.Printf("HTTP server Shutdown: %v", err)
	}

	<-con
}
