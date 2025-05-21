package repository

import (
	"context"
	"os"

	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/logger"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson"
	"go.uber.org/zap"
)

func (sr *shiftSwapRepository) UpdateShiftSwap(
	id string,
	shiftSwapDomain model.ShiftSwapDomainInterface,
) *rest_err.RestErr {
	logger.Info("Init updateShiftSwap repository",
		zap.String("journey", "updateShiftSwap"))

	collection_name := os.Getenv(MONGODB_SHIFTSWAP_COLLECTION_ENV_KEY)
	collection := sr.databaseConnection.Collection(collection_name)

	value := converter.ConvertShiftSwapDomainToEntity(shiftSwapDomain)
	filter := bson.D{{Key: "_id", Value: id}}
	update := bson.D{{Key: "$set", Value: value}}

	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		logger.Error("Error trying to update shift swap",
			err,
			zap.String("journey", "updateShiftSwap"))
		return rest_err.NewInternalServerError(err.Error())
	}

	logger.Info("UpdateShiftSwap repository executed successfully",
		zap.String("shiftSwapID", id),
		zap.String("journey", "updateShiftSwap"))

	return nil
}
