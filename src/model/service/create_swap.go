package service

import (
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/logger"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model"
	"go.uber.org/zap"
)

func (ss *shiftSwapDomainService) CreateShiftSwapServices(
	shiftSwapDomain model.ShiftSwapDomainInterface,
) (model.ShiftSwapDomainInterface, *rest_err.RestErr) {
	logger.Info("Init CreateShiftSwap service",
		zap.String("journey", "createShiftSwap"))

	// Validação: Verificar se os usuários envolvidos existem e pertencem à mesma equipe
	// (Implementar conforme necessidade)

	domainResult, err := ss.repository.CreateShiftSwap(shiftSwapDomain)
	if err != nil {
		logger.Error("Error trying to call repository",
			err,
			zap.String("journey", "createShiftSwap"))
		return nil, err
	}

	logger.Info("ShiftSwap created successfully",
		zap.String("shiftSwapID", domainResult.GetID()),
		zap.String("journey", "createShiftSwap"))

	return domainResult, nil
}