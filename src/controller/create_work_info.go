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

func (wc *workInfoControllerInterface) CreateWorkInfo(c *gin.Context) {
	logger.Info("Init CreateWorkInfo controller",
		zap.String("journey", "createWorkInfo"))

	var workInfoRequest request.WorkInfoRequest
	if err := c.ShouldBindJSON(&workInfoRequest); err != nil {
		logger.Error("Error validating work info", err,
			zap.String("journey", "createWorkInfo"))
		errRest := validation.ValidateUserError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	domain := model.NewWorkInfoDomain(
		c.Param("userId"),
		model.Team(workInfoRequest.Team),
		workInfoRequest.Position,
		model.Shift(workInfoRequest.DefaultShift),
		model.Weekday(workInfoRequest.WeekdayOff),
		model.WeekendDayOff(workInfoRequest.WeekendDayOff),
		workInfoRequest.SuperiorID,
	)

	domainResult, err := wc.service.CreateWorkInfoServices(domain)
	if err != nil {
		logger.Error("Error trying to call CreateWorkInfo service",
			err,
			zap.String("journey", "createWorkInfo"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("WorkInfo created successfully",
		zap.String("userId", domainResult.GetUserId()),
		zap.String("journey", "createWorkInfo"))

	c.JSON(http.StatusOK, view.ConvertWorkInfoDomainToResponse(domainResult))
}
