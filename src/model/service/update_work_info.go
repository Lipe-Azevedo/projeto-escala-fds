package service

import (
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/logger"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model"
	"go.uber.org/zap"
)

func (wd *workInfoDomainService) UpdateWorkInfoServices(
	userId string,
	workInfoDomain model.WorkInfoDomainInterface,
) *rest_err.RestErr {
	logger.Info("Init UpdateWorkInfo service", zap.String("journey", "updateWorkInfo"))

	err := wd.workInfoRepository.UpdateWorkInfo(userId, workInfoDomain)
	if err != nil {
		logger.Error(
			"Error tyring to call repository.",
			err,
			zap.String("journey", "updateWorkInfo"),
		)
		return err
	}

	logger.Info(
		"updateWorkInfo service executed successfully.",
		zap.String("userId", userId),
		zap.String("journey", "updateWorkInfo"),
	)
	return nil
}
