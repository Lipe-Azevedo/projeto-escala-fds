package service

import (
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model/repository"
)

func NewWorkInfiDomainService(
	workInfoRepository repository.WorkInfoRepository,
) WorkInfoDomainService {
	return &workInfoDomainService{workInfoRepository}
}

type workInfoDomainService struct {
	workInfoRepository repository.WorkInfoRepository
}

type WorkInfoDomainService interface {
	CreateWorkInfoServices(
		workInfoDomain model.WorkInfoDomainInterface,
	) (model.WorkInfoDomainInterface, *rest_err.RestErr)

	FindWorkInfoByUserIdServices(
		userId string,
	) (model.WorkInfoDomainInterface, *rest_err.RestErr)

	UpdateWorkInfoServices(
		userId string,
		workInfoDomain model.WorkInfoDomainInterface,
	) *rest_err.RestErr
}
