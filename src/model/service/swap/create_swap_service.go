package swap

import (
	"github.com/Lipe-Azevedo/escala-fds/src/configuration/logger"
	"github.com/Lipe-Azevedo/escala-fds/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/escala-fds/src/model"
	"go.uber.org/zap"
	// Para time.Now() no SetApprovedAt e para o status
)

func (ss *swapDomainService) CreateSwapServices(
	swapDomain model.SwapDomainInterface,
) (model.SwapDomainInterface, *rest_err.RestErr) {
	logger.Info("Init CreateSwapServices", // Nome da função no log
		zap.String("journey", "createSwap"),
		zap.String("requesterId", swapDomain.GetRequesterID()))

	// Validações que podem ser adicionadas aqui:
	// 1. Verificar se RequesterID e RequestedID existem (usando um UserDomainService, se disponível).
	//    Por agora, essa dependência não foi adicionada ao SwapService.
	// 2. Verificar se RequesterID != RequestedID.
	// 3. Verificar se os turnos/folgas são válidos ou fazem sentido (ex: não trocar para o mesmo turno/folga).
	// 4. Definir o status inicial e CreatedAt (o NewSwapDomain já faz isso).
	//    swapDomain.SetStatus(model.StatusPending)
	//    swapDomain.SetCreatedAt(time.Now()) // O seu NewSwapDomain já faz isso.

	// A interface SwapDomainInterface tem SetID, mas quem define o ID é o repositório.
	// O NewSwapDomain já define status como Pending e CreatedAt como time.Now().

	domainResult, err := ss.repository.CreateSwap(swapDomain)
	if err != nil {
		logger.Error("Error calling repository to create swap", err,
			zap.String("journey", "createSwap"))
		return nil, err
	}

	logger.Info("CreateSwapServices executed successfully",
		zap.String("swapID", domainResult.GetID()), // ID é setado pelo repositório
		zap.String("journey", "createSwap"))

	return domainResult, nil
}
