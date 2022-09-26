package repositories

import (
	"api-sw/src/core/domains/planet/entities"

	"go.mongodb.org/mongo-driver/bson"
)

type IPlanetRepository interface {
	Create(document entities.PlanetCreate) (entities.Planet, error)
	FindByID(id string) (entities.Planet, error)
	FindByName(name string) (entities.Planet, error)
	FindAll(filter bson.M) (entities.Planets, error)
	Update(id string, document entities.PlanetUpdate) (entities.Planet, error)
	Delete(id string) error
}
