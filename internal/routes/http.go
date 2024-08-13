package routes

import (
	"fmt"
	"net/http"

	"github.com/5822791760/hr/internal/handlers/https"
	"github.com/5822791760/hr/internal/repos"
	"github.com/5822791760/hr/internal/usecases"
	"github.com/go-chi/chi/v5"
)

func InitRoutes(r *chi.Mux) error {
	// Repos
	authorRepo := repos.NewAuthorRepo()

	// Use Case
	authorUsecase := usecases.NewAuthorUseCase(authorRepo)

	// Handlers
	authorHandler := https.NewAuthorHandler(authorUsecase)

	r.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Get("/authors", authorHandler.QueryAuthors)
			r.Get("/authors/{id}", authorHandler.FindOne)
			r.Post("/authors", authorHandler.CreateAuthor)
			r.Put("/authors/{id}", authorHandler.UpdateAuthor)
			r.Delete("/authors/{id}", authorHandler.DeleteAuthor)
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
