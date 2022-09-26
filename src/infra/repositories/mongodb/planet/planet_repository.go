package planet

import (
	"api-sw/src/core/domains/planet/entities"
	"api-sw/src/shared/database/mongodb"
	"api-sw/src/shared/providers/logger"
	"api-sw/src/shared/tools/namespace"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	collection = "planet"
	Namespace  = namespace.New("infra.repositories.mongodb.planet.planet_repository")
)

type Repository struct {
	Context    context.Context
	Collection *mongo.Collection
	Logger     logger.ILoggerProvider
}

func Setup(ctx context.Context) *Repository {
	connection := mongodb.New(ctx)
	log := logger.Instance

	return &Repository{
		Context:    ctx,
		Collection: connection.MongoDB.Collection(collection),
		Logger:     log,
	}
}

func (r Repository) FindByID(id string) (entities.Planet, error) {
	r.Logger.Info(Namespace.Concat("FindByID"), "")

	var document entities.Planet
	filter := bson.M{"_id": id}

	err := r.Collection.FindOne(r.Context, filter).Decode(&document)

	if err == mongo.ErrNoDocuments {
		return entities.Planet{}, nil
	}

	if err != nil {
		return entities.Planet{}, err
	}

	return document, nil
}

func (r Repository) FindByName(name string) (entities.Planet, error) {
	r.Logger.Info(Namespace.Concat("FindByName"), "")

	var document entities.Planet

	filter := bson.M{"name": bson.M{"$regex": name, "$options": "i"}}

	err := r.Collection.FindOne(r.Context, filter).Decode(&document)

	if err == mongo.ErrNoDocuments {

		return entities.Planet{}, nil
	}
	if err != nil {
		return entities.Planet{}, err
	}
	return document, nil
}

func (r Repository) FindAll(filter bson.M) (entities.Planets, error) {
	r.Logger.Info(Namespace.Concat("findAll"), "")

	var planets entities.Planets

	cursor, err := r.Collection.Find(r.Context, filter)
	if err != nil {
		return nil, err
	}
	for cursor.Next(r.Context) {
		document := entities.Planet{}
		cursor.Decode(&document)
		planets = append(planets, document)
	}

	return planets, nil
}

func (r Repository) Create(document entities.PlanetCreate) (entities.Planet, error) {
	r.Logger.Info(Namespace.Concat("Create"), "")

	_, err := r.Collection.InsertOne(r.Context, document)
	if err != nil {
		return entities.Planet{}, err
	}

	return r.FindByID(document.ID)
}
