package user

import (
	"context"
	"fmt"
	"os"

	"github.com/Lipe-Azevedo/escala-fds/src/configuration/logger"
	"github.com/Lipe-Azevedo/escala-fds/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/escala-fds/src/model/domain"
	"github.com/Lipe-Azevedo/escala-fds/src/model/repository/entity"
	userconv "github.com/Lipe-Azevedo/escala-fds/src/model/repository/entity/converter/user" // IMPORT MODIFICADO

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func (ur *userRepository) FindUserByID(id string) (domain.UserDomainInterface, *rest_err.RestErr) {
	logger.Info(
		"Init FindUserByID repository",
		zap.String("journey", "findUserByID"),
		zap.String("userIdToFind", id))

	collectionName := os.Getenv(MONGODB_USERS_COLLECTION_ENV_KEY)
	if collectionName == "" {
		errMsg := fmt.Sprintf("Environment variable %s not set for users collection name", MONGODB_USERS_COLLECTION_ENV_KEY)
		logger.Error(errMsg, nil, zap.String("journey", "findUserByID"))
		return nil, rest_err.NewInternalServerError("database configuration error: users collection name not set")
	}
	collection := ur.databaseConnection.Collection(collectionName)
	userEntity := &entity.UserEntity{}
	objectId, errHex := primitive.ObjectIDFromHex(id)
	if errHex != nil {
		errorMessage := fmt.Sprintf("Invalid userId format: %s. Must be a valid hex string.", id)
		logger.Error(errorMessage, errHex, zap.String("journey", "findUserByID"), zap.String("userId", id))
		return nil, rest_err.NewBadRequestError(errorMessage)
	}
	filter := bson.D{{Key: "_id", Value: objectId}}
	err := collection.FindOne(context.Background(), filter).Decode(userEntity)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf("User not found with ID: %s", id)
			logger.Warn(errorMessage, zap.String("journey", "findUserByID"), zap.String("userId", id))
			return nil, rest_err.NewNotFoundError(errorMessage)
		}
		errorMessage := fmt.Sprintf("Error trying to find user by ID in repository: %s", id)
		logger.Error(errorMessage, err, zap.String("journey", "findUserByID"))
		return nil, rest_err.NewInternalServerError(fmt.Sprintf("Error finding user by ID: %s", err.Error()))
	}
	logger.Info("FindUserByID repository executed successfully", zap.String("journey", "findUserByID"), zap.String("userId", userEntity.ID.Hex()))
	return userconv.ConvertEntityToDomain(*userEntity), nil // USO MODIFICADO
}

func (ur *userRepository) FindUserByEmail(email string) (domain.UserDomainInterface, *rest_err.RestErr) {
	logger.Info(
		"Init FindUserByEmail repository",
		zap.String("journey", "findUserByEmail"),
		zap.String("emailToFind", email))

	collectionName := os.Getenv(MONGODB_USERS_COLLECTION_ENV_KEY)
	if collectionName == "" {
		errMsg := fmt.Sprintf("Environment variable %s not set for users collection name", MONGODB_USERS_COLLECTION_ENV_KEY)
		logger.Error(errMsg, nil, zap.String("journey", "findUserByEmail"))
		return nil, rest_err.NewInternalServerError("database configuration error: users collection name not set")
	}
	collection := ur.databaseConnection.Collection(collectionName)
	userEntity := &entity.UserEntity{}
	filter := bson.D{{Key: "email", Value: email}}
	err := collection.FindOne(context.Background(), filter).Decode(userEntity)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf("User not found with email: %s", email)
			logger.Warn(errorMessage, zap.String("journey", "findUserByEmail"), zap.String("email", email))
			return nil, rest_err.NewNotFoundError(errorMessage)
		}
		errorMessage := fmt.Sprintf("Error trying to find user by email in repository: %s", email)
		logger.Error(errorMessage, err, zap.String("journey", "findUserByEmail"))
		return nil, rest_err.NewInternalServerError(fmt.Sprintf("Error finding user by email: %s", err.Error()))
	}
	logger.Info("FindUserByEmail repository executed successfully", zap.String("journey", "findUserByEmail"), zap.String("email", email), zap.String("userId", userEntity.ID.Hex()))
	return userconv.ConvertEntityToDomain(*userEntity), nil // USO MODIFICADO
}

func (ur *userRepository) FindAllUsers() ([]domain.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindAllUsers repository",
		zap.String("journey", "findAllUsers"))

	collectionName := os.Getenv(MONGODB_USERS_COLLECTION_ENV_KEY)
	if collectionName == "" {
		errorMessage := fmt.Sprintf("Environment variable %s not set for users collection name", MONGODB_USERS_COLLECTION_ENV_KEY)
		logger.Error(errorMessage, nil, zap.String("journey", "findAllUsers"))
		return nil, rest_err.NewInternalServerError("database configuration error: users collection name not set")
	}
	collection := ur.databaseConnection.Collection(collectionName)

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
		logger.Error("Error decoding all users from cursor in repository", err,
			zap.String("journey", "findAllUsers"))
		return nil, rest_err.NewInternalServerError(fmt.Sprintf("Error decoding all users: %s", err.Error()))
	}

	var userDomains []domain.UserDomainInterface
	for _, ue := range userEntities {
		userDomains = append(userDomains, userconv.ConvertEntityToDomain(ue)) // USO MODIFICADO
	}

	logger.Info("FindAllUsers repository executed successfully",
		zap.Int("count", len(userDomains)),
		zap.String("journey", "findAllUsers"))

	return userDomains, nil
}
