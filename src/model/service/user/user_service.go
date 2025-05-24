package user

import (
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model/repository/user" // ATENÇÃO: Import ajustado para o novo pacote do repositório
)

// UserDomainService define a interface para os serviços de domínio de usuário.
// As implementações dos métodos estarão em arquivos separados neste pacote.
type UserDomainService interface {
	CreateUserServices(userDomain model.UserDomainInterface) (
		model.UserDomainInterface, *rest_err.RestErr)

	FindUserByEmailServices(
		email string,
	) (model.UserDomainInterface, *rest_err.RestErr)

	FindUserByIDServices(
		id string,
	) (model.UserDomainInterface, *rest_err.RestErr)

	UpdateUserServices(userId string, userDomain model.UserDomainInterface) *rest_err.RestErr // Renomeado de UpdateUser para UpdateUserServices

	DeleteUserServices(userId string) *rest_err.RestErr // Renomeado de DeleteUser para DeleteUserServices

	FindAllUsersServices() ([]model.UserDomainInterface, *rest_err.RestErr)
}

// userDomainService é a implementação da interface UserDomainService.
type userDomainService struct {
	userRepository user.UserRepository // Tipo ajustado para user.UserRepository
}

// NewUserDomainService cria uma nova instância de UserDomainService.
func NewUserDomainService(
	userRepository user.UserRepository, // Tipo ajustado para user.UserRepository
) UserDomainService {
	return &userDomainService{userRepository}
}
