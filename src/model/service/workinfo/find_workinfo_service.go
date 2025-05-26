package workinfo

import (
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/logger"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/rest_err"

	// IMPORT ATUALIZADO: Agora importa do subpacote 'domain'
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model/domain"
	"go.uber.org/zap"
)

func (wd *workInfoDomainService) FindWorkInfoByUserIdServices(
	userId string,
) (domain.WorkInfoDomainInterface, *rest_err.RestErr) { // <<< USA domain.WorkInfoDomainInterface
	logger.Info("Init FindWorkInfoByUserIdServices",
		zap.String("journey", "findWorkInfoByUserId"),
		zap.String("userIdToFind", userId))

	if userId == "" {
		logger.Error("UserID for FindWorkInfoByUserIdServices cannot be empty", nil,
			zap.String("journey", "findWorkInfoByUserId"))
		return nil, rest_err.NewBadRequestError("User ID cannot be empty")
	}

	workInfoDomain, err := wd.workInfoRepository.FindWorkInfoByUserId(userId)
	if err != nil {
		logger.Error("Error calling repository to find work info by user ID", err,
			zap.String("journey", "findWorkInfoByUserId"),
			zap.String("userId", userId))
		return nil, err
	}

	logger.Info("FindWorkInfoByUserIdServices executed successfully",
		zap.String("journey", "findWorkInfoByUserId"),
		zap.String("userId", userId))
	return workInfoDomain, nil
}
