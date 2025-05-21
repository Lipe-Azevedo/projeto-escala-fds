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

func (sr *shiftSwapRepository) FindShiftSwapByID(
	id string,
) (model.ShiftSwapDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findShiftSwapByID repository",
		zap.String("journey", "findShiftSwapByID"))

	collection_name := os.Getenv(MONGODB_SHIFTSWAP_COLLECTION_ENV_KEY)
	collection := sr.databaseConnection.Collection(collection_name)

	shiftSwapEntity := &entity.ShiftSwapEntity{}

	filter := bson.D{{Key: "_id", Value: id}}
	err := collection.FindOne(context.Background(), filter).Decode(shiftSwapEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf("ShiftSwap not found with ID: %s", id)
			logger.Error(errorMessage,
				err,
				zap.String("journey", "findShiftSwapByID"))
			return nil, rest_err.NewNotFoundError(errorMessage)
		}

		logger.Error("Error trying to find shift swap",
			err,
			zap.String("journey", "findShiftSwapByID"))
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	logger.Info("FindShiftSwapByID repository executed successfully",
		zap.String("shiftSwapID", id),
		zap.String("journey", "findShiftSwapByID"))

	return converter.ConvertShiftSwapEntityToDomain(*shiftSwapEntity), nil
}
