package workinfo

import (
	"net/http"

	"github.com/Lipe-Azevedo/escala-fds/src/configuration/logger"
	"github.com/Lipe-Azevedo/escala-fds/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/escala-fds/src/model/domain"
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

	actingUserIDClaim, exists := c.Get("userID")
	if !exists {
		logger.Error("userID not found in context (middleware error?)", nil, zap.String("journey", "findWorkInfoByUserId"))
		restErr := rest_err.NewInternalServerError("Could not retrieve acting user ID from context")
		c.JSON(restErr.Code, restErr)
		return
	}
	actingUserID := actingUserIDClaim.(string)

	actingUserTypeClaim, exists := c.Get("userType")
	if !exists {
		logger.Error("userType not found in context (middleware error?)", nil, zap.String("journey", "findWorkInfoByUserId"))
		restErr := rest_err.NewInternalServerError("Could not retrieve acting user type from context")
		c.JSON(restErr.Code, restErr)
		return
	}
	actingUserType := actingUserTypeClaim.(domain.UserType)

	if actingUserType == domain.UserTypeMaster {
		logger.Info("FindWorkInfoByUserId: Access granted for master.",
			zap.String("journey", "findWorkInfoByUserId"),
			zap.String("actingUserID", actingUserID),
			zap.String("targetUserId", targetUserId))
	} else if actingUserType == domain.UserTypeCollaborator {
		if actingUserID != targetUserId {
			logger.Warn("Forbidden attempt by collaborator to find work info for another user.",
				zap.String("journey", "findWorkInfoByUserId"),
				zap.String("actingUserID", actingUserID),
				zap.String("targetUserId", targetUserId))
			restErr := rest_err.NewForbiddenError("You do not have permission to view this information.")
			c.JSON(restErr.Code, restErr)
			return
		}
		logger.Info("FindWorkInfoByUserId: Access granted for collaborator to view their own info.",
			zap.String("journey", "findWorkInfoByUserId"),
			zap.String("actingUserID", actingUserID))
	} else {
		logger.Error("FindWorkInfoByUserId: Unknown or unauthorized user type attempting access.", nil,
			zap.String("journey", "findWorkInfoByUserId"),
			zap.String("actingUserID", actingUserID),
			zap.String("actingUserType", string(actingUserType)))
		restErr := rest_err.NewForbiddenError("You do not have permission to perform this action.")
		c.JSON(restErr.Code, restErr)
		return
	}

	workInfoDomain, serviceErr := wc.service.FindWorkInfoByUserIdServices(targetUserId)
	if serviceErr != nil {
		logger.Error("Error calling FindWorkInfoByUserIdServices from controller",
			serviceErr,
			zap.String("journey", "findWorkInfoByUserId"),
			zap.String("targetUserId", targetUserId))
		c.JSON(serviceErr.Code, serviceErr)
		return
	}

	logger.Info("FindWorkInfoByUserId controller executed successfully",
		zap.String("foundUserIdForWorkInfo", workInfoDomain.GetUserId()),
		zap.String("journey", "findWorkInfoByUserId"))

	c.JSON(http.StatusOK, view.ConvertWorkInfoDomainToResponse(workInfoDomain))
}
