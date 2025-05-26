package user

import (
	"github.com/Lipe-Azevedo/escala-fds/src/configuration/logger"
	"github.com/Lipe-Azevedo/escala-fds/src/configuration/rest_err"

	// IMPORT ATUALIZADO: Agora importa do subpacote 'domain'
	"github.com/Lipe-Azevedo/escala-fds/src/model/domain"
	"go.uber.org/zap"
)

func (ud *userDomainService) UpdateUserServices(
	userId string,
	userUpdateRequestDomain domain.UserDomainInterface, // <<< USA domain.UserDomainInterface
) *rest_err.RestErr {

	logger.Info(
		"Init UpdateUserServices",
		zap.String("journey", "updateUser"),
		zap.String("userId", userId))

	if userUpdateRequestDomain.GetPassword() != "" {
		// A interface domain.UserDomainInterface deve ter EncryptPassword.
		// O objeto userUpdateRequestDomain é uma instância que implementa essa interface.
		if err := userUpdateRequestDomain.EncryptPassword(); err != nil {
			logger.Error("Error encrypting password during user update", err,
				zap.String("journey", "updateUser"),
				zap.String("userId", userId))
			return rest_err.NewInternalServerError("Error processing password for update")
		}
	}

	repoErr := ud.userRepository.UpdateUser(userId, userUpdateRequestDomain)
	if repoErr != nil {
		logger.Error(
			"Error trying to call repository to update user",
			repoErr,
			zap.String("journey", "updateUser"),
			zap.String("userId", userId),
		)
		return repoErr
	}

	logger.Info(
		"UpdateUserServices executed successfully",
		zap.String("userId", userId),
		zap.String("journey", "updateUser"),
	)
	return nil
}
