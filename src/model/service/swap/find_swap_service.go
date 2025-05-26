package swap

import (
	"github.com/Lipe-Azevedo/escala-fds/src/configuration/logger"
	"github.com/Lipe-Azevedo/escala-fds/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/escala-fds/src/model"
	"go.uber.org/zap"
)

func (ss *swapDomainService) FindSwapByIDServices(
	id string,
) (model.SwapDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindSwapByIDServices", // Nome da função no log
		zap.String("journey", "findSwapByID"),
		zap.String("swapIDToFind", id))

	if id == "" {
		logger.Error("Swap ID for FindSwapByIDServices cannot be empty", nil,
			zap.String("journey", "findSwapByID"))
		return nil, rest_err.NewBadRequestError("Swap ID cannot be empty")
	}

	swapDomain, err := ss.repository.FindSwapByID(id)
	if err != nil {
		logger.Error("Error calling repository to find swap by ID", err,
			zap.String("journey", "findSwapByID"),
			zap.String("swapID", id))
		return nil, err
	}

	logger.Info("FindSwapByIDServices executed successfully",
		zap.String("journey", "findSwapByID"),
		zap.String("swapID", id))
	return swapDomain, nil
}

// Se você decidir expor FindSwapsByUserID ou FindSwapsByStatus através do serviço,
// os métodos seriam adicionados aqui, por exemplo:
/*
func (ss *swapDomainService) FindSwapsByUserIDServices(
    userID string,
) ([]model.SwapDomainInterface, *rest_err.RestErr) {
    logger.Info("Init FindSwapsByUserIDServices",
        zap.String("journey", "findSwapsByUserID"),
        zap.String("userID", userID))

    if userID == "" {
        return nil, rest_err.NewBadRequestError("User ID cannot be empty")
    }
    return ss.repository.FindSwapsByUserID(userID)
}

func (ss *swapDomainService) FindSwapsByStatusServices(
    status model.SwapStatus,
) ([]model.SwapDomainInterface, *rest_err.RestErr) {
    logger.Info("Init FindSwapsByStatusServices",
        zap.String("journey", "findSwapsByStatus"),
        zap.String("status", string(status)))

    return ss.repository.FindSwapsByStatus(status)
}
*/
