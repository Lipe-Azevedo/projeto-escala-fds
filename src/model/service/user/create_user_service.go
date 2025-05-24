package user

import (
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/logger"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUserServices(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info(
		"Init CreateUserServices",
		zap.String("journey", "createUser"),
		zap.String("email", userDomain.GetEmail()))

	// Verifica se o e-mail já está registrado
	// Esta chamada usará o método FindUserByEmailServices do mesmo serviço (ud).
	existingUser, _ := ud.FindUserByEmailServices(userDomain.GetEmail())
	if existingUser != nil {
		logger.Warn("Attempt to create user with already registered email",
			zap.String("email", userDomain.GetEmail()),
			zap.String("journey", "createUser"))
		return nil, rest_err.NewConflictError("Email is already registered in another account.") // Alterado para ConflictError
	}

	userDomain.EncryptPassword()

	userDomainRepository, err := ud.userRepository.CreateUser(userDomain)
	if err != nil {
		logger.Error("Error calling repository to create user", err,
			zap.String("journey", "createUser"))
		return nil, err
	}

	logger.Info("CreateUserServices executed successfully",
		zap.String("userId", userDomainRepository.GetID()),
		zap.String("journey", "createUser"))

	return userDomainRepository, nil
}
