package swap

import (
	"context"
	"fmt"
	"os"

	"github.com/Lipe-Azevedo/escala-fds/src/configuration/logger"
	"github.com/Lipe-Azevedo/escala-fds/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/escala-fds/src/model/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (sr *swapRepository) UpdateSwap(
	id string,
	swapDomain domain.SwapDomainInterface,
) *rest_err.RestErr {
	logger.Info("Init UpdateSwap repository", // "updateSwap" para "UpdateSwap" para consistência
		zap.String("journey", "updateSwap"),
		zap.String("swapID_to_update", id))

	collectionName := os.Getenv(MONGODB_SWAPS_COLLECTION_ENV_KEY)
	if collectionName == "" {
		errorMessage := fmt.Sprintf("Environment variable %s not set for swaps collection name", MONGODB_SWAPS_COLLECTION_ENV_KEY)
		logger.Error(errorMessage, nil, zap.String("journey", "updateSwap"))
		return rest_err.NewInternalServerError("database configuration error: swaps collection name not set")
	}
	collection := sr.databaseConnection.Collection(collectionName)

	objectID, errHex := primitive.ObjectIDFromHex(id)
	if errHex != nil {
		errorMessage := fmt.Sprintf("Invalid Swap ID format for update: %s", id)
		logger.Error(errorMessage, errHex,
			zap.String("journey", "updateSwap"),
			zap.String("swapID", id))
		return rest_err.NewBadRequestError(errorMessage)
	}
	filter := bson.D{{Key: "_id", Value: objectID}}

	// Não usamos o conversor aqui porque o swapDomain pode ser uma atualização parcial
	// ou o serviço já construiu o BSON de atualização específico.
	// O seu código original já montava um bson.M com os campos a serem atualizados.
	// Aqui, o swapDomain deve conter todos os campos que precisam ser setados.
	// A entidade SwapEntity tem campos que são ponteiros (ApprovedAt, ApprovedBy),
	// o que é bom para $set não sobrescrever com valor zero se não fornecido.
	// O conversor ConvertSwapDomainToEntity criaria uma entidade completa.
	// A forma mais segura é construir o bson.M explicitamente com os campos do swapDomain
	// que devem ser atualizados.
	updateFields := bson.M{
		"requester_id":    swapDomain.GetRequesterID(),
		"requested_id":    swapDomain.GetRequestedID(),
		"current_shift":   string(swapDomain.GetCurrentShift()),
		"new_shift":       string(swapDomain.GetNewShift()),
		"current_day_off": string(swapDomain.GetCurrentDayOff()),
		"new_day_off":     string(swapDomain.GetNewDayOff()),
		"status":          string(swapDomain.GetStatus()),
		"reason":          swapDomain.GetReason(),
		// CreatedAt geralmente não é atualizado. Se o seu domain tem, ele será setado.
		// Se o GetCreatedAt() do domain for o valor zero de time.Time e você não quer
		// sobrescrever o valor existente no BD, então ele não deveria estar aqui.
		// A sua SwapDomainInterface não tem um SetCreatedAt, então o valor de createdAt
		// viria do NewSwapDomain ou NewSwapUpdateDomain (que usa time.Now()).
		// Isso significa que cada update iria resetar o createdAt se incluído aqui.
		// É mais seguro omiti-lo do $set a menos que a intenção seja realmente atualizá-lo.
		// "created_at":      swapDomain.GetCreatedAt(), // REMOVIDO - Geralmente não se atualiza
		"approved_at": swapDomain.GetApprovedAt(), // Se for nil, será setado como null (ou omitido se bson omitempty)
		"approved_by": swapDomain.GetApprovedBy(), // Se for nil, será setado como null (ou omitido se bson omitempty)
	}
	// Se approved_at ou approved_by forem nil no domain e a entidade tiver `omitempty`,
	// eles não serão incluídos no BSON pelo $set, o que é bom.
	// Se não tiverem omitempty e forem nil, serão setados para null no MongoDB.

	update := bson.D{{Key: "$set", Value: updateFields}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		logger.Error("Error trying to update swap in repository", err,
			zap.String("journey", "updateSwap"),
			zap.String("swapID", id))
		return rest_err.NewInternalServerError(err.Error())
	}

	if result.MatchedCount == 0 {
		errorMessage := fmt.Sprintf("Swap with ID %s not found for update", id)
		logger.Warn(errorMessage,
			zap.String("journey", "updateSwap"),
			zap.String("swapID", id))
		return rest_err.NewNotFoundError(errorMessage)
	}

	logger.Info("UpdateSwap repository executed successfully",
		zap.String("swapID", id),
		zap.Int64("matchedCount", result.MatchedCount),
		zap.Int64("modifiedCount", result.ModifiedCount),
		zap.String("journey", "updateSwap"))

	return nil
}
