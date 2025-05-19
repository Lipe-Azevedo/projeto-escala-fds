package repository

import (
	"context"
	"os"

	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/logger"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model/repository/entity"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func (ur *userRepository) FindWorkInfoByUserId(
	userId string,
) (model.WorkInfoDomainInterface, *rest_err.RestErr) {
	logger.Info(
		"Init findWorkInfoByID repository.",
		zap.String("journey", "createUser"))

	collection_name := os.Getenv(MONGODB_WORK_INFO_DB)

	collection := ur.dataBaseConnection.Collection(collection_name)

	workInfoEntity := &entity.WorkInfoEntity{}
	filter := bson.M{"user_id": userId}

	err := collection.FindOne(context.Background(), filter).Decode(workInfoEntity)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, rest_err.NewNotFoundError("Work info not found")
		}
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	return converter.ConvertWorkInfoEntityToDomain(*workInfoEntity), nil
}
