package user

import (
	"github.com/Lipe-Azevedo/escala-fds/src/configuration/logger"
	"github.com/Lipe-Azevedo/escala-fds/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/escala-fds/src/model/domain"
	"go.uber.org/zap"
)

func (ud *userDomainService) FindUserByIDServices(
	id string,
) (domain.UserDomainInterface, *rest_err.RestErr) {
	logger.Info(
		"Init FindUserByIDServices (fetches only user data)", // Log atualizado
		zap.String("journey", "findUserByID"),
		zap.String("userIdToFind", id))

	userDomainVal, err := ud.userRepository.FindUserByID(id)
	if err != nil {
		logger.Error("Error calling repository to find user by ID", err,
			zap.String("journey", "findUserByID"),
			zap.String("userId", id))
		return nil, err
	}

	// A lógica de buscar WorkInfo foi removida daqui.
	// O controller fará isso.

	logger.Info("FindUserByIDServices executed successfully (user data only)", // Log atualizado
		zap.String("journey", "findUserByID"),
		zap.String("userId", id))
	return userDomainVal, nil
}

func (ud *userDomainService) FindUserByEmailServices(
	email string,
) (domain.UserDomainInterface, *rest_err.RestErr) {
	logger.Info(
		"Init FindUserByEmailServices (fetches only user data)", // Log atualizado
		zap.String("journey", "findUserByEmail"),
		zap.String("emailToFind", email))

	userDomainVal, err := ud.userRepository.FindUserByEmail(email)
	if err != nil {
		logger.Error("Error calling repository to find user by email", err,
			zap.String("journey", "findUserByEmail"),
			zap.String("email", email))
		return nil, err
	}

	// A lógica de buscar WorkInfo opcional foi removida daqui para consistência.
	// O controller, se precisar, fará a chamada adicional ao WorkInfoService.

	logger.Info("FindUserByEmailServices executed successfully (user data only)", // Log atualizado
		zap.String("journey", "findUserByEmail"),
		zap.String("email", email))
	return userDomainVal, nil
}

// FindAllUsersServices não busca WorkInfo para evitar N+1 queries.
func (ud *userDomainService) FindAllUsersServices() ([]domain.UserDomainInterface, *rest_err.RestErr) {
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
