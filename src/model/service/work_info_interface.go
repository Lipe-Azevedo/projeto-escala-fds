package service

import (
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/controller/model/request" // Importação para request.WorkInfoUpdateRequest
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model/repository"
)

func NewWorkInfoDomainService(
	workInfoRepository repository.WorkInfoRepository,
	userDomainService UserDomainService,
) WorkInfoDomainService {
	return &workInfoDomainService{
		workInfoRepository: workInfoRepository,
		userDomainService:  userDomainService,
	}
}

type workInfoDomainService struct {
	workInfoRepository repository.WorkInfoRepository
	userDomainService  UserDomainService
}

type WorkInfoDomainService interface {
	CreateWorkInfoServices(
		workInfoDomain model.WorkInfoDomainInterface, // Para criação, usa WorkInfoRequest (controller converte para domain)
	) (model.WorkInfoDomainInterface, *rest_err.RestErr)

	FindWorkInfoByUserIdServices(
		userId string,
	) (model.WorkInfoDomainInterface, *rest_err.RestErr)

	// UpdateWorkInfoServices agora lida com atualizações parciais via PUT
	// e recebe request.WorkInfoUpdateRequest.
	// Retorna o domínio atualizado.
	UpdateWorkInfoServices(
		userId string,
		updateRequest request.WorkInfoUpdateRequest, // Alterado para request.WorkInfoUpdateRequest
	) (model.WorkInfoDomainInterface, *rest_err.RestErr) // Alterado para retornar domain e erro
}
