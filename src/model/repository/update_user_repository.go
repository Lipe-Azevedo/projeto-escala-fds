package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/logger"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (ur *userRepository) UpdateUser(
	userId string,
	userDomain model.UserDomainInterface,
) *rest_err.RestErr {
	logger.Info(
		"Init updateUser repository.",
		zap.String("journey", "updateUser"),
	)

	collectionNameKey := MONGODB_USERS_COLLECTION_ENV_KEY
	collectionName := os.Getenv(collectionNameKey)

	if collectionName == "" {
		errorMessage := fmt.Sprintf("Environment variable %s not set for users collection name", collectionNameKey)
		logger.Error(errorMessage, nil, zap.String("journey", "updateUser"))
		return rest_err.NewInternalServerError("database configuration error: users collection name not set")
	}
	collection := ur.dataBaseConnection.Collection(collectionName)

	userIdHex, errHex := primitive.ObjectIDFromHex(userId)
	if errHex != nil {
		errorMessage := fmt.Sprintf("Invalid userId format for update: %s", userId)
		logger.Error(errorMessage, errHex, zap.String("journey", "updateUser"))
		return rest_err.NewBadRequestError(errorMessage)
	}

	updateFields := bson.M{}

	if userDomain.GetName() != "" {
		updateFields["name"] = userDomain.GetName()
	}

	if userDomain.GetPassword() != "" {
		updateFields["password"] = userDomain.GetPassword() // Assumindo que já foi criptografada no serviço
	}

	// Não adicionamos userDomain.GetEmail() ou userDomain.GetUserType() aqui,
	// pois eles não são parte da atualização de perfil do usuário via NewUserUpdateDomain
	// e não queremos zerá-los no banco.

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
			"Error trying to update user.",
			err,
			zap.String("journey", "updateUser"),
		)
		return rest_err.NewInternalServerError(err.Error())
	}

	if result.MatchedCount == 0 {
		logger.Warn("No user found with the given ID to update",
			zap.String("userId", userId),
			zap.String("journey", "updateUser"))
		return rest_err.NewNotFoundError(fmt.Sprintf("User not found with ID: %s for update", userId))
	}

	logger.Info(
		"UpdateUser repository executed successfully.",
		zap.String("userId", userId),
		zap.String("matchedCount", fmt.Sprintf("%d", result.MatchedCount)),
		zap.String("modifiedCount", fmt.Sprintf("%d", result.ModifiedCount)),
		zap.String("journey", "updateUser"),
	)
	return nil
}
