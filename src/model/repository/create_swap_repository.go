package repository

import (
	"context"
	"os"

	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/logger"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (sr *swapRepository) CreateSwap(
	swapDomain model.SwapDomainInterface,
) (model.SwapDomainInterface, *rest_err.RestErr) {
	logger.Info("Init createSwap repository",
		zap.String("journey", "createSwap"))

	collection_name := os.Getenv(MONGODB_SWAP_COLLECTION_ENV_KEY)
	collection := sr.databaseConnection.Collection(collection_name)

	value := converter.ConvertSwapDomainToEntity(swapDomain)

	result, err := collection.InsertOne(context.Background(), value)
	if err != nil {
		logger.Error("Error trying to create shift swap",
			err,
			zap.String("journey", "createSwap"))
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	value.ID = result.InsertedID.(primitive.ObjectID).Hex()

	logger.Info("CreateSwap repository executed successfully",
		zap.String("swapID", value.ID),
		zap.String("journey", "createSwap"))

	return converter.ConvertSwapEntityToDomain(*value), nil
}
