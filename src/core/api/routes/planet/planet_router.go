package planet

import (
	"api-sw/src/core/api/handlers/planet"
	"api-sw/src/shared/middlewares"

	"github.com/go-chi/chi"
)

func NewRoutes(router *chi.Mux, planetHandler planet.IHandler) {

	router.Route("/api/v1/planet", func(r chi.Router) {
		r.Post("/", middlewares.Handler(planetHandler.CreateHandler))
	})

}
