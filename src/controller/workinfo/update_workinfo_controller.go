package workinfo

import (
	"net/http"

	"github.com/Lipe-Azevedo/escala-fds/src/configuration/logger"
	"github.com/Lipe-Azevedo/escala-fds/src/configuration/rest_err" // Adicionado
	"github.com/Lipe-Azevedo/escala-fds/src/configuration/validation"

	// Import para o DTO de request de workinfo, usando alias
	workinfo_request_dto "github.com/Lipe-Azevedo/escala-fds/src/controller/workinfo/request" // Para model.UserTypeMaster
	"github.com/Lipe-Azevedo/escala-fds/src/view"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (wc *workInfoControllerInterface) UpdateWorkInfo(c *gin.Context) {
	logger.Info("Init UpdateWorkInfo controller (PUT handling partial updates)",
		zap.String("journey", "updateWorkInfo"))

	// TODO: (Pós-JWT) A verificação de permissão será feita pelo middleware JWT.
	// Somente 'master' pode atualizar WorkInfo.
	/*
		actingUserType := c.GetString("userType") // Vem do token JWT
		actingUserID := c.GetString("userID")     // Vem do token JWT

		if model.UserType(actingUserType) != model.UserTypeMaster {
			logger.Warn("Forbidden attempt to update work info by non-master user",
				zap.String("journey", "updateWorkInfo"),
				zap.String("actingUserID", actingUserID),
				zap.String("actingUserType", actingUserType))
			restErr := rest_err.NewForbiddenError("You do not have permission to perform this action.")
			c.JSON(restErr.Code, restErr)
			return
		}
		logger.Info("UpdateWorkInfo action performed by master user", zap.String("actingUserID", actingUserID))
	*/

	targetUserId := c.Param("userId")
	if targetUserId == "" {
		logger.Error("Target UserID is required in path for UpdateWorkInfo", nil,
			zap.String("journey", "updateWorkInfo"))
		restErr := rest_err.NewBadRequestError("Target UserID is required in the URL path")
		c.JSON(restErr.Code, restErr)
		return
	}

	var workInfoUpdateReq workinfo_request_dto.WorkInfoUpdateRequest // Usando DTO específico
	if err := c.ShouldBindJSON(&workInfoUpdateReq); err != nil {
		logger.Error("Error validating work info update request data", err,
			zap.String("journey", "updateWorkInfo"))
		errRest := validation.ValidateUserError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	// O serviço UpdateWorkInfoServices espera workinfo_request_dto.WorkInfoUpdateRequest
	updatedWorkInfoDomain, serviceErr := wc.service.UpdateWorkInfoServices(targetUserId, workInfoUpdateReq)
	if serviceErr != nil {
		logger.Error("Failed to call workinfo update service", serviceErr,
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
