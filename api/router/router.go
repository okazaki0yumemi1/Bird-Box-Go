package router

import (
	"github.com/go-chi/chi/v5"

	"bird-box-go/api/resource/bird"
	"bird-box-go/api/resource/health"
)

func New() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/health", health.Read)

	r.Route("/v1", func(r chi.Router) {
		birdAPI := &bird.API{}
		r.Get("/birds", birdAPI.List)
		r.Post("/birds", birdAPI.Create)
		r.Get("/birds/{id}", birdAPI.Read)
		r.Put("/birds/{id}", birdAPI.Update)
		r.Delete("/birds/{id}", birdAPI.Delete)
	})

	return r
}
