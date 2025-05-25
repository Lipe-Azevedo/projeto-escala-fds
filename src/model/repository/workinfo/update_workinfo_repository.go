package workinfo

import (
	"context"
	"fmt"
	"os"

	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/logger"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model/repository/entity/converter" // Necessário para converter domain para entity
	"go.mongodb.org/mongo-driver/bson"                                                  // Para Upsert
	"go.uber.org/zap"
)

func (wr *workInfoRepository) UpdateWorkInfo(
	userId string, // Este é o _id do WorkInfo a ser atualizado
	workInfoDomain model.WorkInfoDomainInterface,
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

	// O workInfoDomain contém o estado COMPLETO desejado para o WorkInfo.
	// ConvertDomainToEntity irá mapear workInfoDomain.GetUserId() para o campo UserID na entidade,
	// que por sua vez é mapeado para _id no BSON.
	value := converter.ConvertWorkInfoDomainToEntity(workInfoDomain)

	// O filtro é pelo _id, que é o userId fornecido.
	filter := bson.M{"_id": userId}
	// Os dados para atualização são todos os campos da entidade (exceto _id, que está no filtro).
	// $set garante que apenas os campos fornecidos em 'value' sejam atualizados.
	// Como WorkInfoEntity não tem omitempty em todos os campos, ele tentará setar todos.
	// Se o WorkInfoDomain passado for o objeto completo e atualizado, isso funciona como um replace.
	update := bson.M{"$set": value}

	// Poderíamos usar UpdateOne. Se a intenção for criar se não existir (Upsert),
	// o serviço deveria ter essa lógica e chamar CreateWorkInfo ou UpdateWorkInfo.
	// O arquivo original usava UpdateOne sem upsert.
	// Se a lógica de atualização no serviço carrega o existente, modifica e depois salva,
	// UpdateOne é apropriado.

	// No seu código original, o update do repositório tinha um bug:
	// filter := bson.M{"user_id": userId}
	// Isso está incorreto porque o campo no MongoDB é "_id", não "user_id".
	// O WorkInfoEntity tem `UserID string bson:"_id"`.
	// A correção é usar bson.M{"_id": userId} no filtro.

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
		// Importante: O serviço de update em seu código original busca o WorkInfo primeiro.
		// Se não encontrar, ele retorna um erro de "não encontrado" ANTES de chamar o repositório de update.
		// Então, este erro de MatchedCount == 0 no repositório de update não deveria ocorrer
		// se a lógica do serviço estiver correta. No entanto, é uma boa salvaguarda.
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
