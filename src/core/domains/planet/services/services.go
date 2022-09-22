package services

import (
	"api-sw/src/core/domains/planet/repositories"
	"api-sw/src/core/domains/planet/services/create"
	"api-sw/src/shared/providers/logger"
	"context"
)

type Dependecies struct {
	Context    context.Context
	Repository repositories.IPlanetRepository
	Logger     logger.ILoggerProvider
}

type Services struct {
	Create create.Service
}

func NewPlanet(dep Dependecies) *Services {
	return &Services{
		Create: create.Service{
			Repository: dep.Repository,
			Logger:     dep.Logger,
		},
	}
}
