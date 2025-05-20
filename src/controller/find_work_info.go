package controller

import (
	"net/http"

	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/logger"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/view"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (wc *workInfoControllerInterface) FindWorkInfoByUserId(c *gin.Context) {
	logger.Info("Init findWorkInfoByUserId controller",
		zap.String("journey", "findWorkInfoByUserId"))

	userId := c.Param("userId")

	workInfoDomain, err := wc.service.FindWorkInfoByUserIdServices(userId)
	if err != nil {
		logger.Error("Error trying to call findWorkInfoByUserId service",
			err,
			zap.String("journey", "findWorkInfoByUserId"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("FindWorkInfoByUserId controller executed successfully",
		zap.String("userId", userId),
		zap.String("journey", "findWorkInfoByUserId"))

	c.JSON(http.StatusOK, view.ConvertWorkInfoDomainToResponse(workInfoDomain))
}
