package workinfo

import (
	"net/http"

	"github.com/Lipe-Azevedo/escala-fds/src/configuration/logger"
	"github.com/Lipe-Azevedo/escala-fds/src/configuration/rest_err" // Adicionado para NewForbiddenError

	// Para model.UserTypeMaster
	"github.com/Lipe-Azevedo/escala-fds/src/view"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (wc *workInfoControllerInterface) FindWorkInfoByUserId(c *gin.Context) {
	logger.Info("Init FindWorkInfoByUserId controller",
		zap.String("journey", "findWorkInfoByUserId"))

	targetUserId := c.Param("userId")
	if targetUserId == "" {
		logger.Error("Target UserID is required in path for FindWorkInfoByUserId", nil,
			zap.String("journey", "findWorkInfoByUserId"))
		restErr := rest_err.NewBadRequestError("Target UserID is required in the URL path")
		c.JSON(restErr.Code, restErr)
		return
	}

	// TODO: (Pós-JWT) Lógica de Permissão:
	// - Master pode ver WorkInfo de qualquer um.
	// - Colaborador só pode ver o seu próprio WorkInfo.
	/*
		actingUserID := c.GetString("userID")     // Vem do token JWT
		actingUserType := c.GetString("userType") // Vem do token JWT

		if model.UserType(actingUserType) != model.UserTypeMaster && actingUserID != targetUserId {
			logger.Warn("Forbidden attempt to find work info for another user by non-master/non-self user",
				zap.String("journey", "findWorkInfoByUserId"),
				zap.String("actingUserID", actingUserID),
				zap.String("actingUserType", string(actingUserType)),
				zap.String("targetUserID", targetUserId))
			restErr := rest_err.NewForbiddenError("You do not have permission to view this information.")
			c.JSON(restErr.Code, restErr)
			return
		}
		logger.Info("FindWorkInfoByUserId action allowed",
			zap.String("actingUserID", actingUserID),
			zap.String("actingUserType", string(actingUserType)),
			zap.String("targetUserID", targetUserId))
	*/

	workInfoDomain, serviceErr := wc.service.FindWorkInfoByUserIdServices(targetUserId)
	if serviceErr != nil {
		logger.Error("Failed to call find workinfo by user ID service", serviceErr,
			zap.String("journey", "findWorkInfoByUserId"),
			zap.String("targetUserIdToFind", targetUserId))
		c.JSON(serviceErr.Code, serviceErr)
		return
	}

	logger.Info("FindWorkInfoByUserId controller executed successfully",
		zap.String("foundUserIdForWorkInfo", targetUserId), // Corrigido para targetUserId
		zap.String("journey", "findWorkInfoByUserId"))

	c.JSON(http.StatusOK, view.ConvertWorkInfoDomainToResponse(workInfoDomain))
}
