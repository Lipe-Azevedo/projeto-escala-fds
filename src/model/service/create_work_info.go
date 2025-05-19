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
	logger.Info("Init CreateWorkInfo service", zap.String("journey", "createWorkInfo"))

	workInfo, err := wd.workInfoRepository.CreateWorkInfo(workInfoDomain)
	if err != nil {
		logger.Error("Error trying to create work info", err, zap.String("journey", "createWorkInfo"))
		return nil, err
	}

	logger.Info("CreateWorkInfo service executed successfully",
		zap.String("userId", workInfo.GetUserId()),
		zap.String("journey", "createWorkInfo"))

	return workInfo, nil
}
