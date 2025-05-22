package controller

import (
	"net/http"

	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/logger"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/validation"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/controller/model/request"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (sc *swapControllerInterface) UpdateSwapStatus(c *gin.Context) {
	logger.Info("Init UpdateSwapStatus controller",
		zap.String("journey", "updateSwapStatus"))

	id := c.Param("id")

	var swapRequest request.SwapRequest

	if err := c.ShouldBindJSON(&swapRequest); err != nil {
		logger.Error("Error validating status request", err,
			zap.String("journey", "updateSwapStatus"))
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	approverID := c.GetString("userID") // Obtém do middleware de autenticação

	domain := model.NewSwapUpdateDomain(
		swapRequest.RequestedID,
		model.Shift(swapRequest.CurrentShift),
		model.Shift(swapRequest.NewShift),
		model.Weekday(swapRequest.CurrentDayOff),
		model.Weekday(swapRequest.NewDayOff),
		swapRequest.Reason,
	)

	domain.SetApprovedBy(approverID)

	err := sc.service.UpdateSwapServices(id, domain)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.Status(http.StatusOK)
}
