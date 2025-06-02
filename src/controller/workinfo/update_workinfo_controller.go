package workinfo

import (
	"net/http"

	"github.com/Lipe-Azevedo/escala-fds/src/configuration/logger"
	"github.com/Lipe-Azevedo/escala-fds/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/escala-fds/src/configuration/validation"
	"github.com/Lipe-Azevedo/escala-fds/src/controller/workinfo/request"
	"github.com/Lipe-Azevedo/escala-fds/src/model/domain"
	"github.com/Lipe-Azevedo/escala-fds/src/view"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (wc *workInfoControllerInterface) UpdateWorkInfo(c *gin.Context) {
	logger.Info("Init UpdateWorkInfo controller",
		zap.String("journey", "updateWorkInfo"))

	targetUserId := c.Param("userId")
	if targetUserId == "" {
		logger.Error("Target UserID is required in path for UpdateWorkInfo", nil,
			zap.String("journey", "updateWorkInfo"))
		restErr := rest_err.NewBadRequestError("Target UserID is required in the URL path")
		c.JSON(restErr.Code, restErr)
		return
	}

	actingUserIDClaim, exists := c.Get("userID")
	if !exists {
		logger.Error("userID not found in context (middleware error?)", nil, zap.String("journey", "updateWorkInfo"))
		restErr := rest_err.NewInternalServerError("Could not retrieve acting user ID from context")
		c.JSON(restErr.Code, restErr)
		return
	}
	actingUserID := actingUserIDClaim.(string)

	actingUserTypeClaim, exists := c.Get("userType")
	if !exists {
		logger.Error("userType not found in context (middleware error?)", nil, zap.String("journey", "updateWorkInfo"))
		restErr := rest_err.NewInternalServerError("Could not retrieve acting user type from context")
		c.JSON(restErr.Code, restErr)
		return
	}
	actingUserType := actingUserTypeClaim.(domain.UserType)

	if actingUserType != domain.UserTypeMaster {
		logger.Warn("Forbidden attempt to update work info by non-master user",
			zap.String("journey", "updateWorkInfo"),
			zap.String("actingUserID", actingUserID),
			zap.String("actingUserType", string(actingUserType)),
			zap.String("targetUserId", targetUserId))
		restErr := rest_err.NewForbiddenError("You do not have permission to perform this action.")
		c.JSON(restErr.Code, restErr)
		return
	}
	logger.Info("UpdateWorkInfo action authorized for master user",
		zap.String("journey", "updateWorkInfo"),
		zap.String("actingUserID", actingUserID),
		zap.String("targetUserId", targetUserId))

	var workInfoUpdateReq request.WorkInfoUpdateRequest
	if err := c.ShouldBindJSON(&workInfoUpdateReq); err != nil {
		logger.Error("Error validating work info update request data", err,
			zap.String("journey", "updateWorkInfo"),
			zap.String("targetUserId", targetUserId))
		errRest := validation.ValidateUserError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	updatedWorkInfoDomain, serviceErr := wc.service.UpdateWorkInfoServices(targetUserId, workInfoUpdateReq)
	if serviceErr != nil {
		c.JSON(serviceErr.Code, serviceErr)
		return
	}

	logger.Info("UpdateWorkInfo controller executed successfully",
		zap.String("targetUserId", targetUserId),
		zap.String("journey", "updateWorkInfo"))

	c.JSON(http.StatusOK, view.ConvertWorkInfoDomainToResponse(updatedWorkInfoDomain))
}
