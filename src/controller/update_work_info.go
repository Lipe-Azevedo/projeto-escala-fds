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

func (wc *workInfoControllerInterface) UpdateWorkInfo(c *gin.Context) {
	logger.Info("Init updateWorkInfo controller",
		zap.String("journey", "updateWorkInfo"))

	var workInfoRequest request.WorkInfoRequest
	if err := c.ShouldBindJSON(&workInfoRequest); err != nil {
		logger.Error("Error validating work info",
			err,
			zap.String("journey", "updateWorkInfo"))
		errRest := validation.ValidateUserError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	userId := c.Param("userId")

	domain := model.NewWorkInfoDomain(
		userId,
		model.Team(workInfoRequest.Team),
		workInfoRequest.Position,
		model.Shift(workInfoRequest.DefaultShift),
		model.Weekday(workInfoRequest.WeekdayOff),
		model.WeekendDayOff(workInfoRequest.WeekendDayOff),
		workInfoRequest.SuperiorID,
	)

	err := wc.service.UpdateWorkInfoServices(userId, domain)
	if err != nil {
		logger.Error("Error trying to call updateWorkInfo service",
			err,
			zap.String("journey", "updateWorkInfo"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("UpdateWorkInfo controller executed successfully",
		zap.String("userId", userId),
		zap.String("journey", "updateWorkInfo"))

	c.Status(http.StatusOK)
}
