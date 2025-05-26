package user

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

func (ur *userRepository) UpdateUser(
	userId string,
	userDomain domain.UserDomainInterface,
) *rest_err.RestErr {
	logger.Info(
		"Init UpdateUser repository",
		zap.String("journey", "updateUser"),
		zap.String("userId", userId))

	collectionName := os.Getenv(MONGODB_USERS_COLLECTION_ENV_KEY)
	if collectionName == "" {
		errorMessage := fmt.Sprintf("Environment variable %s not set for users collection name", MONGODB_USERS_COLLECTION_ENV_KEY)
		logger.Error(errorMessage, nil, zap.String("journey", "updateUser"))
		return rest_err.NewInternalServerError("database configuration error: users collection name not set")
	}
	collection := ur.databaseConnection.Collection(collectionName)

	userIdHex, errHex := primitive.ObjectIDFromHex(userId)
	if errHex != nil {
		errorMessage := fmt.Sprintf("Invalid userId format for update: %s. Must be a valid hex string.", userId)
		logger.Error(errorMessage, errHex, zap.String("journey", "updateUser"))
		return rest_err.NewBadRequestError(errorMessage)
	}

	updateFields := bson.M{}
	if name := userDomain.GetName(); name != "" {
		updateFields["name"] = name
	}
	if password := userDomain.GetPassword(); password != "" {
		// A senha já deve estar criptografada pela camada de serviço ANTES de chegar aqui.
		updateFields["password"] = password
	}

	if len(updateFields) == 0 {
		logger.Info("No fields to update for user.",
			zap.String("userId", userId),
			zap.String("journey", "updateUser"))
		return nil
	}

	filter := bson.D{{Key: "_id", Value: userIdHex}}
	update := bson.D{{Key: "$set", Value: updateFields}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		logger.Error(
			"Error trying to update user in repository",
			err,
			zap.String("journey", "updateUser"),
			zap.String("userId", userId),
		)
		return rest_err.NewInternalServerError(err.Error())
	}

	if result.MatchedCount == 0 {
		logger.Warn("No user found with the given ID to update in repository",
			zap.String("userId", userId),
			zap.String("journey", "updateUser"))
		return rest_err.NewNotFoundError(fmt.Sprintf("User not found with ID: %s for update", userId))
	}

	logger.Info(
		"UpdateUser repository executed successfully",
		zap.String("userId", userId),
		zap.Int64("matchedCount", result.MatchedCount),
		zap.Int64("modifiedCount", result.ModifiedCount),
		zap.String("journey", "updateUser"),
	)
	return nil
}
