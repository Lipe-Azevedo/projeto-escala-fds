package user

import (
	"github.com/Lipe-Azevedo/escala-fds/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/escala-fds/src/model/domain"
	repository_user "github.com/Lipe-Azevedo/escala-fds/src/model/repository/user"
	// service_workinfo NÃO É MAIS IMPORTADO AQUI
)

type UserDomainService interface {
	CreateUserServices(userDomainReq domain.UserDomainInterface) (
		domain.UserDomainInterface, *rest_err.RestErr)
	FindUserByEmailServices(email string) (domain.UserDomainInterface, *rest_err.RestErr)
	FindUserByIDServices(id string) (domain.UserDomainInterface, *rest_err.RestErr) // Retorna apenas UserDomain
	UpdateUserServices(userId string, userUpdateRequestDomain domain.UserDomainInterface) *rest_err.RestErr
	DeleteUserServices(userId string) *rest_err.RestErr
	FindAllUsersServices() ([]domain.UserDomainInterface, *rest_err.RestErr)
	LoginUserServices(email, password string) (string, domain.UserDomainInterface, *rest_err.RestErr) // NOVO MÉTODO (retorna token, userDomain, error)
}

type userDomainService struct {
	userRepository repository_user.UserRepository
	// workInfoService  service_workinfo.WorkInfoDomainService // REMOVIDA DEPENDÊNCIA
}

func NewUserDomainService(
	userRepository repository_user.UserRepository,
	// workInfoService service_workinfo.WorkInfoDomainService, // REMOVIDA DEPENDÊNCIA
) UserDomainService {
	return &userDomainService{
		userRepository: userRepository,
		// workInfoService: workInfoService, // REMOVIDA INJEÇÃO
	}
}
