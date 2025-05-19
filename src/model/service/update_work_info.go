package service

import (
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/logger"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) UpdateWorkInfoServices(
	userId string,
	workInfoDomain model.WorkInfoDomainInterface,
) *rest_err.RestErr {
	logger.Info("Init UpdateWorkInfo service", zap.String("journey", "updateWorkInfo"))

	// Verificar se o superior existe
	if workInfoDomain.GetSuperiorID() != "" {
		if _, err := ud.userRepository.FindUserByID(workInfoDomain.GetSuperiorID()); err != nil {
			return rest_err.NewBadRequestError("Superior not found")
		}
	}

	return ud.userRepository.UpdateWorkInfo(userId, workInfoDomain)
}
