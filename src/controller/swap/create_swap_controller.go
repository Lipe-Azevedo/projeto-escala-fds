package swap

import (
	"net/http"

	"github.com/Lipe-Azevedo/escala-fds/src/configuration/logger"
	"github.com/Lipe-Azevedo/escala-fds/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/escala-fds/src/configuration/validation"
	swap_request_dto "github.com/Lipe-Azevedo/escala-fds/src/controller/swap/request"
	"github.com/Lipe-Azevedo/escala-fds/src/model/domain"
	"github.com/Lipe-Azevedo/escala-fds/src/view"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (sc *swapControllerInterface) CreateSwap(c *gin.Context) {
	logger.Info("Init CreateSwap controller",
		zap.String("journey", "createSwap"))

	requesterIDClaim, exists := c.Get("userID")
	if !exists {
		errMsg := "Failed to get userID from token for creating swap"
		logger.Error(errMsg, nil, zap.String("journey", "createSwap"))
		c.JSON(http.StatusInternalServerError, rest_err.NewInternalServerError(errMsg))
		return
	}
	requesterID := requesterIDClaim.(string)

	var swapRequest swap_request_dto.SwapRequest
	if err := c.ShouldBindJSON(&swapRequest); err != nil {
		logger.Error("Error validating swap request data for creation", err,
			zap.String("journey", "createSwap"))
		restErrVal := validation.ValidateUserError(err)
		c.JSON(restErrVal.Code, restErrVal)
		return
	}

	if swapRequest.RequestedID == "" ||
		swapRequest.CurrentShift == "" || swapRequest.NewShift == "" ||
		swapRequest.CurrentDayOff == "" || swapRequest.NewDayOff == "" {
		errMsg := "Missing required fields for swap creation (requested_id, current_shift, new_shift, current_day_off, new_day_off)"
		logger.Error(errMsg, nil, zap.String("journey", "createSwap"))
		restErr := rest_err.NewBadRequestError(errMsg)
		c.JSON(restErr.Code, restErr)
		return
	}

	swapDomain := domain.NewSwapDomain(
		requesterID,
		swapRequest.RequestedID,
		domain.Shift(swapRequest.CurrentShift),
		domain.Shift(swapRequest.NewShift),
		domain.Weekday(swapRequest.CurrentDayOff),
		domain.Weekday(swapRequest.NewDayOff),
		swapRequest.Reason,
	)

	domainResult, serviceErr := sc.service.CreateSwapServices(swapDomain)
	if serviceErr != nil {
		c.JSON(serviceErr.Code, serviceErr)
		return
	}

	logger.Info("Swap created successfully via controller",
		zap.String("swapId", domainResult.GetID()),
		zap.String("requesterId", requesterID),
		zap.String("journey", "createSwap"))

	c.JSON(http.StatusCreated, view.ConvertSwapDomainToResponse(domainResult))
}
