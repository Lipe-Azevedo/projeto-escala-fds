package repository

import (
	"context"
	"os" // Faltava fmt para Sprintf nos logs, mas não é usado no código fornecido diretamente.

	"fmt" // Adicionado para Sprintf

	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/logger"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive" // Adicionar para ObjectIDFromHex
	"go.uber.org/zap"
)

func (sr *swapRepository) UpdateSwap(
	id string, // Este ID é a string hexadecimal do _id do Swap
	swapDomain model.SwapDomainInterface,
) *rest_err.RestErr {
	logger.Info("Init updateSwap repository",
		zap.String("journey", "updateSwap"),
		zap.String("swapID_to_update", id)) // Log do ID que será atualizado

	collectionNameKey := MONGODB_SWAPS_COLLECTION_ENV_KEY
	collectionName := os.Getenv(collectionNameKey)

	if collectionName == "" { // Adicionada verificação de collectionName vazio
		errorMessage := fmt.Sprintf("Environment variable %s not set for swaps collection name", collectionNameKey)
		logger.Error(errorMessage, nil, zap.String("journey", "updateSwap"))
		return rest_err.NewInternalServerError("database configuration error: swaps collection name not set")
	}
	collection := sr.databaseConnection.Collection(collectionName)

	// 1. Converter o ID da string para primitive.ObjectID para o filtro
	objectID, errHex := primitive.ObjectIDFromHex(id)
	if errHex != nil {
		errorMessage := fmt.Sprintf("Invalid Swap ID format: %s", id)
		logger.Error(errorMessage, errHex,
			zap.String("journey", "updateSwap"),
			zap.String("swapID", id))
		return rest_err.NewBadRequestError(errorMessage)
	}
	filter := bson.D{{Key: "_id", Value: objectID}} // Usar ObjectID no filtro

	// 2. Construir o documento de atualização para $set
	// É crucial não tentar dar $set no campo _id.
	// A entidade SwapEntity retornada por ConvertSwapDomainToEntity inclui o ID.
	// O bson marshaller com `omitempty` na tag _id da entidade PODE lidar com isso,
	// mas é mais seguro e explícito construir um bson.M apenas com os campos que devem ser atualizados.
	updateFields := bson.M{
		"requester_id":    swapDomain.GetRequesterID(),
		"requested_id":    swapDomain.GetRequestedID(),
		"current_shift":   string(swapDomain.GetCurrentShift()),
		"new_shift":       string(swapDomain.GetNewShift()),
		"current_day_off": string(swapDomain.GetCurrentDayOff()),
		"new_day_off":     string(swapDomain.GetNewDayOff()),
		"status":          string(swapDomain.GetStatus()),
		"reason":          swapDomain.GetReason(),
		"created_at":      swapDomain.GetCreatedAt(), // Geralmente não se atualiza CreatedAt
		"approved_at":     swapDomain.GetApprovedAt(),
		"approved_by":     swapDomain.GetApprovedBy(),
	}
	// Se CreatedAt não deve ser alterado em um update, remova-o de updateFields.
	// A struct SwapEntity tem CreatedAt. Se o objetivo é que o $set só atualize os campos
	// que realmente mudam (como status, approvedAt, approvedBy), o serviço deveria passar
	// um domain que reflete apenas essas mudanças, ou o repositório constrói o updateFields
	// de forma ainda mais seletiva. Assumindo que swapDomain contém o estado final desejado (exceto _id).

	update := bson.D{{Key: "$set", Value: updateFields}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		logger.Error("Error trying to update swap in repository", // "shift swap" para "swap"
			err,
			zap.String("journey", "updateSwap"),
			zap.String("swapID", id))
		return rest_err.NewInternalServerError(err.Error())
	}

	// 3. Verificar se algum documento foi realmente encontrado e atualizado
	if result.MatchedCount == 0 {
		errorMessage := fmt.Sprintf("Swap with ID %s not found for update", id)
		logger.Warn(errorMessage, // Usar Warn para "não encontrado"
			zap.String("journey", "updateSwap"),
			zap.String("swapID", id))
		return rest_err.NewNotFoundError(errorMessage)
	}

	logger.Info("UpdateSwap repository executed successfully",
		zap.String("swapID", id),
		zap.Int64("matchedCount", result.MatchedCount), // Logar counts
		zap.Int64("modifiedCount", result.ModifiedCount),
		zap.String("journey", "updateSwap"))

	return nil
}
