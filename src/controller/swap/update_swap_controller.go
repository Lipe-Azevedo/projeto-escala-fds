package swap

import (
	"net/http"
	"time"

	"github.com/Lipe-Azevedo/escala-fds/src/configuration/logger"
	"github.com/Lipe-Azevedo/escala-fds/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/escala-fds/src/configuration/validation"

	swap_request_dto "github.com/Lipe-Azevedo/escala-fds/src/controller/swap/request"
	"github.com/Lipe-Azevedo/escala-fds/src/model/domain"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (sc *swapControllerInterface) UpdateSwapStatus(c *gin.Context) {
	logger.Info("Init UpdateSwapStatus controller",
		zap.String("journey", "updateSwapStatus"))

	swapID := c.Param("id")
	if _, err := primitive.ObjectIDFromHex(swapID); err != nil {
		logger.Error("Invalid swap ID format in UpdateSwapStatus controller", err,
			zap.String("journey", "updateSwapStatus"),
			zap.String("swapID", swapID))
		restErrVal := rest_err.NewBadRequestError("Invalid Swap ID format, must be a hex value.")
		c.JSON(restErrVal.Code, restErrVal)
		return
	}

	var statusRequest swap_request_dto.SwapRequest
	if err := c.ShouldBindJSON(&statusRequest); err != nil {
		logger.Error("Error validating status update request data", err,
			zap.String("journey", "updateSwapStatus"))
		restErrVal := validation.ValidateUserError(err)
		c.JSON(restErrVal.Code, restErrVal)
		return
	}

	if statusRequest.Status == "" {
		logger.Error("Missing status field for swap status update", nil,
			zap.String("journey", "updateSwapStatus"))
		restErrVal := rest_err.NewBadRequestError("Missing status field for swap status update")
		c.JSON(restErrVal.Code, restErrVal)
		return
	}

	newStatus := domain.SwapStatus(statusRequest.Status)
	if newStatus != domain.StatusApproved && newStatus != domain.StatusRejected && newStatus != domain.StatusPending {
		logger.Error("Invalid status value for swap status update", nil,
			zap.String("journey", "updateSwapStatus"),
			zap.String("receivedStatus", statusRequest.Status))
		restErrVal := rest_err.NewBadRequestError("Invalid status value. Must be 'approved', 'rejected', or 'pending'.")
		c.JSON(restErrVal.Code, restErrVal)
		return
	}

	approverID := "temp-approver-id"
	if approverID == "" {
		logger.Error("Approver ID not found (simulate JWT)", nil, zap.String("journey", "updateSwapStatus"))
		restErr_ := rest_err.NewUnauthorizedError("Unauthorized: Approver ID not found.")
		c.JSON(restErr_.Code, restErr_)
		return
	}

	updatePayload := domain.NewSwapDomain("", "", "", "", "", "", "")

	updatePayload.SetStatus(newStatus)
	if newStatus == domain.StatusApproved {
		updatePayload.SetApprovedBy(approverID)
		now := time.Now()
		updatePayload.SetApprovedAt(now)
	} else if newStatus == domain.StatusRejected {

	}

	serviceErr := sc.service.UpdateSwapServices(swapID, updatePayload)
	if serviceErr != nil {
		logger.Error("Failed to call swap status update service", serviceErr,
			zap.String("journey", "updateSwapStatus"),
			zap.String("swapID", swapID))
		c.JSON(serviceErr.Code, serviceErr)
		return
	}

	logger.Info("UpdateSwapStatus controller executed successfully",
		zap.String("swapID", swapID),
		zap.String("newStatus", string(newStatus)),
		zap.String("journey", "updateSwapStatus"))

	c.Status(http.StatusOK)
}
