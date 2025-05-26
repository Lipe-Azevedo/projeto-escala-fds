package user

import (
	"github.com/Lipe-Azevedo/escala-fds/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/escala-fds/src/model/domain" // CORRETO
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	MONGODB_USERS_COLLECTION_ENV_KEY = "MONGODB_USERS_COLLECTION"
)

type UserRepository interface {
	CreateUser(
		userDomainReq domain.UserDomainInterface, // CORRETO
	) (domain.UserDomainInterface, *rest_err.RestErr) // CORRETO

	UpdateUser(
		userId string,
		userDomainReq domain.UserDomainInterface, // CORRETO
	) *rest_err.RestErr

	DeleteUser(
		userId string,
	) *rest_err.RestErr

	FindUserByEmail(
		email string,
	) (domain.UserDomainInterface, *rest_err.RestErr) // CORRETO

	FindUserByID(
		id string,
	) (domain.UserDomainInterface, *rest_err.RestErr) // CORRETO

	FindAllUsers() ([]domain.UserDomainInterface, *rest_err.RestErr) // CORRETO E CRUCIAL PARA O ERRO ATUAL
}

type userRepository struct {
	databaseConnection *mongo.Database
}

func NewUserRepository(
	database *mongo.Database,
) UserRepository {
	return &userRepository{
		databaseConnection: database,
	}
}
