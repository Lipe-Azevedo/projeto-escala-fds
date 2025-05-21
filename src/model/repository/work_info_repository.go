package repository

import (
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	// MONGODB_WORKINFO_COLLECTION_ENV_KEY armazena o nome da variável de ambiente que contém o nome da coleção work_info.
	MONGODB_WORKINFO_COLLECTION_ENV_KEY = "MONGODB_WORKINFO_COLLECTION"
)

func NewWorkInfoRepository(
	database *mongo.Database,
) WorkInfoRepository {
	return &workInfoRepository{
		database,
	}
}

type workInfoRepository struct {
	dataBaseConnection *mongo.Database
}

type WorkInfoRepository interface {
	CreateWorkInfo(
		workInfoDomain model.WorkInfoDomainInterface,
	) (model.WorkInfoDomainInterface, *rest_err.RestErr)

	FindWorkInfoByUserId(
		userId string,
	) (model.WorkInfoDomainInterface, *rest_err.RestErr)

	UpdateWorkInfo(
		userId string,
		workInfoDomain model.WorkInfoDomainInterface,
	) *rest_err.RestErr
}
