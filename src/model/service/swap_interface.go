package service

import (
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model/repository"
)

func NewSwapDomainService(
	repository repository.SwapRepository,
) SwapDomainService {
	return &swapDomainService{
		repository,
	}
}

type swapDomainService struct {
	repository repository.SwapRepository
}

type SwapDomainService interface {
	CreateSwapServices(
		swapDomain model.SwapDomainInterface,
	) (model.SwapDomainInterface, *rest_err.RestErr)

	FindSwapByIDServices(
		id string,
	) (model.SwapDomainInterface, *rest_err.RestErr)

	UpdateSwapServices(
		id string,
		swapDomain model.SwapDomainInterface,
	) *rest_err.RestErr
}
