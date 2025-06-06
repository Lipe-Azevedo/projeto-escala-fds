package workinfo

import (
	"context"
	"fmt"
	"os"

	"github.com/Lipe-Azevedo/escala-fds/src/configuration/logger"
	"github.com/Lipe-Azevedo/escala-fds/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/escala-fds/src/model/domain"
	workinfoconv "github.com/Lipe-Azevedo/escala-fds/src/model/repository/entity/converter/workinfo" // IMPORT MODIFICADO
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func (wr *workInfoRepository) CreateWorkInfo(
	workInfoDomain domain.WorkInfoDomainInterface,
) (domain.WorkInfoDomainInterface, *rest_err.RestErr) {
	logger.Info(
		"Init CreateWorkInfo repository",
		zap.String("journey", "createWorkInfo"),
		zap.String("userId", workInfoDomain.GetUserId()))

	collectionName := os.Getenv(MONGODB_WORKINFO_COLLECTION_ENV_KEY)
	if collectionName == "" {
		errorMessage := fmt.Sprintf("Environment variable %s not set for work_info collection name", MONGODB_WORKINFO_COLLECTION_ENV_KEY)
		logger.Error(errorMessage, nil, zap.String("journey", "createWorkInfo"))
		return nil, rest_err.NewInternalServerError("database configuration error: work_info collection name not set")
	}
	collection := wr.databaseConnection.Collection(collectionName)

	value := workinfoconv.ConvertWorkInfoDomainToEntity(workInfoDomain) // USO MODIFICADO
	// Em WorkInfoEntity, UserID é mapeado para _id e é uma string.

	_, err := collection.InsertOne(context.Background(), value)
	if err != nil {
		if writeException, ok := err.(mongo.WriteException); ok {
			for _, writeError := range writeException.WriteErrors {
				if writeError.Code == 11000 {
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
		zap.String("userID", workInfoDomain.GetUserId()),
		zap.String("journey", "createWorkInfo"))

	return workInfoDomain, nil
}
