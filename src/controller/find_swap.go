package controller

import (
	"net/http"

	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/logger"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/view"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (sc *shiftSwapControllerInterface) FindShiftSwapByID(c *gin.Context) {
	logger.Info("Init FindShiftSwapByID controller",
		zap.String("journey", "findShiftSwapByID"))

	id := c.Param("id")

	shiftSwapDomain, err := sc.service.FindShiftSwapByIDServices(id)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, view.ConvertShiftSwapDomainToResponse(shiftSwapDomain))
}
