package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/logger"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func (wr *workInfoRepository) CreateWorkInfo(
	workInfoDomain model.WorkInfoDomainInterface,
) (model.WorkInfoDomainInterface, *rest_err.RestErr) {
	logger.Info(
		"Init createWorkInfo repository.",
		zap.String("journey", "createWorkInfo"))

	collectionNameKey := MONGODB_WORKINFO_COLLECTION_ENV_KEY
	collectionName := os.Getenv(collectionNameKey)

	if collectionName == "" {
		errorMessage := fmt.Sprintf("Environment variable %s not set for work_info collection name", collectionNameKey)
		logger.Error(errorMessage, nil, zap.String("journey", "createWorkInfo"))
		return nil, rest_err.NewInternalServerError("database configuration error: work_info collection name not set")
	}
	collection := wr.dataBaseConnection.Collection(collectionName)

	value := converter.ConvertWorkInfoDomainToEntity(workInfoDomain)
	// 'value' agora tem UserID (com tag bson:"_id") preenchido com workInfoDomain.GetUserId()

	_, err := collection.InsertOne(context.Background(), value)
	if err != nil {
		if writeException, ok := err.(mongo.WriteException); ok {
			for _, writeError := range writeException.WriteErrors {
				if writeError.Code == 11000 { // Erro de chave duplicada
					errorMessage := fmt.Sprintf("WorkInfo for user ID %s already exists (duplicate _id)", value.UserID)
					logger.Error(errorMessage, err,
						zap.String("journey", "createWorkInfo"),
						zap.String("userID", value.UserID))
					return nil, rest_err.NewConflictError(errorMessage)
				}
			}
		}
		logger.Error("Error creating work info in repository", err,
			zap.String("journey", "createWorkInfo"),
			zap.String("userID", value.UserID))
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	logger.Info("CreateWorkInfo repository executed successfully",
		zap.String("userID", workInfoDomain.GetUserId()), // O UserID do domain é o _id
		zap.String("journey", "createWorkInfo"))

	return workInfoDomain, nil
}
