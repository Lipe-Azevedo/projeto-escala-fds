package mongobd

import (
	"context"
	"errors" // Importar para erro customizado
	"fmt"    // Importar para formatação de erro
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MONGODB_URL_ENV           = "MONGODB_URL"
	MONGODB_DATABASE_NAME_ENV = "MONGODB_DATABASE"
)

func NewMongoDBConnection(
	ctx context.Context,
) (*mongo.Database, error) {
	mongodb_uri := os.Getenv(MONGODB_URL_ENV) // A URI é lida desta variável de ambiente
	mongodb_database_name := os.Getenv(MONGODB_DATABASE_NAME_ENV)

	if mongodb_uri == "" {
		// Retorna um erro mais claro se a variável não estiver definida
		return nil, errors.New("MONGODB_URL environment variable not set or empty")
	}
	if mongodb_database_name == "" {
		// Retorna um erro mais claro se a variável não estiver definida
		return nil, errors.New("MONGODB_DATABASE environment variable not set or empty")
	}

	// A função options.Client().ApplyURI() é que analisa a URI.
	// Se mongodb_uri não começar com "mongodb://" ou "mongodb+srv://",
	// ela retornará o erro "scheme must be..." que é encapsulado por mongo.Connect.
	client, err := mongo.Connect(
		ctx,
		options.Client().ApplyURI(mongodb_uri))
	if err != nil {
		// O erro "error parsing uri: scheme must be..." é retornado aqui.
		// Podemos adicionar mais contexto ao erro, se desejado.
		return nil, fmt.Errorf("error connecting to MongoDB (uri: %s): %w", mongodb_uri, err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("error pinging MongoDB: %w", err)
	}

	return client.Database(mongodb_database_name), nil
}
