package workinfo

import (
	"context"
	"fmt"
	"os"

	"github.com/Lipe-Azevedo/escala-fds/src/configuration/logger"
	"github.com/Lipe-Azevedo/escala-fds/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/escala-fds/src/model/domain"
	workinfoconv "github.com/Lipe-Azevedo/escala-fds/src/model/repository/entity/converter/workinfo" // IMPORT MODIFICADO
	"go.mongodb.org/mongo-driver/bson"
	"go.uber.org/zap"
	// "go.mongodb.org/mongo-driver/mongo/options" // Não usado aqui, mas pode ser para Upsert
)

func (wr *workInfoRepository) UpdateWorkInfo(
	userId string,
	workInfoDomain domain.WorkInfoDomainInterface,
) *rest_err.RestErr {
	logger.Info(
		"Init UpdateWorkInfo repository",
		zap.String("journey", "updateWorkInfo"),
		zap.String("userId", userId))

	collectionName := os.Getenv(MONGODB_WORKINFO_COLLECTION_ENV_KEY)
	if collectionName == "" {
		errorMessage := fmt.Sprintf("Environment variable %s not set for work_info collection name", MONGODB_WORKINFO_COLLECTION_ENV_KEY)
		logger.Error(errorMessage, nil, zap.String("journey", "updateWorkInfo"))
		return rest_err.NewInternalServerError("database configuration error: work_info collection name not set")
	}
	collection := wr.databaseConnection.Collection(collectionName)

	// workInfoDomain deve conter o estado COMPLETO e ATUALIZADO do WorkInfo.
	// O userID no workInfoDomain deve corresponder ao _id do documento a ser atualizado.
	// ConvertWorkInfoDomainToEntity irá mapear workInfoDomain.GetUserId() para WorkInfoEntity.UserID,
	// que é o campo `bson:"_id"`.
	value := workinfoconv.ConvertWorkInfoDomainToEntity(workInfoDomain) // USO MODIFICADO

	// O filtro é pelo _id, que é o userId fornecido como parâmetro (e deve ser o mesmo que workInfoDomain.GetUserId()).
	filter := bson.M{"_id": userId}

	// $set garante que apenas os campos fornecidos em 'value' (que é a entidade completa) sejam atualizados/substituídos.
	// Como WorkInfoEntity não tem 'omitempty' na maioria dos campos, ele tentará setar todos.
	update := bson.M{"$set": value}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		logger.Error("Error trying to update WorkInfo in repository", err,
			zap.String("journey", "updateWorkInfo"),
			zap.String("userId", userId))
		return rest_err.NewInternalServerError(err.Error())
	}

	if result.MatchedCount == 0 {
		logger.Warn("No WorkInfo found with the given user ID to update in repository",
			zap.String("userId", userId),
			zap.String("journey", "updateWorkInfo"))
		return rest_err.NewNotFoundError(fmt.Sprintf("WorkInfo for user ID %s not found for update", userId))
	}

	if result.ModifiedCount == 0 && result.MatchedCount == 1 {
		logger.Info("WorkInfo found but no fields were modified",
			zap.String("userId", userId),
			zap.String("journey", "updateWorkInfo"))
	}

	logger.Info("UpdateWorkInfo repository executed successfully",
		zap.String("userId", userId),
		zap.Int64("matchedCount", result.MatchedCount),
		zap.Int64("modifiedCount", result.ModifiedCount),
		zap.String("journey", "updateWorkInfo"))

	return nil
}
