package swap

import (
	"context"
	"fmt"
	"os"

	"github.com/Lipe-Azevedo/escala-fds/src/configuration/logger"
	"github.com/Lipe-Azevedo/escala-fds/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/escala-fds/src/model/domain"
	swapconv "github.com/Lipe-Azevedo/escala-fds/src/model/repository/entity/converter/swap" // IMPORT MODIFICADO
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func (sr *swapRepository) CreateSwap(
	swapDomain domain.SwapDomainInterface,
) (domain.SwapDomainInterface, *rest_err.RestErr) {
	logger.Info("Init CreateSwap repository",
		zap.String("journey", "createSwap"))

	collectionName := os.Getenv(MONGODB_SWAPS_COLLECTION_ENV_KEY)
	if collectionName == "" {
		errorMessage := fmt.Sprintf("Environment variable %s not set for swaps collection name", MONGODB_SWAPS_COLLECTION_ENV_KEY)
		logger.Error(errorMessage, nil, zap.String("journey", "createSwap"))
		return nil, rest_err.NewInternalServerError("database configuration error: swaps collection name not set")
	}
	collection := sr.databaseConnection.Collection(collectionName)

	value := swapconv.ConvertSwapDomainToEntity(swapDomain) // USO MODIFICADO

	// Se o ID da entidade é omitempty e a string está vazia, o MongoDB gerará um ID.
	// Se o ID já estiver definido no domain (e, portanto, na entidade), o MongoDB tentará usá-lo.
	// A sua SwapEntity.ID é `bson:"_id,omitempty"`.
	// O ConvertSwapDomainToEntity que você forneceu popula `ID: domain.GetID()`.
	// Se domain.GetID() estiver vazio para um novo swap, o MongoDB gerará o ID.
	// Se domain.GetID() NÃO estiver vazio, o MongoDB usará esse valor.
	// A lógica usual é deixar o MongoDB gerar o ID na inserção para novos documentos.
	// No seu construtor NewSwapDomain, o ID não é inicializado, então domain.GetID() será "" para um novo swap.

	result, err := collection.InsertOne(context.Background(), value)
	if err != nil {
		if writeException, ok := err.(mongo.WriteException); ok {
			for _, writeError := range writeException.WriteErrors {
				if writeError.Code == 11000 { // Código para chave duplicada
					errorMessage := fmt.Sprintf("Duplicate key error on creating swap (possibly _id if provided, or other unique index): %s", writeError.Message)
					logger.Error(errorMessage, err,
						zap.String("journey", "createSwap"))
					return nil, rest_err.NewConflictError(errorMessage)
				}
			}
		}
		logger.Error("Error trying to create swap in repository", err,
			zap.String("journey", "createSwap"))
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	// Se o ID foi gerado pelo MongoDB, result.InsertedID será um primitive.ObjectID.
	// A entidade SwapEntity tem ID como string.
	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		value.ID = oid.Hex() // Atualiza o ID na entidade com o valor gerado e convertido para Hex string
	} else if strId, ok := result.InsertedID.(string); ok {
		value.ID = strId // Se por algum motivo o InsertedID já for uma string (menos comum para _id omitido)
	}
	// Se o ID já estava em 'value' e foi usado, InsertedID pode ser esse mesmo valor.

	logger.Info("CreateSwap repository executed successfully",
		zap.String("swapID", value.ID), // value.ID agora deve ser a string correta do ID
		zap.String("journey", "createSwap"))

	return swapconv.ConvertSwapEntityToDomain(*value), nil // USO MODIFICADO
}
