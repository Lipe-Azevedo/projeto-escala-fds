package service

import (
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model/repository"
)

func NewWorkInfiDomainService(
	workInfoRepository repository.UserRepository,
) UserDomainService {
	return &userDomainService{userRepository}
}

type workInfoDomainService struct {
	userRepository repository.UserRepository
}

type orkInfonService interface {
	CreateUserServices(model.UserDomainInterface) (
		model.UserDomainInterface, *rest_err.RestErr)

	FindUserByEmailServices(
		email string,
	) (model.UserDomainInterface, *rest_err.RestErr)

	FindUserByIDServices(
		id string,
	) (model.UserDomainInterface, *rest_err.RestErr)

	UpdateUser(string, model.UserDomainInterface) *rest_err.RestErr

	DeleteUser(string) *rest_err.RestErr
}
