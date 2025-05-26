package user

import (
	"github.com/Lipe-Azevedo/escala-fds/src/configuration/logger"
	"github.com/Lipe-Azevedo/escala-fds/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/escala-fds/src/model/domain" // CORRETO
	"go.uber.org/zap"
)

func (ud *userDomainService) FindUserByIDServices(
	id string,
) (domain.UserDomainInterface, *rest_err.RestErr) {
	logger.Info(
		"Init FindUserByIDServices",
		zap.String("journey", "findUserByID"),
		zap.String("userIdToFind", id))

	userDomainVal, err := ud.userRepository.FindUserByID(id) // Espera-se que retorne domain.UserDomainInterface
	if err != nil {
		logger.Error("Error calling repository to find user by ID", err,
			zap.String("journey", "findUserByID"),
			zap.String("userId", id))
		return nil, err
	}

	logger.Info("FindUserByIDServices executed successfully",
		zap.String("journey", "findUserByID"),
		zap.String("userId", id))
	return userDomainVal, nil
}

func (ud *userDomainService) FindUserByEmailServices(
	email string,
) (domain.UserDomainInterface, *rest_err.RestErr) {
	logger.Info(
		"Init FindUserByEmailServices",
		zap.String("journey", "findUserByEmail"),
		zap.String("emailToFind", email))

	userDomainVal, err := ud.userRepository.FindUserByEmail(email) // Espera-se que retorne domain.UserDomainInterface
	if err != nil {
		logger.Error("Error calling repository to find user by email", err,
			zap.String("journey", "findUserByEmail"),
			zap.String("email", email))
		return nil, err
	}

	logger.Info("FindUserByEmailServices executed successfully",
		zap.String("journey", "findUserByEmail"),
		zap.String("email", email))
	return userDomainVal, nil
}

func (ud *userDomainService) FindAllUsersServices() ([]domain.UserDomainInterface, *rest_err.RestErr) { // CORRETO: Assinatura
	logger.Info(
		"Init FindAllUsersServices",
		zap.String("journey", "findAllUsers"))

	// ud.userRepository.FindAllUsers() deve retornar []domain.UserDomainInterface
	userDomains, err := ud.userRepository.FindAllUsers()
	if err != nil {
		logger.Error("Error calling repository for FindAllUsers", err,
			zap.String("journey", "findAllUsers"))
		return nil, err
	}

	logger.Info("FindAllUsersServices executed successfully",
		zap.Int("count", len(userDomains)),
		zap.String("journey", "findAllUsers"))
	return userDomains, nil // Esta linha deve ser válida agora
}
