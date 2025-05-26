package swap

import (
	"net/http"

	"github.com/Lipe-Azevedo/escala-fds/src/configuration/logger"
	"github.com/Lipe-Azevedo/escala-fds/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/escala-fds/src/configuration/validation"

	// Import para o DTO de request de swap, usando alias
	swap_request_dto "github.com/Lipe-Azevedo/escala-fds/src/controller/swap/request"
	"github.com/Lipe-Azevedo/escala-fds/src/model"
	"github.com/Lipe-Azevedo/escala-fds/src/view"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (sc *swapControllerInterface) CreateSwap(c *gin.Context) {
	logger.Info("Init CreateSwap controller",
		zap.String("journey", "createSwap"))

	var swapRequest swap_request_dto.SwapRequest // Usando DTO específico
	if err := c.ShouldBindJSON(&swapRequest); err != nil {
		logger.Error("Error validating swap request data for creation", err, // Mensagem de log mais genérica
			zap.String("journey", "createSwap"))
		restErrVal := validation.ValidateUserError(err) // Nome da var para evitar conflito com rest_err do pacote
		c.JSON(restErrVal.Code, restErrVal)
		return
	}

	// Validação para garantir que os campos de criação estão presentes
	if swapRequest.RequestedID == "" || swapRequest.CurrentShift == "" || swapRequest.NewShift == "" || swapRequest.CurrentDayOff == "" || swapRequest.NewDayOff == "" {
		logger.Error("Missing required fields for swap creation", nil,
			zap.String("journey", "createSwap"))
		restErr := rest_err.NewBadRequestError("Missing required fields for swap creation (requested_id, current_shift, new_shift, current_day_off, new_day_off)")
		c.JSON(restErr.Code, restErr)
		return
	}

	// TODO: (Pós-JWT) Obter o requesterID do token JWT.
	// Por enquanto, vamos simular ou deixar um placeholder.
	// requesterID := c.GetString("userID") // Exemplo: viria do middleware JWT
	requesterID := "temp-requester-id" // Placeholder - REMOVER/AJUSTAR COM JWT
	if requesterID == "" {
		logger.Error("Requester ID not found (simulate JWT)", nil, zap.String("journey", "createSwap"))
		restErr := rest_err.NewUnauthorizedError("Unauthorized: Requester ID not found.")
		c.JSON(restErr.Code, restErr)
		return
	}

	domain := model.NewSwapDomain(
		requesterID, // Usar o ID do usuário autenticado
		swapRequest.RequestedID,
		model.Shift(swapRequest.CurrentShift),
		model.Shift(swapRequest.NewShift),
		model.Weekday(swapRequest.CurrentDayOff),
		model.Weekday(swapRequest.NewDayOff),
		swapRequest.Reason,
	)
	// O status é definido como "pending" e CreatedAt por NewSwapDomain.

	domainResult, serviceErr := sc.service.CreateSwapServices(domain)
	if serviceErr != nil {
		logger.Error("Failed to call swap creation service", serviceErr,
			zap.String("journey", "createSwap"))
		c.JSON(serviceErr.Code, serviceErr)
		return
	}

	logger.Info("Swap created successfully via controller",
		zap.String("swapId", domainResult.GetID()),
		zap.String("journey", "createSwap"))

	// view.ConvertSwapDomainToResponse será ajustado para usar swap_response_dto
	c.JSON(http.StatusCreated, view.ConvertSwapDomainToResponse(domainResult))
}
