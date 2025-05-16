package controller

import (
	"net/http"

	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/logger"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/controller/model/request"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type WorkInfoControllerInterface interface {
	UpdateWorkInfo(c *gin.Context)
}

type workInfoController struct {
	service model.WorkInfoServiceInterface
}

func NewWorkInfoController(
	service model.WorkInfoServiceInterface,
) WorkInfoControllerInterface {
	return &workInfoController{
		service: service,
	}
}

func (wc *workInfoController) UpdateWorkInfo(c *gin.Context) {
	logger.Info("Init UpdateWorkInfo controller",
		zap.String("journey", "updateWorkInfo"))

	userId := c.Param("userId")

	var workInfoRequest request.WorkInfoRequest
	if err := c.ShouldBindJSON(&workInfoRequest); err != nil {
		logger.Error("Error validating work info data",
			err,
			zap.String("journey", "updateWorkInfo"))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	workInfo := model.WorkInfo{
		Team:          workInfoRequest.Team,
		Position:      workInfoRequest.Position,
		DefaultShift:  model.Shift(workInfoRequest.DefaultShift),
		WeekdayOff:    model.Weekday(workInfoRequest.WeekdayOff),
		WeekendDayOff: model.WeekendDayOff(workInfoRequest.WeekendDayOff),
		SuperiorID:    workInfoRequest.SuperiorID,
	}

	if err := wc.service.UpdateWorkInfo(userId, &workInfo); err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("WorkInfo updated successfully",
		zap.String("userId", userId),
		zap.String("journey", "updateWorkInfo"))

	c.Status(http.StatusOK)
}
