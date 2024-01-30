package mongodb

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	DB_URL = os.Getenv("DB_URL")
	DB_NAME = os.Getenv("DB_NAME")
)

func NewMongoDBConnection(ctx context.Context) (*mongo.Database, error) {
	client, err := mongo.Connect(
		ctx, options.Client().ApplyURI(DB_URL),
	)
	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	database := client.Database(DB_NAME)

	return database, nil
}
