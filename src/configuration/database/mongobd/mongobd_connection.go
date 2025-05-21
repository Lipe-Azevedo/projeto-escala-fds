package mongobd

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MONGODB_URL_ENV           = "MONGODB_URL"
	MONGODB_DATABASE_NAME_ENV = "MONGODB_DATABASE" // Renomeado para clareza de MONGODB_USER_DB
)

func NewMongoDBConnection(
	ctx context.Context,
) (*mongo.Database, error) {
	mongodb_uri := os.Getenv(MONGODB_URL_ENV)
	mongodb_database_name := os.Getenv(MONGODB_DATABASE_NAME_ENV) // Usar variável renomeada

	client, err := mongo.Connect(
		ctx,
		options.Client().ApplyURI(mongodb_uri))
	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return client.Database(mongodb_database_name), nil
}
