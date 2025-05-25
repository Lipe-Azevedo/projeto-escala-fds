package workinfo

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

func (wr *workInfoRepository) FindWorkInfoByUserId(
	userId string,
) (model.WorkInfoDomainInterface, *rest_err.RestErr) {
	logger.Info(
		"Init FindWorkInfoByUserId repository",
		zap.String("journey", "findWorkInfoByUserId"),
		zap.String("userId_to_find", userId))

	collectionName := os.Getenv(MONGODB_WORKINFO_COLLECTION_ENV_KEY)
	if collectionName == "" {
		errorMessage := fmt.Sprintf("Environment variable %s not set for work_info collection name", MONGODB_WORKINFO_COLLECTION_ENV_KEY)
		logger.Error(errorMessage, nil, zap.String("journey", "findWorkInfoByUserId"))
		return nil, rest_err.NewInternalServerError("database configuration error: work_info collection name not set")
	}
	collection := wr.databaseConnection.Collection(collectionName)

	workInfoEntity := &entity.WorkInfoEntity{}
	// Filtro é pelo campo "_id", pois UserID na entidade WorkInfoEntity está mapeado para _id.
	filter := bson.M{"_id": userId}

	err := collection.FindOne(context.Background(), filter).Decode(workInfoEntity)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			logger.Warn("WorkInfo not found in repository for user ID (as _id)",
				zap.String("userId_as_id", userId),
				zap.String("journey", "findWorkInfoByUserId"))
			return nil, rest_err.NewNotFoundError(fmt.Sprintf("WorkInfo not found for user ID: %s", userId))
		}
		logger.Error("Error finding WorkInfo by user ID (as _id) in repository", err,
			zap.String("userId_as_id", userId),
			zap.String("journey", "findWorkInfoByUserId"))
		return nil, rest_err.NewInternalServerError(fmt.Sprintf("Error finding WorkInfo: %s", err.Error()))
	}

	logger.Info("FindWorkInfoByUserId repository executed successfully (found by _id)",
		zap.String("userID_found", workInfoEntity.UserID),
		zap.String("journey", "findWorkInfoByUserId"))

	return converter.ConvertWorkInfoEntityToDomain(*workInfoEntity), nil
}
