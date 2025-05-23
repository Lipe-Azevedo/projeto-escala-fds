package repository

import (
	"context"
	"os" // Faltava fmt para Sprintf, mas não é usado diretamente aqui. Adicionado para consistência com outros arquivos.

	"fmt" // Adicionado para Sprintf

	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/logger"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo" // Adicionar para mongo.WriteException
	"go.uber.org/zap"
)

func (sr *swapRepository) CreateSwap(
	swapDomain model.SwapDomainInterface,
) (model.SwapDomainInterface, *rest_err.RestErr) {
	logger.Info("Init createSwap repository",
		zap.String("journey", "createSwap"))

	// collection_name foi renomeado para collectionName para seguir o padrão Go.
	collectionNameKey := MONGODB_SWAPS_COLLECTION_ENV_KEY // Deve ser definido em swap_repository.go
	collectionName := os.Getenv(collectionNameKey)

	if collectionName == "" { // Adicionada verificação de collectionName vazio
		errorMessage := fmt.Sprintf("Environment variable %s not set for swaps collection name", collectionNameKey)
		logger.Error(errorMessage, nil, zap.String("journey", "createSwap"))
		return nil, rest_err.NewInternalServerError("database configuration error: swaps collection name not set")
	}
	collection := sr.databaseConnection.Collection(collectionName)

	value := converter.ConvertSwapDomainToEntity(swapDomain)
	// O ID no swapDomain (se preenchido antes de chegar aqui) seria usado se a tag _id na entidade SwapEntity
	// não tivesse 'omitempty' e fosse preenchido no 'value'.
	// No entanto, a prática comum para InsertOne é deixar o MongoDB gerar o _id.
	// A sua lógica atual atribui o ID gerado DEPOIS da inserção.

	result, err := collection.InsertOne(context.Background(), value)
	if err != nil {
		// Verifica se o erro é uma exceção de escrita do MongoDB
		// Embora para InsertOne, um erro de chave duplicada no _id seja menos comum
		// se o _id não estiver sendo pré-definido no 'value' (o que não parece ser o caso aqui,
		// já que você atribui o InsertedID depois).
		// Mas, se houvesse outro índice único que pudesse ser violado:
		if writeException, ok := err.(mongo.WriteException); ok {
			for _, writeError := range writeException.WriteErrors {
				if writeError.Code == 11000 { // Erro de chave duplicada
					// Adapte a mensagem para o contexto de Swap, se aplicável
					// Por exemplo, se houvesse um índice único em (requester_id, requested_id, day)
					errorMessage := fmt.Sprintf("Duplicate key error on creating swap: %s", writeError.Message)
					logger.Error(errorMessage, err,
						zap.String("journey", "createSwap"))
					return nil, rest_err.NewConflictError(errorMessage)
				}
			}
		}
		logger.Error("Error trying to create swap", // "shift swap" para "swap"
			err,
			zap.String("journey", "createSwap"))
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	// Atribui o ID gerado pelo MongoDB (convertido para Hex string) ao campo ID da entidade 'value'.
	// A entidade SwapEntity tem um campo ID string `bson:"_id,omitempty"`.
	generatedID := result.InsertedID.(primitive.ObjectID)
	value.ID = generatedID.Hex()

	logger.Info("CreateSwap repository executed successfully",
		zap.String("swapID", value.ID),
		zap.String("journey", "createSwap"))

	// Retorna um novo domain convertido a partir da entidade 'value', que agora possui o ID do banco.
	return converter.ConvertSwapEntityToDomain(*value), nil
}
