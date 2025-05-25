package workinfo

import (
	"fmt"

	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/logger"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model"
	"go.uber.org/zap"
)

func (wd *workInfoDomainService) CreateWorkInfoServices(
	workInfoDomain model.WorkInfoDomainInterface,
) (model.WorkInfoDomainInterface, *rest_err.RestErr) {
	logger.Info("Init CreateWorkInfoServices",
		zap.String("journey", "createWorkInfo"),
		zap.String("targetUserId", workInfoDomain.GetUserId()))

	targetUserID := workInfoDomain.GetUserId()
	if targetUserID == "" {
		logger.Error("Target UserID for WorkInfo cannot be empty", nil, zap.String("journey", "createWorkInfo"))
		return nil, rest_err.NewBadRequestError("Target UserID for WorkInfo cannot be empty")
	}

	// Validação 1: Verificar se o usuário para o qual o WorkInfo está sendo criado existe.
	// wd.userDomainService é do tipo service_user.UserDomainService (já reorganizado)
	targetUser, errUser := wd.userDomainService.FindUserByIDServices(targetUserID)
	if errUser != nil {
		logger.Error("Target user for WorkInfo creation not found by service", errUser,
			zap.String("journey", "createWorkInfo"),
			zap.String("targetUserID", targetUserID))
		if errUser.Code == 404 { // http.StatusNotFound
			return nil, rest_err.NewBadRequestError(fmt.Sprintf("Target user (ID: %s) for WorkInfo does not exist", targetUserID))
		}
		return nil, errUser
	}
	// Adicional: Verificar se o usuário é um 'colaborador'
	if targetUser.GetUserType() != model.UserTypeCollaborator {
		logger.Warn("Attempt to create WorkInfo for a non-collaborator user",
			zap.String("journey", "createWorkInfo"),
			zap.String("targetUserID", targetUserID),
			zap.String("userType", string(targetUser.GetUserType())))
		return nil, rest_err.NewBadRequestError(fmt.Sprintf("WorkInfo can only be created for 'colaborador' users. User ID: %s is a '%s'", targetUserID, targetUser.GetUserType()))
	}

	// Validação 2: Verificar se o SuperiorID (se fornecido e não vazio) corresponde a um usuário existente.
	superiorID := workInfoDomain.GetSuperiorID()
	if superiorID != "" {
		// Opcional: Validar se o superiorID é diferente do targetUserID
		if superiorID == targetUserID {
			logger.Error("SuperiorID cannot be the same as TargetUserID for WorkInfo", nil,
				zap.String("journey", "createWorkInfo"),
				zap.String("targetUserID", targetUserID))
			return nil, rest_err.NewBadRequestError("Superior can't be the same as the user")
		}
		_, errSuperior := wd.userDomainService.FindUserByIDServices(superiorID)
		if errSuperior != nil {
			logger.Error("Superior user specified in WorkInfo not found by service", errSuperior,
				zap.String("journey", "createWorkInfo"),
				zap.String("superiorID", superiorID))
			if errSuperior.Code == 404 { // http.StatusNotFound
				return nil, rest_err.NewBadRequestError(fmt.Sprintf("Superior user (ID: %s) specified in WorkInfo does not exist", superiorID))
			}
			return nil, errSuperior
		}
	}

	// Validação 3: Verificar se já existe WorkInfo para este usuário
	// (O repositório já faz isso com o ConflictError, mas uma verificação aqui pode ser mais explícita)
	existingWorkInfo, _ := wd.workInfoRepository.FindWorkInfoByUserId(targetUserID)
	if existingWorkInfo != nil {
		logger.Warn("WorkInfo already exists for this user",
			zap.String("journey", "createWorkInfo"),
			zap.String("targetUserID", targetUserID))
		return nil, rest_err.NewConflictError(fmt.Sprintf("WorkInfo already exists for user ID: %s", targetUserID))
	}

	// wd.workInfoRepository é do tipo repository_workinfo.WorkInfoRepository (já reorganizado)
	workInfo, errRepo := wd.workInfoRepository.CreateWorkInfo(workInfoDomain)
	if errRepo != nil {
		logger.Error("Error calling repository to create work info", errRepo,
			zap.String("journey", "createWorkInfo"))
		return nil, errRepo
	}

	logger.Info("CreateWorkInfoServices executed successfully",
		zap.String("userId", workInfo.GetUserId()),
		zap.String("journey", "createWorkInfo"))

	return workInfo, nil
}
