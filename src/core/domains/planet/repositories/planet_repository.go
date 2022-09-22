package repositories

import "api-sw/src/core/domains/planet/entities"

type IPlanetRepository interface {
	Create(document entities.PlanetCreate) (entities.Planet, error)
	FindByID(id string) (entities.Planet, error)
	FindByName(name string) (entities.Planet, error)
}
