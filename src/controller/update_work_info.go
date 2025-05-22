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

	// TODO: Reativar/Confirmar verificação de permissão após implementar JWT/AuthN.
	// A lógica abaixo assume que c.GetString("userType") e c.GetString("userID")
	// são preenchidos por um middleware de autenticação.
	// Se JWT não estiver implementado, c.GetString("userType") retornará ""
	// e a condição abaixo bloqueará a ação. Comente o bloco if para testes sem JWT.

	/* // BLOCO DE PERMISSÃO TEMPORARIAMENTE COMENTADO PARA TESTES SEM JWT
	actingUserType := c.GetString("userType")
	if model.UserType(actingUserType) != model.UserTypeMaster {
		logger.Warn("Forbidden attempt to update work info by non-master user (or auth not implemented)",
			zap.String("journey", "updateWorkInfo"),
			zap.String("actingUserID", c.GetString("userID")),
			zap.String("actingUserType", actingUserType))
		restErr := rest_err.NewForbiddenError("You do not have permission to perform this action or user type not identified")
		c.JSON(restErr.Code, restErr)
		return
	}
	*/

	var workInfoRequest request.WorkInfoRequest
	if err := c.ShouldBindJSON(&workInfoRequest); err != nil {
		logger.Error("Error validating work info update request data",
			err,
			zap.String("journey", "updateWorkInfo"))
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

	serviceErr := wc.service.UpdateWorkInfoServices(targetUserId, domain)
	if serviceErr != nil {
		logger.Error("Error trying to call updateWorkInfo service",
			serviceErr,
			zap.String("journey", "updateWorkInfo"))
		c.JSON(serviceErr.Code, serviceErr)
		return
	}

	logger.Info("UpdateWorkInfo controller executed successfully",
		zap.String("targetUserId", targetUserId),
		zap.String("journey", "updateWorkInfo"))

	c.Status(http.StatusOK)
}
