package controller

import (
	"net/http"

	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/logger"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/validation"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/controller/model/request" // Importa o pacote de request
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/view"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (wc *workInfoControllerInterface) UpdateWorkInfo(c *gin.Context) {
	logger.Info("Init updateWorkInfo controller (PUT handling partial updates)",
		zap.String("journey", "updateWorkInfo"))

	// TODO: Reativar/Confirmar verificação de permissão após implementar JWT/AuthN.
	/*
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

	targetUserId := c.Param("userId")
	if targetUserId == "" {
		logger.Error("Target UserID is required for WorkInfo update", nil,
			zap.String("journey", "updateWorkInfo"))
		restErr := rest_err.NewBadRequestError("Target UserID is required")
		c.JSON(restErr.Code, restErr)
		return
	}

	var workInfoUpdateReq request.WorkInfoUpdateRequest // Usa WorkInfoUpdateRequest (com ponteiros)
	if err := c.ShouldBindJSON(&workInfoUpdateReq); err != nil {
		logger.Error("Error validating work info update request data (WorkInfoUpdateRequest)", // Especifica qual request
			err,
			zap.String("journey", "updateWorkInfo"))
		errRest := validation.ValidateUserError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	// Chama o serviço UpdateWorkInfoServices, que agora espera request.WorkInfoUpdateRequest
	// e implementa a lógica de atualização parcial.
	updatedWorkInfoDomain, serviceErr := wc.service.UpdateWorkInfoServices(targetUserId, workInfoUpdateReq)
	if serviceErr != nil {
		logger.Error("Error calling UpdateWorkInfoServices for update", serviceErr,
			zap.String("journey", "updateWorkInfo"),
			zap.String("targetUserId", targetUserId))
		c.JSON(serviceErr.Code, serviceErr)
		return
	}

	logger.Info("UpdateWorkInfo controller (handling partial update) executed successfully",
		zap.String("targetUserId", targetUserId),
		zap.String("journey", "updateWorkInfo"))

	c.JSON(http.StatusOK, view.ConvertWorkInfoDomainToResponse(updatedWorkInfoDomain))
}
