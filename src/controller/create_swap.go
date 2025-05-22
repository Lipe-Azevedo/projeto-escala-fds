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

func (sc *swapControllerInterface) CreateSwap(c *gin.Context) {
	logger.Info("Init CreateSwap controller",
		zap.String("journey", "createSwap"))

	var swapRequest request.SwapRequest

	if err := c.ShouldBindJSON(&swapRequest); err != nil {
		logger.Error("Error validating shift swap request", err,
			zap.String("journey", "createSwap"))
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	requesterID := c.GetString("userID") // Obtém do middleware de autenticação

	domain := model.NewSwapDomain(
		requesterID,
		swapRequest.RequestedID,
		model.Shift(swapRequest.CurrentShift),
		model.Shift(swapRequest.NewShift),
		model.Weekday(swapRequest.CurrentDayOff),
		model.Weekday(swapRequest.NewDayOff),
		swapRequest.Reason,
	)

	domainResult, err := sc.service.CreateSwapServices(domain)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusCreated, view.ConvertSwapDomainToResponse(domainResult))
}
