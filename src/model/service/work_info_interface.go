package service

import (
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model/repository"
)

// NewWorkInfoDomainService (alterado para NewWorkInfiDomainService no seu código original, corrigindo para Info)
func NewWorkInfoDomainService( // Corrigido: NewWorkInfiDomainService para NewWorkInfoDomainService
	workInfoRepository repository.WorkInfoRepository,
	userDomainService UserDomainService, // Adicionada dependência
) WorkInfoDomainService {
	return &workInfoDomainService{
		workInfoRepository: workInfoRepository,
		userDomainService:  userDomainService, // Adicionada atribuição
	}
}

type workInfoDomainService struct {
	workInfoRepository repository.WorkInfoRepository
	userDomainService  UserDomainService // Adicionado campo
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
