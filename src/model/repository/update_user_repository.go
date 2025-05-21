package repository

import (
	"context"
	"os"

	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/logger"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model/repository/entity/converter"
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

	collection_name := os.Getenv(MONGODB_USERS_COLLECTION_ENV_KEY) // Usa a chave definida em user_repository.go
	collection := ur.dataBaseConnection.Collection(collection_name)

	value := converter.ConvertDomainToEntity(userDomain)
	userIdHex, _ := primitive.ObjectIDFromHex(userId)

	filter := bson.D{{Key: "_id", Value: userIdHex}}
	update := bson.D{{Key: "$set", Value: value}}

	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		logger.Error(
			"Error trying to update user.",
			err,
			zap.String("journey", "updateUser"),
		)
		return rest_err.NewInternalServerError(err.Error())
	}

	logger.Info(
		"updateUser repository executed suceeefully.",
		zap.String("userId", userId),
		zap.String("journey", "updateUser"),
	)
	return nil
}
