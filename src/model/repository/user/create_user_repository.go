package user

import (
	"context"
	"fmt"
	"os"

	"github.com/Lipe-Azevedo/escala-fds/src/configuration/logger"
	"github.com/Lipe-Azevedo/escala-fds/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/escala-fds/src/model"
	"github.com/Lipe-Azevedo/escala-fds/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func (ur *userRepository) CreateUser(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info(
		"Init CreateUser repository",
		zap.String("journey", "createUser"),
		zap.String("email", userDomain.GetEmail()))

	collectionName := os.Getenv(MONGODB_USERS_COLLECTION_ENV_KEY)
	if collectionName == "" {
		errMsg := fmt.Sprintf("Environment variable %s not set for users collection name", MONGODB_USERS_COLLECTION_ENV_KEY)
		logger.Error(errMsg, nil, zap.String("journey", "createUser"))
		return nil, rest_err.NewInternalServerError("database configuration error: users collection name not set")
	}
	collection := ur.databaseConnection.Collection(collectionName)

	value := converter.ConvertDomainToEntity(userDomain)

	result, err := collection.InsertOne(context.Background(), value)
	if err != nil {
		if writeException, ok := err.(mongo.WriteException); ok {
			for _, writeError := range writeException.WriteErrors {
				if writeError.Code == 11000 { // Código de erro para chave duplicada
					// Verifica se o erro de duplicidade é no índice do email (comum)
					// A mensagem de erro do MongoDB geralmente indica qual chave causou o problema.
					errorMessage := fmt.Sprintf("User with email %s already exists or another unique constraint violated.", value.Email)
					// Poderia analisar writeError.Message para ser mais específico se necessário.
					logger.Error(errorMessage, err,
						zap.String("journey", "createUser"),
						zap.String("email", value.Email))
					return nil, rest_err.NewConflictError(errorMessage)
				}
			}
		}
		logger.Error(
			"Error trying to create user in repository",
			err,
			zap.String("journey", "createUser"))
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	value.ID = result.InsertedID.(primitive.ObjectID)
	logger.Info(
		"CreateUser repository executed successfully",
		zap.String("userId", value.ID.Hex()),
		zap.String("journey", "createUser"),
	)

	return converter.ConvertEntityToDomain(*value), nil
}
