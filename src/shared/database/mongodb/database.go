package mongodb

import (
	"api-sw/server/config"
	"context"
	"fmt"
	"html"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	mongotrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/go.mongodb.org/mongo-driver/mongo"
)

type Storage struct {
	Error   error
	MongoDB *mongo.Database
}

var singletonStorage *Storage = nil

func New(ctx context.Context) Storage {
	cfg := *config.Instance

	if singletonStorage == nil {

		mongoDB, err := connect(ctx, cfg)

		singletonStorage = &Storage{
			Error:   err,
			MongoDB: mongoDB,
		}
	}
	return *singletonStorage

}

func connect(ctx context.Context, cfg config.Config) (*mongo.Database, error) {

	mongoURI := html.UnescapeString(fmt.Sprintf("mongodb://%s:%s@%s/%s", cfg.Mongo.User, cfg.Mongo.Pass, cfg.Mongo.Host, cfg.Mongo.Database))

	if cfg.Mongo.Args != "" {
		mongoURI = html.UnescapeString(fmt.Sprintf("mongodb://%s:%s@%s/%s?%s", cfg.Mongo.User, cfg.Mongo.Pass, cfg.Mongo.Host, cfg.Mongo.Database, cfg.Mongo.Args))
	}

	opts := options.Client()
	opts.Monitor = mongotrace.NewMonitor()
	opts.ApplyURI(mongoURI)

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, err
	}

	fmt.Println(cfg.Mongo.Database)

	return client.Database(cfg.Mongo.Database), nil

}
