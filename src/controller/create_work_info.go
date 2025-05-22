package controller

import (
	"net/http"

	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/logger"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/rest_err"
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

	// Verificação de Permissão: Somente usuários master podem criar work info
	actingUserType := c.GetString("userType") // Assume que isso é definido pelo middleware de autenticação
	if model.UserType(actingUserType) != model.UserTypeMaster {
		logger.Warn("Forbidden attempt to create work info by non-master user", // Mudado para Warn
			zap.String("journey", "createWorkInfo"),
			zap.String("actingUserID", c.GetString("userID")), // Adicionado ID do usuário que tentou a ação
			zap.String("actingUserType", actingUserType))
		restErr := rest_err.NewForbiddenError("You do not have permission to perform this action")
		c.JSON(restErr.Code, restErr)
		return
	}

	var workInfoRequest request.WorkInfoRequest
	if err := c.ShouldBindJSON(&workInfoRequest); err != nil {
		logger.Error("Error validating work info request data", err, // Mensagem de log mais específica
			zap.String("journey", "createWorkInfo"))
		errRest := validation.ValidateUserError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	// targetUserId é o ID do usuário para o qual a work info está sendo criada.
	// Este usuário deve ser, tipicamente, um 'colaborador'.
	// Validações adicionais podem ser adicionadas aqui ou na camada de serviço
	// para garantir que o usuário alvo exista e seja um colaborador.
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
			serviceErr, // serviceErr já é *rest_err.RestErr
			zap.String("journey", "createWorkInfo"))
		c.JSON(serviceErr.Code, serviceErr)
		return
	}

	logger.Info("WorkInfo created successfully",
		zap.String("targetUserId", domainResult.GetUserId()),
		zap.String("journey", "createWorkInfo"))

	c.JSON(http.StatusCreated, view.ConvertWorkInfoDomainToResponse(domainResult)) // Alterado para http.StatusCreated
}
