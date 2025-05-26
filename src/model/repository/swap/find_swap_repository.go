package swap

import (
	"context"
	"fmt"
	"os"

	"github.com/Lipe-Azevedo/escala-fds/src/configuration/logger"
	"github.com/Lipe-Azevedo/escala-fds/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/escala-fds/src/model"
	"github.com/Lipe-Azevedo/escala-fds/src/model/repository/entity"           // Entidades ainda globais
	"github.com/Lipe-Azevedo/escala-fds/src/model/repository/entity/converter" // Conversores ainda globais
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func (sr *swapRepository) FindSwapByID(
	id string,
) (model.SwapDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindSwapByID repository",
		zap.String("journey", "findSwapByID"),
		zap.String("swapIDToFind", id))

	collectionName := os.Getenv(MONGODB_SWAPS_COLLECTION_ENV_KEY)
	if collectionName == "" {
		errorMessage := fmt.Sprintf("Environment variable %s not set for swaps collection name", MONGODB_SWAPS_COLLECTION_ENV_KEY)
		logger.Error(errorMessage, nil, zap.String("journey", "findSwapByID"))
		return nil, rest_err.NewInternalServerError("database configuration error: swaps collection name not set")
	}
	collection := sr.databaseConnection.Collection(collectionName)

	swapEntity := &entity.SwapEntity{} // Usando entidade global

	objectID, errHex := primitive.ObjectIDFromHex(id)
	if errHex != nil {
		errorMessage := fmt.Sprintf("Invalid Swap ID format: %s", id)
		logger.Error(errorMessage, errHex,
			zap.String("journey", "findSwapByID"),
			zap.String("swapID", id))
		return nil, rest_err.NewBadRequestError(errorMessage)
	}

	filter := bson.D{{Key: "_id", Value: objectID}}
	err := collection.FindOne(context.Background(), filter).Decode(swapEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf("Swap not found with ID: %s", id)
			logger.Warn(errorMessage,
				zap.String("journey", "findSwapByID"),
				zap.String("swapID", id))
			return nil, rest_err.NewNotFoundError(errorMessage)
		}
		logger.Error("Error trying to find swap by ID in repository", err,
			zap.String("journey", "findSwapByID"),
			zap.String("swapID", id))
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	logger.Info("FindSwapByID repository executed successfully",
		zap.String("swapID", swapEntity.ID),
		zap.String("journey", "findSwapByID"))

	return converter.ConvertSwapEntityToDomain(*swapEntity), nil // Usando conversor global
}

func (sr *swapRepository) FindSwapsByUserID(
	userID string,
) ([]model.SwapDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindSwapsByUserID repository",
		zap.String("journey", "findSwapsByUserID"), // Corrigido journey para corresponder à função
		zap.String("userID", userID))

	collectionName := os.Getenv(MONGODB_SWAPS_COLLECTION_ENV_KEY)
	if collectionName == "" {
		errorMessage := fmt.Sprintf("Environment variable %s not set for swaps collection name", MONGODB_SWAPS_COLLECTION_ENV_KEY)
		logger.Error(errorMessage, nil, zap.String("journey", "findSwapsByUserID"))
		return nil, rest_err.NewInternalServerError("database configuration error: swaps collection name not set")
	}
	collection := sr.databaseConnection.Collection(collectionName)

	filter := bson.M{
		"$or": []bson.M{
			{"requester_id": userID},
			{"requested_id": userID},
		},
	}

	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		logger.Error("Error finding swaps by user ID in repository", err,
			zap.String("journey", "findSwapsByUserID"),
			zap.String("userID", userID))
		return nil, rest_err.NewInternalServerError(fmt.Sprintf("Error finding swaps for user %s: %s", userID, err.Error()))
	}
	defer cursor.Close(context.Background())

	var swapEntities []entity.SwapEntity // Usando entidade global
	if err = cursor.All(context.Background(), &swapEntities); err != nil {
		logger.Error("Error decoding swaps by user ID from cursor", err,
			zap.String("journey", "findSwapsByUserID"),
			zap.String("userID", userID))
		return nil, rest_err.NewInternalServerError(fmt.Sprintf("Error decoding swaps for user %s: %s", userID, err.Error()))
	}

	var swapDomains []model.SwapDomainInterface
	for _, se := range swapEntities {
		domain := converter.ConvertSwapEntityToDomain(se) // Usando conversor global
		swapDomains = append(swapDomains, domain)
	}

	logger.Info("Successfully found swaps by userID in repository",
		zap.String("userID", userID),
		zap.Int("count", len(swapDomains)),
		zap.String("journey", "findSwapsByUserID"))

	return swapDomains, nil
}

func (sr *swapRepository) FindSwapsByStatus(
	status model.SwapStatus,
) ([]model.SwapDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindSwapsByStatus repository",
		zap.String("journey", "findSwapsByStatus"),
		zap.String("status", string(status)))

	collectionName := os.Getenv(MONGODB_SWAPS_COLLECTION_ENV_KEY)
	if collectionName == "" {
		errorMessage := fmt.Sprintf("Environment variable %s not set for swaps collection name", MONGODB_SWAPS_COLLECTION_ENV_KEY)
		logger.Error(errorMessage, nil, zap.String("journey", "findSwapsByStatus"))
		return nil, rest_err.NewInternalServerError("database configuration error: swaps collection name not set")
	}
	collection := sr.databaseConnection.Collection(collectionName)

	filter := bson.M{"status": string(status)}

	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		logger.Error("Error finding swaps by status in repository", err,
			zap.String("journey", "findSwapsByStatus"),
			zap.String("status", string(status)))
		return nil, rest_err.NewInternalServerError(fmt.Sprintf("Error finding swaps with status %s: %s", status, err.Error()))
	}
	defer cursor.Close(context.Background())

	var swapEntities []entity.SwapEntity // Usando entidade global
	if err = cursor.All(context.Background(), &swapEntities); err != nil {
		logger.Error("Error decoding swaps by status from cursor", err,
			zap.String("journey", "findSwapsByStatus"),
			zap.String("status", string(status)))
		return nil, rest_err.NewInternalServerError(fmt.Sprintf("Error decoding swaps with status %s: %s", status, err.Error()))
	}

	var swapDomains []model.SwapDomainInterface
	for _, se := range swapEntities {
		swapDomains = append(swapDomains, converter.ConvertSwapEntityToDomain(se)) // Usando conversor global
	}

	logger.Info("Successfully found swaps by status in repository",
		zap.String("status", string(status)),
		zap.Int("count", len(swapDomains)),
		zap.String("journey", "findSwapsByStatus"))

	return swapDomains, nil
}
