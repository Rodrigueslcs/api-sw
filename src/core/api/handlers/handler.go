package handlers

import (
	"api-sw/src/core/api/handlers/planet"
	"api-sw/src/shared/providers/logger"
)

type Dependencies struct {
	Logger logger.ILoggerProvider
}

type IHandler interface {
	NewPlanetHandler() planet.IHandler
}

type handler struct {
	PlanetHandler planet.IHandler
}

func NewHandler(dep Dependencies) handler {
	return handler{
		PlanetHandler: planet.NewHandler(dep.Logger),
	}
}

func (h handler) NewPlanetHandler() planet.IHandler {
	return h.PlanetHandler
}
