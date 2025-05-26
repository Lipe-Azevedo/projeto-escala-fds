package workinfo

import (
	"fmt"

	"github.com/Lipe-Azevedo/escala-fds/src/configuration/logger"
	"github.com/Lipe-Azevedo/escala-fds/src/configuration/rest_err"

	// IMPORT ATUALIZADO: Agora importa do subpacote 'domain'
	"github.com/Lipe-Azevedo/escala-fds/src/model/domain"
	"go.uber.org/zap"
)

func (wd *workInfoDomainService) CreateWorkInfoServices(
	workInfoReqDomain domain.WorkInfoDomainInterface, // <<< USA domain.WorkInfoDomainInterface
) (domain.WorkInfoDomainInterface, *rest_err.RestErr) { // <<< USA domain.WorkInfoDomainInterface
	logger.Info("Init CreateWorkInfoServices",
		zap.String("journey", "createWorkInfo"),
		zap.String("targetUserId", workInfoReqDomain.GetUserId()))

	targetUserID := workInfoReqDomain.GetUserId()
	if targetUserID == "" {
		logger.Error("Target UserID for WorkInfo cannot be empty", nil, zap.String("journey", "createWorkInfo"))
		return nil, rest_err.NewBadRequestError("Target UserID for WorkInfo cannot be empty")
	}

	targetUser, errUser := wd.userDomainService.FindUserByIDServices(targetUserID)
	if errUser != nil {
		logger.Error("Target user for WorkInfo creation not found by service", errUser,
			zap.String("journey", "createWorkInfo"),
			zap.String("targetUserID", targetUserID))
		if errUser.Code == 404 {
			return nil, rest_err.NewBadRequestError(fmt.Sprintf("Target user (ID: %s) for WorkInfo does not exist", targetUserID))
		}
		return nil, errUser
	}

	if targetUser.GetUserType() != domain.UserTypeCollaborator { // <<< USA domain.UserTypeCollaborator
		logger.Warn("Attempt to create WorkInfo for a non-collaborator user",
			zap.String("journey", "createWorkInfo"),
			zap.String("targetUserID", targetUserID),
			zap.String("userType", string(targetUser.GetUserType())))
		return nil, rest_err.NewBadRequestError(fmt.Sprintf("WorkInfo can only be created for 'colaborador' users. User ID: %s is a '%s'", targetUserID, targetUser.GetUserType()))
	}

	superiorID := workInfoReqDomain.GetSuperiorID()
	if superiorID != "" {
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
			if errSuperior.Code == 404 {
				return nil, rest_err.NewBadRequestError(fmt.Sprintf("Superior user (ID: %s) specified in WorkInfo does not exist", superiorID))
			}
			return nil, errSuperior
		}
	}

	existingWorkInfo, _ := wd.workInfoRepository.FindWorkInfoByUserId(targetUserID)
	if existingWorkInfo != nil {
		logger.Warn("WorkInfo already exists for this user",
			zap.String("journey", "createWorkInfo"),
			zap.String("targetUserID", targetUserID))
		return nil, rest_err.NewConflictError(fmt.Sprintf("WorkInfo already exists for user ID: %s", targetUserID))
	}

	workInfo, errRepo := wd.workInfoRepository.CreateWorkInfo(workInfoReqDomain)
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
