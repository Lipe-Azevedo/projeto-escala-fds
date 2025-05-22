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

	// TODO: Reativar/Confirmar verificação de permissão após implementar JWT/AuthN.
	// A lógica abaixo assume que c.GetString("userType") e c.GetString("userID")
	// são preenchidos por um middleware de autenticação.
	// Se JWT não estiver implementado, c.GetString("userType") retornará ""
	// e a condição abaixo bloqueará a ação. Comente o bloco if para testes sem JWT.

	/* // BLOCO DE PERMISSÃO TEMPORARIAMENTE COMENTADO PARA TESTES SEM JWT
	actingUserType := c.GetString("userType")
	if model.UserType(actingUserType) != model.UserTypeMaster {
		logger.Warn("Forbidden attempt to create work info by non-master user (or auth not implemented)",
			zap.String("journey", "createWorkInfo"),
			zap.String("actingUserID", c.GetString("userID")),
			zap.String("actingUserType", actingUserType))
		restErr := rest_err.NewForbiddenError("You do not have permission to perform this action or user type not identified")
		c.JSON(restErr.Code, restErr)
		return
	}
	*/

	var workInfoRequest request.WorkInfoRequest
	if err := c.ShouldBindJSON(&workInfoRequest); err != nil {
		logger.Error("Error validating work info request data", err,
			zap.String("journey", "createWorkInfo"))
		errRest := validation.ValidateUserError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	targetUserId := c.Param("userId")

	domain := model.NewWorkInfoDomain(
		targetUserId,
		model.Team(workInfoRequest.Team),
		workInfoRequest.Position,
		model.Shift(workInfoRequest.DefaultShift),
		model.Weekday(workInfoRequest.WeekdayOff),
		model.WeekendDayOff(workInfoRequest.WeekendDayOff),
		workInfoRequest.SuperiorID,
	)

	domainResult, serviceErr := wc.service.CreateWorkInfoServices(domain)
	if serviceErr != nil {
		logger.Error("Error trying to call CreateWorkInfo service",
			serviceErr,
			zap.String("journey", "createWorkInfo"))
		c.JSON(serviceErr.Code, serviceErr)
		return
	}

	logger.Info("WorkInfo created successfully",
		zap.String("targetUserId", domainResult.GetUserId()),
		zap.String("journey", "createWorkInfo"))

	c.JSON(http.StatusCreated, view.ConvertWorkInfoDomainToResponse(domainResult))
}
