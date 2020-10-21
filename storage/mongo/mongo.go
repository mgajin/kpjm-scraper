package mongo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Storage struct {
	*mongo.Client
	*mongo.Database
}

type Config struct {
	ConnectionURL string
	DatabaseName  string
}

func NewStorage(config *Config) (*Storage, error) {

	clientOptions := options.Client().ApplyURI(config.ConnectionURL)
	clientOptions.SetMaxPoolSize(200)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		return nil, err
	}

	database := client.Database(config.DatabaseName)

	return &Storage{
		Client:   client,
		Database: database,
	}, nil
}

func (s *Storage) Collection(name string) *mongo.Collection {
	return s.Database.Collection(name)
}
