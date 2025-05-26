package swap

import (
	"github.com/Lipe-Azevedo/escala-fds/src/configuration/logger"
	"github.com/Lipe-Azevedo/escala-fds/src/configuration/rest_err"

	// IMPORT ATUALIZADO: Agora importa do subpacote 'domain'
	"github.com/Lipe-Azevedo/escala-fds/src/model/domain"
	"go.uber.org/zap"
	// "time" // Removido se não estiver sendo usado diretamente neste arquivo específico
)

func (ss *swapDomainService) CreateSwapServices(
	swapReqDomain domain.SwapDomainInterface, // <<< USA domain.SwapDomainInterface
) (domain.SwapDomainInterface, *rest_err.RestErr) { // <<< USA domain.SwapDomainInterface
	logger.Info("Init CreateSwapServices",
		zap.String("journey", "createSwap"),
		zap.String("requesterId", swapReqDomain.GetRequesterID()))

	// O NewSwapDomain (chamado pelo controller) já define o status como Pending
	// e CreatedAt como time.Now().

	domainResult, err := ss.repository.CreateSwap(swapReqDomain)
	if err != nil {
		logger.Error("Error calling repository to create swap", err,
			zap.String("journey", "createSwap"))
		return nil, err
	}

	logger.Info("CreateSwapServices executed successfully",
		zap.String("swapID", domainResult.GetID()),
		zap.String("journey", "createSwap"))

	return domainResult, nil
}
