package swap

import (
	"net/http"

	"github.com/Lipe-Azevedo/escala-fds/src/configuration/logger"
	"github.com/Lipe-Azevedo/escala-fds/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/escala-fds/src/model/domain"
	"github.com/Lipe-Azevedo/escala-fds/src/view"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (sc *swapControllerInterface) FindSwapByID(c *gin.Context) {
	logger.Info("Init FindSwapByID controller",
		zap.String("journey", "findSwapByID"))

	swapID := c.Param("id")
	if _, err := primitive.ObjectIDFromHex(swapID); err != nil {
		logger.Error("Invalid swap ID format in FindSwapByID controller", err,
			zap.String("journey", "findSwapByID"),
			zap.String("swapID", swapID))
		restErrVal := rest_err.NewBadRequestError("Invalid Swap ID format, must be a hex value.")
		c.JSON(restErrVal.Code, restErrVal)
		return
	}

	actingUserIDClaim, exists := c.Get("userID")
	if !exists {
		restErr := rest_err.NewInternalServerError("User ID not found in context")
		logger.Error(restErr.Message, nil, zap.String("journey", "findSwapByID"))
		c.JSON(restErr.Code, restErr)
		return
	}
	actingUserID := actingUserIDClaim.(string)

	actingUserTypeClaim, exists := c.Get("userType")
	if !exists {
		restErr := rest_err.NewInternalServerError("User type not found in context")
		logger.Error(restErr.Message, nil, zap.String("journey", "findSwapByID"))
		c.JSON(restErr.Code, restErr)
		return
	}
	actingUserType := actingUserTypeClaim.(domain.UserType)

	swapDomain, serviceErr := sc.service.FindSwapByIDServices(swapID)
	if serviceErr != nil {

		c.JSON(serviceErr.Code, serviceErr)
		return
	}

	isParticipant := (swapDomain.GetRequesterID() == actingUserID || swapDomain.GetRequestedID() == actingUserID)

	if actingUserType == domain.UserTypeMaster {

		logger.Info("FindSwapByID: Access granted for master.",
			zap.String("journey", "findSwapByID"),
			zap.String("actingUserID", actingUserID),
			zap.String("swapID", swapID))
	} else if actingUserType == domain.UserTypeCollaborator && isParticipant {

		logger.Info("FindSwapByID: Access granted for collaborator participant.",
			zap.String("journey", "findSwapByID"),
			zap.String("actingUserID", actingUserID),
			zap.String("swapID", swapID))
	} else {
		logger.Warn("Forbidden attempt to find swap by non-participant/non-master.",
			zap.String("journey", "findSwapByID"),
			zap.String("actingUserID", actingUserID),
			zap.String("actingUserType", string(actingUserType)),
			zap.String("swapID", swapID))
		restErr := rest_err.NewForbiddenError("You do not have permission to view this swap request.")
		c.JSON(restErr.Code, restErr)
		return
	}

	logger.Info("FindSwapByID controller executed successfully",
		zap.String("swapID", swapID),
		zap.String("journey", "findSwapByID"))

	c.JSON(http.StatusOK, view.ConvertSwapDomainToResponse(swapDomain))
}
