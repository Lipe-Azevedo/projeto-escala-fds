package controller

import (
	"net/http"

	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/logger"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/validation"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/controller/model/request"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/view"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (sc *shiftSwapControllerInterface) CreateShiftSwap(c *gin.Context) {
	logger.Info("Init CreateShiftSwap controller",
		zap.String("journey", "createShiftSwap"))

	var shiftSwapRequest request.ShiftSwapRequest
	if err := c.ShouldBindJSON(&shiftSwapRequest); err != nil {
		logger.Error("Error validating shift swap request", err,
			zap.String("journey", "createShiftSwap"))
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	requesterID := c.GetString("userID") // Obtém do middleware de autenticação

	domain := model.NewShiftSwapDomain(
		requesterID,
		shiftSwapRequest.RequestedID,
		model.Shift(shiftSwapRequest.CurrentShift),
		model.Shift(shiftSwapRequest.NewShift),
		model.Weekday(shiftSwapRequest.CurrentDayOff),
		model.Weekday(shiftSwapRequest.NewDayOff),
		shiftSwapRequest.Reason,
	)

	domainResult, err := sc.service.CreateShiftSwapServices(domain)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusCreated, view.ConvertShiftSwapDomainToResponse(domainResult))
}