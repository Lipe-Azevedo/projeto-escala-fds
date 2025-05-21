package service

import (
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/logger"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model"
	"go.uber.org/zap"
)

func (ss *shiftSwapDomainService) UpdateShiftSwapServices(
	id string,
	shiftSwapDomain model.ShiftSwapDomainInterface,
) *rest_err.RestErr {
	logger.Info("Init UpdateShiftSwap service",
		zap.String("journey", "updateShiftSwap"))

	// Validação: Verificar se o usuário que está aprovando é um master/superior
	// (Implementar conforme necessidade)

	err := ss.repository.UpdateShiftSwap(id, shiftSwapDomain)
	if err != nil {
		logger.Error("Error trying to call repository",
			err,
			zap.String("journey", "updateShiftSwap"))
		return err
	}

	logger.Info("ShiftSwap updated successfully",
		zap.String("shiftSwapID", id),
		zap.String("journey", "updateShiftSwap"))

	return nil
}