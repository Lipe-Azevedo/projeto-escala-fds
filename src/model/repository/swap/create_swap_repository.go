package swap

import (
	"context"
	"fmt"
	"os"

	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/logger"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model/repository/entity/converter" // Este caminho permanecerá global por enquanto
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func (sr *swapRepository) CreateSwap(
	swapDomain model.SwapDomainInterface,
) (model.SwapDomainInterface, *rest_err.RestErr) {
	logger.Info("Init CreateSwap repository",
		zap.String("journey", "createSwap"))

	collectionName := os.Getenv(MONGODB_SWAPS_COLLECTION_ENV_KEY) // Usando a constante do pacote
	if collectionName == "" {
		errorMessage := fmt.Sprintf("Environment variable %s not set for swaps collection name", MONGODB_SWAPS_COLLECTION_ENV_KEY)
		logger.Error(errorMessage, nil, zap.String("journey", "createSwap"))
		return nil, rest_err.NewInternalServerError("database configuration error: swaps collection name not set")
	}
	collection := sr.databaseConnection.Collection(collectionName)

	// Os conversores de entidade ainda estão em um local global.
	// Se decidirmos movê-los para subpastas de entidade (ex: entity/user, entity/swap),
	// este import mudaria. Por enquanto, está ok.
	value := converter.ConvertSwapDomainToEntity(swapDomain)

	result, err := collection.InsertOne(context.Background(), value)
	if err != nil {
		if writeException, ok := err.(mongo.WriteException); ok {
			for _, writeError := range writeException.WriteErrors {
				if writeError.Code == 11000 {
					errorMessage := fmt.Sprintf("Duplicate key error on creating swap: %s", writeError.Message)
					logger.Error(errorMessage, err,
						zap.String("journey", "createSwap"))
					return nil, rest_err.NewConflictError(errorMessage)
				}
			}
		}
		logger.Error("Error trying to create swap in repository", err,
			zap.String("journey", "createSwap"))
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	generatedID := result.InsertedID.(primitive.ObjectID)
	value.ID = generatedID.Hex() // A entidade SwapEntity tem ID string `bson:"_id,omitempty"`

	logger.Info("CreateSwap repository executed successfully",
		zap.String("swapID", value.ID),
		zap.String("journey", "createSwap"))

	return converter.ConvertSwapEntityToDomain(*value), nil
}
