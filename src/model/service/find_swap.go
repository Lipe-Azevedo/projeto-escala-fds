package service

import (
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/logger"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model"
	"go.uber.org/zap"
)

func (ss *swapDomainService) FindSwapByIDServices(
	id string,
) (model.SwapDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindSwapByID service",
		zap.String("journey", "findSwapByID"))

	return ss.repository.FindSwapByID(id)
}
