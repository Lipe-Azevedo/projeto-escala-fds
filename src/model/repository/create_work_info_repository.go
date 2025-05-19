package repository

import (
	"context"
	"os"

	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/logger"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model/repository/entity/converter"
	"go.uber.org/zap"
)

func (wr *workInfoRepository)CreateWorkInfo(
	workInfoDomain model.WorkInfoDomainInterface,
) (model.WorkInfoDomainInterface, *rest_err.RestErr) {
	logger.Info(
		"Init createWorkInfo repository.",
		zap.String("journey", "createUser"))

	collection_name := os.Getenv(MONGODB_WORK_INFO_DB)

	collection := wr.dataBaseConnection.Collection(collection_name)

	value := converter.ConvertWorkInfoDomainToEntity(workInfoDomain)

	_, err := collection.InsertOne(context.Background(), value)
	if err != nil {
		logger.Error("Error creating work info", err)
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	return workInfoDomain, nil
}
