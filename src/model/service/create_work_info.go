package service

import (
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/logger"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateWorkInfoServices(
	workInfoDomain model.WorkInfoDomainInterface,
) (model.WorkInfoDomainInterface, *rest_err.RestErr) {
	logger.Info("Init CreateWorkInfo service", zap.String("journey", "createWorkInfo"))

	// Verificar se o usuário existe e é do tipo colaborador
	user, err := ud.userRepository.FindUserByID(workInfoDomain.GetUserId())
	if err != nil {
		return nil, err
	}

	if user.GetUserType() != model.UserTypeCollaborator {
		return nil, rest_err.NewBadRequestError("Only collaborators can have work info")
	}

	// Verificar se o superior existe
	if _, err := ud.userRepository.FindUserByID(workInfoDomain.GetSuperiorID()); err != nil {
		return nil, rest_err.NewBadRequestError("Superior not found")
	}

	workInfo, err := ud.userRepository.CreateWorkInfo(workInfoDomain)
	if err != nil {
		logger.Error("Error trying to create work info", err, zap.String("journey", "createWorkInfo"))
		return nil, err
	}

	logger.Info("CreateWorkInfo service executed successfully",
		zap.String("userId", workInfo.GetUserId()),
		zap.String("journey", "createWorkInfo"))

	return workInfo, nil
}
