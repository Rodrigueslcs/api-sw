package routes

import (
	"api-sw/src/core/api/handlers"
	"api-sw/src/core/api/routes/planet"
	"api-sw/src/shared/middlewares"

	"github.com/go-chi/chi"
)

type router struct {
	Client   *chi.Mux
	Handlers handlers.IHandler
}

func NewRoutes(handlers handlers.IHandler) *router {
	return &router{
		Client:   chi.NewRouter(),
		Handlers: handlers,
	}
}

func (r *router) Setup() {
	middlewares.Default(r.Client)
	planet.NewRoutes(r.Client, r.Handlers.NewPlanetHandler())
}
