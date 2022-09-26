package entities

import (
	"api-sw/src/shared/tools/apisw"

	"github.com/google/uuid"
)

type Planet struct {
	ID          string `json:"id" bson:"_id"`
	Name        string `json:"name" bson:"name"`
	Climate     string `json:"climate" bson:"climate"`
	Terrain     string `json:"terrain" bson:"terrain"`
	Apparitions int    `json:"apparitions" bson:"Apparitions"`
}

type PlanetCreate struct {
	ID          string `json:"id" bson:"_id"`
	Name        string `json:"name" bson:"name"`
	Climate     string `json:"climate" bson:"climate"`
	Terrain     string `json:"terrain" bson:"terrain"`
	Apparitions int    `json:"apparitions" bson:"Apparitions"`
}

type PlanetUpdate struct {
	Name    string `json:"name" bson:"name"`
	Climate string `json:"climate" bson:"climate"`
	Terrain string `json:"terrain" bson:"terrain"`
}

type Planets []Planet

func (m *PlanetCreate) Populate() {
	if qtde, err := apisw.GetQtdeFilm(m.Name); err == nil {
		m.Apparitions = qtde
	}

	m.ID = uuid.New().String()
}
