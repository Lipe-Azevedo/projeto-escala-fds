package repository

import (
	"context"

	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/logger"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model/repository/entity/converter"
)

func (ur *userRepository) CreateWorkInfo(
	workInfoDomain model.WorkInfoDomainInterface,
) (model.WorkInfoDomainInterface, *rest_err.RestErr) {
	collection := ur.dataBaseConnection.Collection("work_infos") // Coleção separada

	value := converter.ConvertWorkInfoDomainToEntity(workInfoDomain)

	_, err := collection.InsertOne(context.Background(), value)
	if err != nil {
		logger.Error("Error creating work info", err)
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	return workInfoDomain, nil
}
