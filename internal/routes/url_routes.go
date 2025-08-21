package routes

import (
	"URLSHORTENER/internal/handlers"

	"github.com/go-chi/chi/v5"
)

func RegisterRoutes() *chi.Mux {
	r := chi.NewRouter()
	r.Post("/generatedPath", handlers.SaveInDb)
	r.Get("/", handlers.Redirect)

	return r

}
