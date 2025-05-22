package controller

import (
	"net/http"

	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/logger"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/view"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (sc *swapControllerInterface) FindSwapByID(c *gin.Context) {
	logger.Info("Init FindSwapByID controller",
		zap.String("journey", "findSwapByID"))

	id := c.Param("id")

	swapDomain, err := sc.service.FindSwapByIDServices(id)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, view.ConvertSwapDomainToResponse(swapDomain))
}
