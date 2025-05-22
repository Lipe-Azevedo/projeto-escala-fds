package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/logger"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model/repository/entity"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson"
	"go.uber.org/zap"
)

func (ur *userRepository) FindAllUsers() ([]model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindAllUsers repository",
		zap.String("journey", "findAllUsers"))

	collectionNameKey := MONGODB_USERS_COLLECTION_ENV_KEY
	collectionName := os.Getenv(collectionNameKey)

	if collectionName == "" {
		errorMessage := fmt.Sprintf("Environment variable %s not set for users collection name", collectionNameKey)
		logger.Error(errorMessage, nil, zap.String("journey", "findAllUsers"))
		return nil, rest_err.NewInternalServerError("database configuration error: users collection name not set")
	}
	collection := ur.dataBaseConnection.Collection(collectionName)

	// Usar um filtro vazio para buscar todos os documentos
	filter := bson.D{}
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		logger.Error("Error finding all users in repository", err,
			zap.String("journey", "findAllUsers"))
		return nil, rest_err.NewInternalServerError(fmt.Sprintf("Error finding all users: %s", err.Error()))
	}
	defer cursor.Close(context.Background())

	var userEntities []entity.UserEntity
	if err = cursor.All(context.Background(), &userEntities); err != nil {
		logger.Error("Error decoding all users from cursor", err,
			zap.String("journey", "findAllUsers"))
		return nil, rest_err.NewInternalServerError(fmt.Sprintf("Error decoding all users: %s", err.Error()))
	}

	var userDomains []model.UserDomainInterface
	for _, ue := range userEntities {
		userDomains = append(userDomains, converter.ConvertEntityToDomain(ue))
	}

	logger.Info("FindAllUsers repository executed successfully",
		zap.Int("count", len(userDomains)),
		zap.String("journey", "findAllUsers"))

	return userDomains, nil
}
