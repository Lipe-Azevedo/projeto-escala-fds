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

func (sc *shiftSwapControllerInterface) UpdateShiftSwapStatus(c *gin.Context) {
	logger.Info("Init UpdateShiftSwapStatus controller",
		zap.String("journey", "updateShiftSwapStatus"))

	id := c.Param("id")

	var shiftSwapRequest request.ShiftSwapRequest

	if err := c.ShouldBindJSON(&shiftSwapRequest); err != nil {
		logger.Error("Error validating status request", err,
			zap.String("journey", "updateShiftSwapStatus"))
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	approverID := c.GetString("userID") // Obtém do middleware de autenticação

	domain := model.NewShiftSwapUpdateDomain(
		shiftSwapRequest.RequestedID,
		model.Shift(shiftSwapRequest.CurrentShift),
		model.Shift(shiftSwapRequest.NewShift),
		model.Weekday(shiftSwapRequest.CurrentDayOff),
		model.Weekday(shiftSwapRequest.NewDayOff),
		shiftSwapRequest.Reason,
	)

	domain.SetApprovedBy(approverID)

	err := sc.service.UpdateShiftSwapServices(id, domain)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.Status(http.StatusOK)
}
