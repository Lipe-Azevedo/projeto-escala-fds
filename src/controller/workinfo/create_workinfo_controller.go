package workinfo

import (
	"net/http"

	"github.com/Lipe-Azevedo/escala-fds/src/configuration/logger"
	"github.com/Lipe-Azevedo/escala-fds/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/escala-fds/src/configuration/validation"
	workinfo_request_dto "github.com/Lipe-Azevedo/escala-fds/src/controller/workinfo/request"
	"github.com/Lipe-Azevedo/escala-fds/src/model/domain"
	"github.com/Lipe-Azevedo/escala-fds/src/view"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (wc *workInfoControllerInterface) CreateWorkInfo(c *gin.Context) {
	logger.Info("Init CreateWorkInfo controller",
		zap.String("journey", "createWorkInfo"))

	actingUserTypeClaim, exists := c.Get("userType")
	if !exists {
		logger.Error("userType not found in context (middleware error?)", nil, zap.String("journey", "createWorkInfo"))
		restErr := rest_err.NewInternalServerError("Could not retrieve user type from context")
		c.JSON(restErr.Code, restErr)
		return
	}
	actingUserType, okAssertionUserType := actingUserTypeClaim.(domain.UserType)
	if !okAssertionUserType {
		logger.Error("userType in context is not of type domain.UserType", nil,
			zap.String("journey", "createWorkInfo"),
			zap.Any("retrievedType", actingUserTypeClaim))
		restErr := rest_err.NewInternalServerError("Invalid user type format in context")
		c.JSON(restErr.Code, restErr)
		return
	}

	actingUserIDClaim, exists := c.Get("userID")
	if !exists {
		logger.Error("userID not found in context (middleware error?)", nil, zap.String("journey", "createWorkInfo"))
		restErr := rest_err.NewInternalServerError("Could not retrieve user ID from context")
		c.JSON(restErr.Code, restErr)
		return
	}
	actingUserID, okAssertionUserID := actingUserIDClaim.(string)
	if !okAssertionUserID {
		logger.Error("userID in context is not of type string", nil,
			zap.String("journey", "createWorkInfo"),
			zap.Any("retrievedType", actingUserIDClaim))
		restErr := rest_err.NewInternalServerError("Invalid user ID format in context")
		c.JSON(restErr.Code, restErr)
		return
	}

	logger.Info("Verificando permissão para CreateWorkInfo",
		zap.String("journey", "createWorkInfo"),
		zap.String("actingUserID_from_token", actingUserID),
		zap.Any("actingUserTypeClaim_from_context_raw", actingUserTypeClaim),
		zap.Bool("typeAssertion_userType_ok", okAssertionUserType),
		zap.String("asserted_actingUserType_as_string", string(actingUserType)),
		zap.String("expected_master_type_constant", string(domain.UserTypeMaster)))

	if actingUserType != domain.UserTypeMaster {
		logger.Warn("Forbidden attempt to create work info by non-master user",
			zap.String("journey", "createWorkInfo"),
			zap.String("actingUserID", actingUserID),
			zap.String("actingUserType_evaluated", string(actingUserType)))
		restErr := rest_err.NewForbiddenError("You do not have permission to perform this action.")
		c.JSON(restErr.Code, restErr)
		return
	}
	logger.Info("CreateWorkInfo action authorized for master user",
		zap.String("actingUserID", actingUserID),
		zap.String("journey", "createWorkInfo"))

	var workInfoRequest workinfo_request_dto.WorkInfoRequest
	if err := c.ShouldBindJSON(&workInfoRequest); err != nil {
		logger.Error("Error validating work info request data for creation", err,
			zap.String("journey", "createWorkInfo"))
		errRest := validation.ValidateUserError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	targetUserId := c.Param("userId")
	if targetUserId == "" {
		logger.Error("Target UserID is required in path for CreateWorkInfo", nil,
			zap.String("journey", "createWorkInfo"))
		restErr := rest_err.NewBadRequestError("Target UserID is required in the URL path")
		c.JSON(restErr.Code, restErr)
		return
	}

	domainInstance := domain.NewWorkInfoDomain(
		targetUserId,
		domain.Team(workInfoRequest.Team),
		workInfoRequest.Position,
		domain.Shift(workInfoRequest.DefaultShift),
		domain.Weekday(workInfoRequest.WeekdayOff),
		domain.WeekendDayOff(workInfoRequest.WeekendDayOff),
		workInfoRequest.SuperiorID,
	)

	domainResult, serviceErr := wc.service.CreateWorkInfoServices(domainInstance)
	if serviceErr != nil {
		c.JSON(serviceErr.Code, serviceErr)
		return
	}

	logger.Info("WorkInfo created successfully via controller",
		zap.String("targetUserIdForResult", domainResult.GetUserId()),
		zap.String("journey", "createWorkInfo"))

	c.JSON(http.StatusCreated, view.ConvertWorkInfoDomainToResponse(domainResult))
}
