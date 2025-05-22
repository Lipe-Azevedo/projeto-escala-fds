package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/logger"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model/repository/entity"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func (sr *swapRepository) FindSwapByID(
	id string,
) (model.SwapDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findSwapByID repository",
		zap.String("journey", "findSwapByID"))

	collection_name := os.Getenv(MONGODB_SWAP_COLLECTION_ENV_KEY)
	collection := sr.databaseConnection.Collection(collection_name)

	SwapEntity := &entity.SwapEntity{}

	filter := bson.D{{Key: "_id", Value: id}}
	err := collection.FindOne(context.Background(), filter).Decode(SwapEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf("Swap not found with ID: %s", id)
			logger.Error(errorMessage,
				err,
				zap.String("journey", "findSwapByID"))
			return nil, rest_err.NewNotFoundError(errorMessage)
		}

		logger.Error("Error trying to find shift swap",
			err,
			zap.String("journey", "findSwapByID"))
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	logger.Info("FindSwapByID repository executed successfully",
		zap.String("SwapID", id),
		zap.String("journey", "findSwapByID"))

	return converter.ConvertSwapEntityToDomain(*SwapEntity), nil
}
