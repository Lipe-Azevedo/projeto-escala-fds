package workinfo

import (
	"fmt"

	"github.com/Lipe-Azevedo/escala-fds/src/configuration/logger"
	"github.com/Lipe-Azevedo/escala-fds/src/configuration/rest_err"
	workinfo_request_dto "github.com/Lipe-Azevedo/escala-fds/src/controller/workinfo/request"

	// IMPORT ATUALIZADO: Agora importa do subpacote 'domain'
	"github.com/Lipe-Azevedo/escala-fds/src/model/domain"
	"go.uber.org/zap"
)

func (wd *workInfoDomainService) UpdateWorkInfoServices(
	userId string,
	updateRequest workinfo_request_dto.WorkInfoUpdateRequest,
) (domain.WorkInfoDomainInterface, *rest_err.RestErr) { // <<< USA domain.WorkInfoDomainInterface
	logger.Info("Init UpdateWorkInfoServices",
		zap.String("journey", "updateWorkInfo"),
		zap.String("userID", userId))

	if userId == "" {
		logger.Error("UserID for WorkInfo update cannot be empty", nil, zap.String("journey", "updateWorkInfo"))
		return nil, rest_err.NewBadRequestError("User ID for WorkInfo update cannot be empty")
	}

	targetUser, errUser := wd.userDomainService.FindUserByIDServices(userId)
	if errUser != nil {
		logger.Error("User for WorkInfo update not found by service", errUser,
			zap.String("journey", "updateWorkInfo"),
			zap.String("userID", userId))
		if errUser.Code == 404 {
			return nil, rest_err.NewBadRequestError(fmt.Sprintf("User (ID: %s) for WorkInfo update does not exist", userId))
		}
		return nil, errUser
	}
	if targetUser.GetUserType() != domain.UserTypeCollaborator { // <<< USA domain.UserTypeCollaborator
		logger.Warn("Attempt to update WorkInfo for a non-collaborator user",
			zap.String("journey", "updateWorkInfo"),
			zap.String("targetUserID", userId),
			zap.String("userType", string(targetUser.GetUserType())))
		return nil, rest_err.NewForbiddenError(fmt.Sprintf("WorkInfo can only be updated for 'colaborador' users. User ID: %s is a '%s'", userId, targetUser.GetUserType()))
	}

	existingWorkInfoDomain, err := wd.workInfoRepository.FindWorkInfoByUserId(userId)
	if err != nil {
		logger.Error("WorkInfo to update not found by service", err,
			zap.String("journey", "updateWorkInfo"),
			zap.String("userID", userId))
		return nil, err
	}

	fieldsUpdated := false
	if updateRequest.Team != nil {
		newTeam := domain.Team(*updateRequest.Team) // <<< USA domain.Team
		if newTeam != existingWorkInfoDomain.GetTeam() {
			existingWorkInfoDomain.SetTeam(newTeam)
			fieldsUpdated = true
		}
	}
	if updateRequest.Position != nil {
		if *updateRequest.Position != existingWorkInfoDomain.GetPosition() {
			existingWorkInfoDomain.SetPosition(*updateRequest.Position)
			fieldsUpdated = true
		}
	}
	if updateRequest.DefaultShift != nil {
		newShift := domain.Shift(*updateRequest.DefaultShift) // <<< USA domain.Shift
		if newShift != existingWorkInfoDomain.GetDefaultShift() {
			existingWorkInfoDomain.SetDefaultShift(newShift)
			fieldsUpdated = true
		}
	}
	if updateRequest.WeekdayOff != nil {
		newWeekdayOff := domain.Weekday(*updateRequest.WeekdayOff) // <<< USA domain.Weekday
		if newWeekdayOff != existingWorkInfoDomain.GetWeekdayOff() {
			existingWorkInfoDomain.SetWeekdayOff(newWeekdayOff)
			fieldsUpdated = true
		}
	}
	if updateRequest.WeekendDayOff != nil {
		newWeekendDayOff := domain.WeekendDayOff(*updateRequest.WeekendDayOff) // <<< USA domain.WeekendDayOff
		if newWeekendDayOff != existingWorkInfoDomain.GetWeekendDayOff() {
			existingWorkInfoDomain.SetWeekendDayOff(newWeekendDayOff)
			fieldsUpdated = true
		}
	}
	if updateRequest.SuperiorID != nil {
		newSuperiorID := *updateRequest.SuperiorID
		if newSuperiorID != "" {
			if newSuperiorID == userId {
				logger.Error("New SuperiorID cannot be the same as TargetUserID for WorkInfo update", nil,
					zap.String("journey", "updateWorkInfo"),
					zap.String("targetUserID", userId))
				return nil, rest_err.NewBadRequestError("New Superior can't be the same as the user")
			}
			_, errSuperior := wd.userDomainService.FindUserByIDServices(newSuperiorID)
			if errSuperior != nil {
				logger.Error("New Superior user (for update) not found by service", errSuperior,
					zap.String("journey", "updateWorkInfo"),
					zap.String("newSuperiorID", newSuperiorID))
				if errSuperior.Code == 404 {
					return nil, rest_err.NewBadRequestError(fmt.Sprintf("New Superior user (ID: %s) specified in WorkInfo update does not exist", newSuperiorID))
				}
				return nil, errSuperior
			}
		}
		if newSuperiorID != existingWorkInfoDomain.GetSuperiorID() {
			existingWorkInfoDomain.SetSuperiorID(newSuperiorID)
			fieldsUpdated = true
		}
	}

	if !fieldsUpdated {
		logger.Info("No actual changes detected for WorkInfo update.",
			zap.String("journey", "updateWorkInfo"),
			zap.String("userID", userId))
		return existingWorkInfoDomain, nil
	}

	repoErr := wd.workInfoRepository.UpdateWorkInfo(userId, existingWorkInfoDomain)
	if repoErr != nil {
		logger.Error("Error calling repository to update WorkInfo", repoErr,
			zap.String("journey", "updateWorkInfo"),
			zap.String("userID", userId))
		return nil, repoErr
	}

	logger.Info("UpdateWorkInfoServices executed successfully.",
		zap.String("userId", userId),
		zap.String("journey", "updateWorkInfo"),
	)
	return existingWorkInfoDomain, nil
}
