package service

import (
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/logger"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model"
	"go.uber.org/zap"
)

func (ss *swapDomainService) UpdateSwapServices(
	id string,
	swapDomain model.SwapDomainInterface,
) *rest_err.RestErr {
	logger.Info("Init UpdateSwap service",
		zap.String("journey", "updateSwap"))

	// Validação: Verificar se o usuário que está aprovando é um master/superior
	// (Implementar conforme necessidade)

	err := ss.repository.UpdateSwap(id, swapDomain)
	if err != nil {
		logger.Error("Error trying to call repository",
			err,
			zap.String("journey", "updateSwap"))
		return err
	}

	logger.Info("Swap updated successfully",
		zap.String("swapID", id),
		zap.String("journey", "updateSwap"))

	return nil
}
