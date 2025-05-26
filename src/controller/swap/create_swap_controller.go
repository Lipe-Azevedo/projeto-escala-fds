package swap

import (
	"net/http"

	"github.com/Lipe-Azevedo/escala-fds/src/configuration/logger"
	"github.com/Lipe-Azevedo/escala-fds/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/escala-fds/src/configuration/validation"

	// Import para o DTO de request de swap, usando alias
	swap_request_dto "github.com/Lipe-Azevedo/escala-fds/src/controller/swap/request"
	"github.com/Lipe-Azevedo/escala-fds/src/model/domain"
	"github.com/Lipe-Azevedo/escala-fds/src/view"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (sc *swapControllerInterface) CreateSwap(c *gin.Context) {
	logger.Info("Init CreateSwap controller",
		zap.String("journey", "createSwap"))

	var swapRequest swap_request_dto.SwapRequest
	if err := c.ShouldBindJSON(&swapRequest); err != nil {
		logger.Error("Error validating swap request data for creation", err,
			zap.String("journey", "createSwap"))
		restErrVal := validation.ValidateUserError(err)
		c.JSON(restErrVal.Code, restErrVal)
		return
	}

	if swapRequest.RequestedID == "" || swapRequest.CurrentShift == "" || swapRequest.NewShift == "" || swapRequest.CurrentDayOff == "" || swapRequest.NewDayOff == "" {
		logger.Error("Missing required fields for swap creation", nil,
			zap.String("journey", "createSwap"))
		restErr := rest_err.NewBadRequestError("Missing required fields for swap creation (requested_id, current_shift, new_shift, current_day_off, new_day_off)")
		c.JSON(restErr.Code, restErr)
		return
	}

	requesterID := "temp-requester-id"
	if requesterID == "" {
		logger.Error("Requester ID not found (simulate JWT)", nil, zap.String("journey", "createSwap"))
		restErr := rest_err.NewUnauthorizedError("Unauthorized: Requester ID not found.")
		c.JSON(restErr.Code, restErr)
		return
	}

	domain := domain.NewSwapDomain(
		requesterID,
		swapRequest.RequestedID,
		domain.Shift(swapRequest.CurrentShift),
		domain.Shift(swapRequest.NewShift),
		domain.Weekday(swapRequest.CurrentDayOff),
		domain.Weekday(swapRequest.NewDayOff),
		swapRequest.Reason,
	)

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

	c.JSON(http.StatusCreated, view.ConvertSwapDomainToResponse(domainResult))
}
