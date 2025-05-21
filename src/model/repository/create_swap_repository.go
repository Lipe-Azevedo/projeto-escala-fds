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

func (sr *shiftSwapRepository) CreateShiftSwap(
	shiftSwapDomain model.ShiftSwapDomainInterface,
) (model.ShiftSwapDomainInterface, *rest_err.RestErr) {
	logger.Info("Init createShiftSwap repository",
		zap.String("journey", "createShiftSwap"))

	collection_name := os.Getenv(MONGODB_SHIFTSWAP_COLLECTION_ENV_KEY)
	collection := sr.databaseConnection.Collection(collection_name)

	value := converter.ConvertShiftSwapDomainToEntity(shiftSwapDomain)

	result, err := collection.InsertOne(context.Background(), value)
	if err != nil {
		logger.Error("Error trying to create shift swap",
			err,
			zap.String("journey", "createShiftSwap"))
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	value.ID = result.InsertedID.(primitive.ObjectID).Hex()

	logger.Info("CreateShiftSwap repository executed successfully",
		zap.String("shiftSwapID", value.ID),
		zap.String("journey", "createShiftSwap"))

	return converter.ConvertShiftSwapEntityToDomain(*value), nil
}
