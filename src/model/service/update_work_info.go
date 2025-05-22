package service

import (
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/logger"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model"
	"go.uber.org/zap"
)

func (wd *workInfoDomainService) UpdateWorkInfoServices(
	userId string, // ID do usuário cujo WorkInfo está sendo atualizado
	workInfoDomain model.WorkInfoDomainInterface,
) *rest_err.RestErr {
	logger.Info("Init UpdateWorkInfoServices", zap.String("journey", "updateWorkInfo")) // Nome da função corrigido no log

	// Validação 1: Verificar se o usuário para o qual o WorkInfo está sendo atualizado existe.
	// Embora o repositório retorne NotFound se o WorkInfo não existir (baseado no userId),
	// é uma boa prática garantir que o próprio usuário exista primeiro.
	_, errUser := wd.userDomainService.FindUserByIDServices(userId)
	if errUser != nil {
		logger.Error("User for WorkInfo update not found", errUser,
			zap.String("journey", "updateWorkInfo"),
			zap.String("userID", userId))
		if errUser.Code == 404 {
			return rest_err.NewBadRequestError("User for WorkInfo update does not exist")
		}
		return errUser
	}

	// Validação 2: Verificar se o SuperiorID (se fornecido e não vazio no domain de atualização) é válido
	superiorID := workInfoDomain.GetSuperiorID()
	if superiorID != "" {
		_, errSuperior := wd.userDomainService.FindUserByIDServices(superiorID)
		if errSuperior != nil {
			logger.Error("Superior user (for update) not found", errSuperior,
				zap.String("journey", "updateWorkInfo"),
				zap.String("superiorID", superiorID))
			if errSuperior.Code == 404 {
				return rest_err.NewBadRequestError("Superior user specified in WorkInfo update does not exist")
			}
			return errSuperior
		}
	}

	// Validação 3: Garantir que o workInfoDomain.GetUserId() (se presente e usado internamente pelo domain)
	// corresponda ao userId do parâmetro, para evitar inconsistências.
	// No seu NewWorkInfoDomain, o UserID é o primeiro parâmetro.
	// E o WorkInfoDomain passado para UpdateWorkInfoServices é construído com o userId da rota.
	// Então workInfoDomain.GetUserId() deve ser igual a userId.
	if workInfoDomain.GetUserId() != userId {
		logger.Error("Mismatch between route userId and workInfoDomain's userId during update", nil,
			zap.String("journey", "updateWorkInfo"),
			zap.String("routeUserId", userId),
			zap.String("domainUserId", workInfoDomain.GetUserId()))
		return rest_err.NewBadRequestError("User ID mismatch in update request")
	}

	err := wd.workInfoRepository.UpdateWorkInfo(userId, workInfoDomain)
	if err != nil {
		logger.Error(
			"Error trying to call repository for WorkInfo update.", // "tyring to call repository."
			err,
			zap.String("journey", "updateWorkInfo"),
		)
		return err
	}

	logger.Info(
		"UpdateWorkInfoServices executed successfully.", // Nome da função corrigido no log
		zap.String("userId", userId),
		zap.String("journey", "updateWorkInfo"),
	)
	return nil
}
