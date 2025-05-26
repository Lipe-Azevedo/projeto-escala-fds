package swap

import (
	"github.com/Lipe-Azevedo/escala-fds/src/configuration/logger"
	"github.com/Lipe-Azevedo/escala-fds/src/configuration/rest_err"

	// IMPORT ATUALIZADO: Agora importa do subpacote 'domain'
	"github.com/Lipe-Azevedo/escala-fds/src/model/domain"
	"go.uber.org/zap"
)

func (ss *swapDomainService) FindSwapByIDServices(
	id string,
) (domain.SwapDomainInterface, *rest_err.RestErr) { // <<< USA domain.SwapDomainInterface
	logger.Info("Init FindSwapByIDServices",
		zap.String("journey", "findSwapByID"),
		zap.String("swapIDToFind", id))

	if id == "" {
		logger.Error("Swap ID for FindSwapByIDServices cannot be empty", nil,
			zap.String("journey", "findSwapByID"))
		return nil, rest_err.NewBadRequestError("Swap ID cannot be empty")
	}

	swapDomainVal, err := ss.repository.FindSwapByID(id) // Renomeado para swapDomainVal
	if err != nil {
		logger.Error("Error calling repository to find swap by ID", err,
			zap.String("journey", "findSwapByID"),
			zap.String("swapID", id))
		return nil, err
	}

	logger.Info("FindSwapByIDServices executed successfully",
		zap.String("journey", "findSwapByID"),
		zap.String("swapID", id))
	return swapDomainVal, nil
}

/*
// Se você decidir expor FindSwapsByUserID ou FindSwapsByStatus através do serviço:
func (ss *swapDomainService) FindSwapsByUserIDServices(
    userID string,
) ([]domain.SwapDomainInterface, *rest_err.RestErr) { // <<< USA domain.SwapDomainInterface
    logger.Info("Init FindSwapsByUserIDServices",
        zap.String("journey", "findSwapsByUserID"),
        zap.String("userID", userID))

    if userID == "" {
        return nil, rest_err.NewBadRequestError("User ID cannot be empty")
    }
    // ss.repository.FindSwapsByUserID já deve retornar []domain.SwapDomainInterface
    return ss.repository.FindSwapsByUserID(userID)
}

func (ss *swapDomainService) FindSwapsByStatusServices(
    status domain.SwapStatus, // <<< USA domain.SwapStatus
) ([]domain.SwapDomainInterface, *rest_err.RestErr) { // <<< USA domain.SwapDomainInterface
    logger.Info("Init FindSwapsByStatusServices",
        zap.String("journey", "findSwapsByStatus"),
        zap.String("status", string(status)))

    // ss.repository.FindSwapsByStatus já deve retornar []domain.SwapDomainInterface
    return ss.repository.FindSwapsByStatus(status)
}
*/
