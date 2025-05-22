package service

import (
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/logger"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model"
	"go.uber.org/zap"
)

func (wd *workInfoDomainService) CreateWorkInfoServices(
	workInfoDomain model.WorkInfoDomainInterface,
) (model.WorkInfoDomainInterface, *rest_err.RestErr) {
	logger.Info("Init CreateWorkInfoServices", zap.String("journey", "createWorkInfo")) // Nome da função corrigido no log

	// Validação 1: Verificar se o usuário para o qual o WorkInfo está sendo criado existe
	targetUserID := workInfoDomain.GetUserId()
	if targetUserID == "" { // Adicionado para garantir que o UserID não seja vazio
		logger.Error("Target UserID for WorkInfo cannot be empty", nil, zap.String("journey", "createWorkInfo"))
		return nil, rest_err.NewBadRequestError("Target UserID for WorkInfo cannot be empty")
	}
	_, errUser := wd.userDomainService.FindUserByIDServices(targetUserID)
	if errUser != nil {
		logger.Error("Target user for WorkInfo not found", errUser,
			zap.String("journey", "createWorkInfo"),
			zap.String("targetUserID", targetUserID))
		// Retornar um erro mais específico se o usuário não foi encontrado
		if errUser.Code == 404 { // Código para NotFoundError
			return nil, rest_err.NewBadRequestError("Target user for WorkInfo does not exist")
		}
		return nil, errUser // Outros erros do FindUserByIDServices
	}

	// Validação 2: Verificar se o SuperiorID (se fornecido e não vazio) corresponde a um usuário existente
	superiorID := workInfoDomain.GetSuperiorID()
	if superiorID != "" { // Apenas validar se um superiorID foi fornecido
		_, errSuperior := wd.userDomainService.FindUserByIDServices(superiorID)
		if errSuperior != nil {
			logger.Error("Superior user not found", errSuperior,
				zap.String("journey", "createWorkInfo"),
				zap.String("superiorID", superiorID))
			// Retornar um erro mais específico se o superior não foi encontrado
			if errSuperior.Code == 404 { // Código para NotFoundError
				return nil, rest_err.NewBadRequestError("Superior user specified in WorkInfo does not exist")
			}
			return nil, errSuperior // Outros erros do FindUserByIDServices
		}
	}

	workInfo, err := wd.workInfoRepository.CreateWorkInfo(workInfoDomain)
	if err != nil {
		logger.Error("Error trying to create work info in repository", err, zap.String("journey", "createWorkInfo")) // Mensagem de log melhorada
		return nil, err
	}

	logger.Info("CreateWorkInfoServices executed successfully", // Nome da função corrigido no log
		zap.String("userId", workInfo.GetUserId()),
		zap.String("journey", "createWorkInfo"))

	return workInfo, nil
}
