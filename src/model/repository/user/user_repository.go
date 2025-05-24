package user

import (
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	// MONGODB_USERS_COLLECTION_ENV_KEY armazena o nome da variável de ambiente que contém o nome da coleção de usuários.
	MONGODB_USERS_COLLECTION_ENV_KEY = "MONGODB_USERS_COLLECTION"
)

// UserRepository define a interface para o repositório de usuários.
// As implementações dos métodos estarão em arquivos separados neste pacote.
type UserRepository interface {
	CreateUser(
		userDomain model.UserDomainInterface,
	) (model.UserDomainInterface, *rest_err.RestErr)

	UpdateUser(
		userId string,
		userDomain model.UserDomainInterface,
	) *rest_err.RestErr

	DeleteUser(
		userId string,
	) *rest_err.RestErr

	FindUserByEmail(
		email string,
	) (model.UserDomainInterface, *rest_err.RestErr)

	FindUserByID(
		id string,
	) (model.UserDomainInterface, *rest_err.RestErr)

	FindAllUsers() ([]model.UserDomainInterface, *rest_err.RestErr)
}

// userRepository é a implementação da interface UserRepository.
// O campo databaseConnection é usado pelos métodos definidos em outros arquivos deste pacote.
type userRepository struct {
	databaseConnection *mongo.Database
}

// NewUserRepository cria uma nova instância de UserRepository.
func NewUserRepository(
	database *mongo.Database,
) UserRepository {
	return &userRepository{
		databaseConnection: database,
	}
}
