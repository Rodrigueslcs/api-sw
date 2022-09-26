package planet

import (
	"api-sw/src/core/api/handlers/planet"
	"api-sw/src/shared/middlewares"

	"github.com/go-chi/chi"
)

func NewRoutes(router *chi.Mux, planetHandler planet.IHandler) {

	router.Route("/api/v1/planet", func(r chi.Router) {
		r.Post("/", middlewares.Handler(planetHandler.CreatePlanetHandler))
		r.Get("/", middlewares.Handler(planetHandler.ListPlanetHandler))
		r.Route("/{id}", func(r chi.Router) {
			r.Delete("/", middlewares.Handler(planetHandler.DeletePlanetHandler))
			r.Put("/", middlewares.Handler(planetHandler.UpdatePlanetHandler))
		})

	})

}
