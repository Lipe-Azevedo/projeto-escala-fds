package repository

import (
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	MONGODB_WORK_INFO_DB = "MONGODB_WORK_INFO_DB"
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
