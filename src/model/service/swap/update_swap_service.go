package swap

import (
	// "fmt" // Removido se não usado
	// IMPORT ATUALIZADO para o novo nome do módulo
	"time"

	"github.com/Lipe-Azevedo/escala-fds/src/configuration/logger"
	"github.com/Lipe-Azevedo/escala-fds/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/escala-fds/src/model/domain"
	"go.uber.org/zap"
)

func (ss *swapDomainService) UpdateSwapServices(
	id string,
	swapUpdateInfo domain.SwapDomainInterface,
) *rest_err.RestErr {
	logger.Info("Init UpdateSwapServices",
		zap.String("journey", "updateSwap"),
		zap.String("swapID", id))

	if id == "" {
		logger.Error("Swap ID for UpdateSwapServices cannot be empty", nil,
			zap.String("journey", "updateSwap"))
		return rest_err.NewBadRequestError("Swap ID cannot be empty")
	}

	existingSwap, findErr := ss.repository.FindSwapByID(id)
	if findErr != nil {
		logger.Error("Swap to update not found by service", findErr,
			zap.String("journey", "updateSwap"),
			zap.String("swapID", id))
		return findErr
	}

	newStatus := swapUpdateInfo.GetStatus()
	existingSwap.SetStatus(newStatus)

	if newStatus == domain.StatusApproved {
		if swapUpdateInfo.GetApprovedBy() != nil && *swapUpdateInfo.GetApprovedBy() != "" {
			existingSwap.SetApprovedBy(*swapUpdateInfo.GetApprovedBy())
		}
		if swapUpdateInfo.GetApprovedAt() != nil && !swapUpdateInfo.GetApprovedAt().IsZero() {
			existingSwap.SetApprovedAt(*swapUpdateInfo.GetApprovedAt())
		} else {
			now := time.Now()
			existingSwap.SetApprovedAt(now)
		}
	}
	// Adicionar lógica para limpar ApprovedBy/At se o status mudar para pending/rejected, se necessário.

	err := ss.repository.UpdateSwap(id, existingSwap)
	if err != nil {
		logger.Error("Error calling repository to update swap", err,
			zap.String("journey", "updateSwap"),
			zap.String("swapID", id))
		return err
	}

	logger.Info("UpdateSwapServices executed successfully",
		zap.String("swapID", id),
		zap.String("journey", "updateSwap"))

	return nil
}
