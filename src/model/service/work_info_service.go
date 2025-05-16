package service

import (
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/logger"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model/repository"
	"go.uber.org/zap"
)

type workInfoService struct {
	userRepository repository.UserRepository
}

func NewWorkInfoService(
	userRepository repository.UserRepository,
) *workInfoService {
	return &workInfoService{
		userRepository: userRepository,
	}
}

func (ws *workInfoService) UpdateWorkInfo(
	userId string,
	workInfo model.WorkInfoInterface,
) *rest_err.RestErr {
	logger.Info("Init updateWorkInfo service",
		zap.String("journey", "updateWorkInfo"))

	// Validação básica
	if err := workInfo.Validate(); err != nil {
		return err
	}

	// Verificar se o superior existe
	if _, err := ws.userRepository.FindUserByID(workInfo.GetSuperiorID()); err != nil {
		return rest_err.NewBadRequestError("Superior not found")
	}

	// Atualizar no banco de dados
	if err := ws.userRepository.UpdateWorkInfo(userId, workInfo); err != nil {
		logger.Error("Error updating work info",
			err,
			zap.String("journey", "updateWorkInfo"))
		return err
	}

	logger.Info("WorkInfo updated successfully",
		zap.String("userId", userId),
		zap.String("journey", "updateWorkInfo"))
	return nil
}
