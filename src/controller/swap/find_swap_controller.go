package swap

import (
	"net/http"

	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/logger"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/view"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive" // Para validar o formato do ID
	"go.uber.org/zap"
)

func (sc *swapControllerInterface) FindSwapByID(c *gin.Context) {
	logger.Info("Init FindSwapByID controller",
		zap.String("journey", "findSwapByID"))

	swapID := c.Param("id")
	if _, err := primitive.ObjectIDFromHex(swapID); err != nil {
		logger.Error("Invalid swap ID format in FindSwapByID controller", err,
			zap.String("journey", "findSwapByID"),
			zap.String("swapID", swapID))
		restErrVal := rest_err.NewBadRequestError("Invalid Swap ID format, must be a hex value.")
		c.JSON(restErrVal.Code, restErrVal)
		return
	}

	// TODO: (Pós-JWT) Lógica de Permissão:
	// - Master pode ver qualquer Swap.
	// - Colaborador só pode ver Swaps onde ele é requester ou requested.
	//   Isso exigiria buscar o swap e então verificar os IDs contra o c.GetString("userID").

	swapDomain, serviceErr := sc.service.FindSwapByIDServices(swapID)
	if serviceErr != nil {
		logger.Error("Failed to call find swap by ID service", serviceErr,
			zap.String("journey", "findSwapByID"),
			zap.String("swapID", swapID))
		c.JSON(serviceErr.Code, serviceErr)
		return
	}

	logger.Info("FindSwapByID controller executed successfully",
		zap.String("swapID", swapID),
		zap.String("journey", "findSwapByID"))

	c.JSON(http.StatusOK, view.ConvertSwapDomainToResponse(swapDomain))
}
