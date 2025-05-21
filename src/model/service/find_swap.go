package service

import (
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/logger"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model"
	"go.uber.org/zap"
)

func (ss *shiftSwapDomainService) FindShiftSwapByIDServices(
	id string,
) (model.ShiftSwapDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindShiftSwapByID service",
		zap.String("journey", "findShiftSwapByID"))

	return ss.repository.FindShiftSwapByID(id)
}