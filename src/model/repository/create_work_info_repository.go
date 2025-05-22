package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/logger"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/mongo" // Adicionar import para mongo.WriteException
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
	// Agora, value.UserID (que tem a tag `bson:"_id"`) contém o userId do domain.
	// Ao fazer InsertOne, o MongoDB usará o valor de value.UserID como o _id do documento.

	_, err := collection.InsertOne(context.Background(), value)
	if err != nil {
		// Verifica se o erro é uma exceção de escrita do MongoDB
		if writeException, ok := err.(mongo.WriteException); ok {
			for _, writeError := range writeException.WriteErrors {
				// O código 11000 indica um erro de chave duplicada
				if writeError.Code == 11000 {
					errorMessage := fmt.Sprintf("WorkInfo for user ID %s already exists", value.UserID)
					logger.Error(errorMessage, err,
						zap.String("journey", "createWorkInfo"),
						zap.String("userID", value.UserID))
					return nil, rest_err.NewConflictError(errorMessage)
				}
			}
		}
		// Para outros erros de InsertOne
		logger.Error("Error creating work info in repository", err,
			zap.String("journey", "createWorkInfo"),
			zap.String("userID", value.UserID)) // Adicionado UserID ao log de erro genérico
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	logger.Info("CreateWorkInfo repository executed successfully",
		zap.String("userID", workInfoDomain.GetUserId()),
		zap.String("journey", "createWorkInfo"))

	return workInfoDomain, nil
}
