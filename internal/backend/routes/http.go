package routes

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/5822791760/hr/internal/backend/handlers/httpv1"
	"github.com/5822791760/hr/internal/backend/repos/authorrepo"
	"github.com/5822791760/hr/internal/backend/usecases/authorusecase"
	"github.com/5822791760/hr/pkg/coreutil"
	"github.com/go-chi/chi/v5"
)

func InitRoutes(r *chi.Mux, db *sql.DB) error {
	clock := coreutil.NewClock()

	// Author Repo
	authorReadRepo := authorrepo.NewReadRepo()
	authorWriteRepo := authorrepo.NewWriteRepo(authorReadRepo, clock)

	// Use Case
	authorUsecase := authorusecase.NewAuthorUseCase(authorReadRepo, authorWriteRepo)

	// Handlers
	authorHandler := httpv1.NewAuthorHandler(db, authorUsecase)

	r.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Get("/authors", authorHandler.GetAll)
			r.Get("/authors/{id}", authorHandler.GetOne)
			r.Post("/authors", authorHandler.Create)
			r.Put("/authors/{id}", authorHandler.Update)
			r.Delete("/authors/{id}", authorHandler.Delete)
		})
	})

	if err := PrintRoutes(r); err != nil {
		return err
	}

	return nil
}

func PrintRoutes(r chi.Router) error {
	fmt.Println()

	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		fmt.Printf("%s %s\n", method, route)
		return nil
	}

	if err := chi.Walk(r, walkFunc); err != nil {
		return err
	}

	return nil
}
