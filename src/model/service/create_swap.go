package service

import (
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/logger"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model"
	"go.uber.org/zap"
)

func (ss *swapDomainService) CreateSwapServices(
	swapDomain model.SwapDomainInterface,
) (model.SwapDomainInterface, *rest_err.RestErr) {
	logger.Info("Init CreateSwap service",
		zap.String("journey", "createSwap"))

	// Validação: Verificar se os usuários envolvidos existem e pertencem à mesma equipe
	// (Implementar conforme necessidade)

	domainResult, err := ss.repository.CreateSwap(swapDomain)
	if err != nil {
		logger.Error("Error trying to call repository",
			err,
			zap.String("journey", "createSwap"))
		return nil, err
	}

	logger.Info("Swap created successfully",
		zap.String("swapID", domainResult.GetID()),
		zap.String("journey", "createSwap"))

	return domainResult, nil
}
