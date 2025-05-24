package user

import (
	"context"
	"fmt"
	"os"

	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/logger"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/rest_err"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (ur *userRepository) DeleteUser(
	userId string,
) *rest_err.RestErr {
	logger.Info(
		"Init DeleteUser repository",
		zap.String("journey", "deleteUser"),
		zap.String("userId", userId))

	collectionName := os.Getenv(MONGODB_USERS_COLLECTION_ENV_KEY)
	if collectionName == "" {
		errMsg := fmt.Sprintf("Environment variable %s not set for users collection name", MONGODB_USERS_COLLECTION_ENV_KEY)
		logger.Error(errMsg, nil, zap.String("journey", "deleteUser"))
		return rest_err.NewInternalServerError("database configuration error: users collection name not set")
	}
	collection := ur.databaseConnection.Collection(collectionName)

	userIdHex, errHex := primitive.ObjectIDFromHex(userId)
	if errHex != nil {
		// Corrigido para retornar o erro aqui, caso o ID seja inválido.
		errMsg := fmt.Sprintf("Invalid userId format for delete: %s. Must be a valid hex string.", userId)
		logger.Error(errMsg, errHex,
			zap.String("journey", "deleteUser"),
			zap.String("userId", userId))
		return rest_err.NewBadRequestError(errMsg)
	}

	filter := bson.D{{Key: "_id", Value: userIdHex}}

	result, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		logger.Error(
			"Error trying to delete user in repository",
			err,
			zap.String("journey", "deleteUser"),
			zap.String("userId", userId),
		)
		return rest_err.NewInternalServerError(err.Error())
	}

	if result.DeletedCount == 0 {
		logger.Warn("No user found with the given ID to delete in repository", // Adicionado "in repository"
			zap.String("userId", userId),
			zap.String("journey", "deleteUser"))
		return rest_err.NewNotFoundError(fmt.Sprintf("User not found with ID: %s for deletion", userId))
	}

	logger.Info(
		"DeleteUser repository executed successfully",
		zap.String("userId", userId),
		zap.String("journey", "deleteUser"),
	)
	return nil
}
