package services

import (
	"api-sw/src/core/domains/planet/repositories"
	"api-sw/src/core/domains/planet/services/create"
	"api-sw/src/core/domains/planet/services/getbyid"
	"api-sw/src/core/domains/planet/services/getbyname"
	"api-sw/src/core/domains/planet/services/list"
	"api-sw/src/shared/providers/logger"
	"context"
)

type Dependecies struct {
	Context    context.Context
	Repository repositories.IPlanetRepository
	Logger     logger.ILoggerProvider
}

type Services struct {
	Create    create.Service
	List      list.Service
	GetByID   getbyid.Service
	GetByName getbyname.Service
}

func NewPlanet(dep Dependecies) *Services {
	return &Services{
		Create: create.Service{
			Repository: dep.Repository,
			Logger:     dep.Logger,
		},
		List: list.Service{
			Repository: dep.Repository,
			Logger:     dep.Logger,
		},
		GetByID: getbyid.Service{
			Repository: dep.Repository,
			Logger:     dep.Logger,
		},
		GetByName: getbyname.Service{
			Repository: dep.Repository,
			Logger:     dep.Logger,
		},
	}
}
