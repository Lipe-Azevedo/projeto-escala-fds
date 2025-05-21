package service

import (
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model/repository"
)

func NewShiftSwapDomainService(
	repository repository.ShiftSwapRepository,
) ShiftSwapDomainService {
	return &shiftSwapDomainService{
		repository,
	}
}

type shiftSwapDomainService struct {
	repository repository.ShiftSwapRepository
}

type ShiftSwapDomainService interface {
	CreateShiftSwapServices(
		shiftSwapDomain model.ShiftSwapDomainInterface,
	) (model.ShiftSwapDomainInterface, *rest_err.RestErr)

	FindShiftSwapByIDServices(
		id string,
	) (model.ShiftSwapDomainInterface, *rest_err.RestErr)

	UpdateShiftSwapServices(
		id string,
		shiftSwapDomain model.ShiftSwapDomainInterface,
	) *rest_err.RestErr
}
