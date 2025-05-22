package service

import (
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/logger"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/controller/model/request"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model"
	"go.uber.org/zap"
)

func (wd *workInfoDomainService) UpdateWorkInfoServices(
	userId string,
	updateRequest request.WorkInfoUpdateRequest,
) (model.WorkInfoDomainInterface, *rest_err.RestErr) {
	logger.Info("Init UpdateWorkInfoServices (handling partial update via PUT)",
		zap.String("journey", "updateWorkInfo"),
		zap.String("userID", userId))

	_, errUser := wd.userDomainService.FindUserByIDServices(userId)
	if errUser != nil {
		logger.Error("User for WorkInfo update not found", errUser,
			zap.String("journey", "updateWorkInfo"),
			zap.String("userID", userId))
		if errUser.Code == 404 {
			return nil, rest_err.NewBadRequestError("User for WorkInfo update does not exist")
		}
		return nil, errUser
	}

	// Esta é a chamada crucial que está resultando em "Work info not found"
	existingWorkInfoDomain, err := wd.workInfoRepository.FindWorkInfoByUserId(userId)
	if err != nil {
		// O log de erro já acontece no repositório se não encontrado.
		// O serviço apenas repassa o erro.
		// Se err for *rest_err.RestErr com Code 404, ele será retornado como está.
		logger.Error("WorkInfo not found by service for update, or other repository error", err,
			zap.String("journey", "updateWorkInfo"),
			zap.String("userID", userId))
		return nil, err
	}

	fieldsUpdated := false

	if updateRequest.Team != nil {
		if model.Team(*updateRequest.Team) != existingWorkInfoDomain.GetTeam() {
			existingWorkInfoDomain.SetTeam(model.Team(*updateRequest.Team))
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
		if model.Shift(*updateRequest.DefaultShift) != existingWorkInfoDomain.GetDefaultShift() {
			existingWorkInfoDomain.SetDefaultShift(model.Shift(*updateRequest.DefaultShift))
			fieldsUpdated = true
		}
	}
	if updateRequest.WeekdayOff != nil {
		if model.Weekday(*updateRequest.WeekdayOff) != existingWorkInfoDomain.GetWeekdayOff() {
			existingWorkInfoDomain.SetWeekdayOff(model.Weekday(*updateRequest.WeekdayOff))
			fieldsUpdated = true
		}
	}
	if updateRequest.WeekendDayOff != nil {
		if model.WeekendDayOff(*updateRequest.WeekendDayOff) != existingWorkInfoDomain.GetWeekendDayOff() {
			existingWorkInfoDomain.SetWeekendDayOff(model.WeekendDayOff(*updateRequest.WeekendDayOff))
			fieldsUpdated = true
		}
	}
	if updateRequest.SuperiorID != nil {
		if *updateRequest.SuperiorID != "" {
			_, errSuperior := wd.userDomainService.FindUserByIDServices(*updateRequest.SuperiorID)
			if errSuperior != nil {
				logger.Error("New Superior user (for update) not found", errSuperior,
					zap.String("journey", "updateWorkInfo"),
					zap.String("newSuperiorID", *updateRequest.SuperiorID))
				if errSuperior.Code == 404 {
					return nil, rest_err.NewBadRequestError("New Superior user specified in WorkInfo update does not exist")
				}
				return nil, errSuperior
			}
		}
		if *updateRequest.SuperiorID != existingWorkInfoDomain.GetSuperiorID() {
			existingWorkInfoDomain.SetSuperiorID(*updateRequest.SuperiorID)
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
		logger.Error("Error updating WorkInfo in repository", repoErr,
			zap.String("journey", "updateWorkInfo"),
			zap.String("userID", userId))
		return nil, repoErr
	}

	logger.Info("UpdateWorkInfoServices (handling partial update via PUT) executed successfully.",
		zap.String("userId", userId),
		zap.String("journey", "updateWorkInfo"),
	)
	return existingWorkInfoDomain, nil
}
