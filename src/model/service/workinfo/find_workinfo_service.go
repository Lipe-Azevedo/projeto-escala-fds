package workinfo

import (
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/logger"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model"
	"go.uber.org/zap"
)

func (wd *workInfoDomainService) FindWorkInfoByUserIdServices(
	userId string,
) (model.WorkInfoDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindWorkInfoByUserIdServices", // Nome da função no log
		zap.String("journey", "findWorkInfoByUserId"), // Journey padronizado
		zap.String("userIdToFind", userId))

	if userId == "" { // Adicionada verificação de userId vazio
		logger.Error("UserID for FindWorkInfoByUserIdServices cannot be empty", nil,
			zap.String("journey", "findWorkInfoByUserId"))
		return nil, rest_err.NewBadRequestError("User ID cannot be empty")
	}

	workInfoDomain, err := wd.workInfoRepository.FindWorkInfoByUserId(userId)
	if err != nil {
		logger.Error("Error calling repository to find work info by user ID", err,
			zap.String("journey", "findWorkInfoByUserId"),
			zap.String("userId", userId))
		return nil, err // Erro já é formatado pelo repositório (ex: NotFoundError)
	}

	logger.Info("FindWorkInfoByUserIdServices executed successfully",
		zap.String("journey", "findWorkInfoByUserId"),
		zap.String("userId", userId))
	return workInfoDomain, nil
}
