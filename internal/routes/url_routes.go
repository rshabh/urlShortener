package routes

import (
	"URLSHORTENER/internal/handlers"

	"github.com/go-chi/chi/v5"
)

func RegisterRoutes() *chi.Mux {
	r := chi.NewRouter()
	r.Post("/generatedPath", handlers.SaveInMap)
	r.Get("/seeMap", handlers.GetMap)
	r.Get("/redirect/{s}", handlers.Redirect)

	return r

}
