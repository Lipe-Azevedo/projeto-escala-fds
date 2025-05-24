package user

import (
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/logger"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) FindUserByIDServices(
	id string,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info(
		"Init FindUserByIDServices",
		zap.String("journey", "findUserByID"),
		zap.String("userIdToFind", id))

	userDomain, err := ud.userRepository.FindUserByID(id)
	if err != nil {
		// O repositório já loga o erro específico do MongoDB.
		// O serviço loga que a chamada ao repositório falhou.
		logger.Error("Error calling repository to find user by ID", err,
			zap.String("journey", "findUserByID"),
			zap.String("userId", id))
		return nil, err
	}

	logger.Info("FindUserByIDServices executed successfully",
		zap.String("journey", "findUserByID"),
		zap.String("userId", id))
	return userDomain, nil
}

func (ud *userDomainService) FindUserByEmailServices(
	email string,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info(
		"Init FindUserByEmailServices",
		zap.String("journey", "findUserByEmail"),
		zap.String("emailToFind", email))

	userDomain, err := ud.userRepository.FindUserByEmail(email)
	if err != nil {
		logger.Error("Error calling repository to find user by email", err,
			zap.String("journey", "findUserByEmail"),
			zap.String("email", email))
		return nil, err
	}

	logger.Info("FindUserByEmailServices executed successfully",
		zap.String("journey", "findUserByEmail"),
		zap.String("email", email))
	return userDomain, nil
}

func (ud *userDomainService) FindAllUsersServices() ([]model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info(
		"Init FindAllUsersServices",
		zap.String("journey", "findAllUsers"))

	userDomains, err := ud.userRepository.FindAllUsers()
	if err != nil {
		logger.Error("Error calling repository for FindAllUsers", err,
			zap.String("journey", "findAllUsers"))
		return nil, err
	}

	logger.Info("FindAllUsersServices executed successfully",
		zap.Int("count", len(userDomains)),
		zap.String("journey", "findAllUsers"))
	return userDomains, nil
}
