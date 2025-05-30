package user

import (
	"log" // Para o panic na inicialização se a chave estiver ausente

	"github.com/Lipe-Azevedo/escala-fds/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/escala-fds/src/model/domain"
	repository_user "github.com/Lipe-Azevedo/escala-fds/src/model/repository/user"
)

type UserDomainService interface {
	CreateUserServices(userDomainReq domain.UserDomainInterface) (
		domain.UserDomainInterface, *rest_err.RestErr)
	FindUserByEmailServices(email string) (domain.UserDomainInterface, *rest_err.RestErr)
	FindUserByIDServices(id string) (domain.UserDomainInterface, *rest_err.RestErr)
	UpdateUserServices(userId string, userUpdateRequestDomain domain.UserDomainInterface) *rest_err.RestErr
	DeleteUserServices(userId string) *rest_err.RestErr
	FindAllUsersServices() ([]domain.UserDomainInterface, *rest_err.RestErr)
	LoginUserServices(email, password string) (string, domain.UserDomainInterface, *rest_err.RestErr)
}

type userDomainService struct {
	userRepository repository_user.UserRepository
	jwtSecret      string // NOVO CAMPO para armazenar a chave JWT
}

func NewUserDomainService(
	userRepository repository_user.UserRepository,
	jwtSecret string, // NOVO PARÂMETRO
) UserDomainService {
	// É crucial que a chave JWT não esteja vazia aqui.
	// A verificação principal disso já está em main.go antes de chamar initDependencies.
	if jwtSecret == "" {
		// Este panic interromperia a inicialização se a chave não fosse passada.
		// Isso é uma salvaguarda, mas o main.go já deve garantir isso.
		log.Fatal("CRITICAL: jwtSecret cannot be empty when creating UserDomainService.")
	}
	return &userDomainService{
		userRepository: userRepository,
		jwtSecret:      jwtSecret, // ARMAZENA A CHAVE
	}
}
