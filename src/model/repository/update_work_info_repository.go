package repository

import (
	"context"
	"os"

	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/logger"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson"
	"go.uber.org/zap"
)

func (ur *userRepository) UpdateWorkInfo(
	userId string,
	workInfoDomain model.WorkInfoDomainInterface,
) *rest_err.RestErr {
	logger.Info(
		"Init updateWorkInfo repository.",
		zap.String("journey", "createUser"))

	collection_name := os.Getenv(MONGODB_WORK_INFO_DB)

	collection := ur.dataBaseConnection.Collection(collection_name)

	value := converter.ConvertWorkInfoDomainToEntity(workInfoDomain)
	filter := bson.M{"user_id": userId}
	update := bson.M{"$set": value}

	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return rest_err.NewInternalServerError(err.Error())
	}

	return nil
}
