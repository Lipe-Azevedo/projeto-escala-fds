package workinfo

import (
	"fmt"

	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/logger"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/controller/model/request" // Import temporário
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model"
	"go.uber.org/zap"
)

func (wd *workInfoDomainService) UpdateWorkInfoServices(
	userId string, // UserID do WorkInfo a ser atualizado
	updateRequest request.WorkInfoUpdateRequest, // DTO ainda do local antigo
) (model.WorkInfoDomainInterface, *rest_err.RestErr) {
	logger.Info("Init UpdateWorkInfoServices",
		zap.String("journey", "updateWorkInfo"),
		zap.String("userID", userId))

	if userId == "" {
		logger.Error("UserID for WorkInfo update cannot be empty", nil, zap.String("journey", "updateWorkInfo"))
		return nil, rest_err.NewBadRequestError("User ID for WorkInfo update cannot be empty")
	}

	// 1. Verificar se o usuário (dono do WorkInfo) existe e é um colaborador
	targetUser, errUser := wd.userDomainService.FindUserByIDServices(userId)
	if errUser != nil {
		logger.Error("User for WorkInfo update not found by service", errUser,
			zap.String("journey", "updateWorkInfo"),
			zap.String("userID", userId))
		if errUser.Code == 404 { // http.StatusNotFound
			return nil, rest_err.NewBadRequestError(fmt.Sprintf("User (ID: %s) for WorkInfo update does not exist", userId))
		}
		return nil, errUser
	}
	if targetUser.GetUserType() != model.UserTypeCollaborator {
		logger.Warn("Attempt to update WorkInfo for a non-collaborator user",
			zap.String("journey", "updateWorkInfo"),
			zap.String("targetUserID", userId),
			zap.String("userType", string(targetUser.GetUserType())))
		return nil, rest_err.NewForbiddenError(fmt.Sprintf("WorkInfo can only be updated for 'colaborador' users. User ID: %s is a '%s'", userId, targetUser.GetUserType()))
	}

	// 2. Buscar o WorkInfo existente
	existingWorkInfoDomain, err := wd.workInfoRepository.FindWorkInfoByUserId(userId)
	if err != nil {
		logger.Error("WorkInfo to update not found by service", err,
			zap.String("journey", "updateWorkInfo"),
			zap.String("userID", userId))
		// O repositório já retorna NotFoundError se for o caso
		return nil, err
	}

	// 3. Aplicar atualizações do request ao domain existente
	fieldsUpdated := false
	if updateRequest.Team != nil {
		newTeam := model.Team(*updateRequest.Team)
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
		newShift := model.Shift(*updateRequest.DefaultShift)
		if newShift != existingWorkInfoDomain.GetDefaultShift() {
			existingWorkInfoDomain.SetDefaultShift(newShift)
			fieldsUpdated = true
		}
	}
	if updateRequest.WeekdayOff != nil {
		newWeekdayOff := model.Weekday(*updateRequest.WeekdayOff)
		if newWeekdayOff != existingWorkInfoDomain.GetWeekdayOff() {
			existingWorkInfoDomain.SetWeekdayOff(newWeekdayOff)
			fieldsUpdated = true
		}
	}
	if updateRequest.WeekendDayOff != nil {
		newWeekendDayOff := model.WeekendDayOff(*updateRequest.WeekendDayOff)
		if newWeekendDayOff != existingWorkInfoDomain.GetWeekendDayOff() {
			existingWorkInfoDomain.SetWeekendDayOff(newWeekendDayOff)
			fieldsUpdated = true
		}
	}
	if updateRequest.SuperiorID != nil {
		newSuperiorID := *updateRequest.SuperiorID
		// Validar novo SuperiorID
		if newSuperiorID != "" { // Se estiver tentando definir um novo superior
			if newSuperiorID == userId { // Não pode ser ele mesmo
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
		// Aplicar alteração se diferente
		if newSuperiorID != existingWorkInfoDomain.GetSuperiorID() {
			existingWorkInfoDomain.SetSuperiorID(newSuperiorID) // Permite definir como string vazia para remover superior
			fieldsUpdated = true
		}
	}

	if !fieldsUpdated {
		logger.Info("No actual changes detected for WorkInfo update.",
			zap.String("journey", "updateWorkInfo"),
			zap.String("userID", userId))
		return existingWorkInfoDomain, nil // Retorna o domínio existente sem chamar o repo
	}

	// 4. Chamar o repositório para persistir as alterações
	// O método UpdateWorkInfo no repositório agora espera o userId e o domain completo.
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
	return existingWorkInfoDomain, nil // Retorna o domínio atualizado
}
