package user

import (
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/rest_err"
	// IMPORT ATUALIZADO: Agora importa do subpacote 'domain'
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model/domain"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model/repository/user"
)

// UserDomainService define a interface para os serviços de domínio de usuário.
type UserDomainService interface {
	CreateUserServices(userDomainReq domain.UserDomainInterface) ( // <<< USA domain.UserDomainInterface
		domain.UserDomainInterface, *rest_err.RestErr) // <<< USA domain.UserDomainInterface

	FindUserByEmailServices(
		email string,
	) (domain.UserDomainInterface, *rest_err.RestErr) // <<< USA domain.UserDomainInterface

	FindUserByIDServices(
		id string,
	) (domain.UserDomainInterface, *rest_err.RestErr) // <<< USA domain.UserDomainInterface

	UpdateUserServices(userId string, userUpdateRequestDomain domain.UserDomainInterface) *rest_err.RestErr // <<< USA domain.UserDomainInterface

	DeleteUserServices(userId string) *rest_err.RestErr

	FindAllUsersServices() ([]domain.UserDomainInterface, *rest_err.RestErr) // <<< USA []domain.UserDomainInterface
}

// userDomainService é a implementação da interface UserDomainService.
type userDomainService struct {
	userRepository user.UserRepository
}

// NewUserDomainService cria uma nova instância de UserDomainService.
func NewUserDomainService(
	userRepository user.UserRepository,
) UserDomainService {
	return &userDomainService{userRepository}
}
