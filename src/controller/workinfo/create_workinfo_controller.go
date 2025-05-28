package workinfo

import (
	"net/http"

	"github.com/Lipe-Azevedo/escala-fds/src/configuration/logger"
	"github.com/Lipe-Azevedo/escala-fds/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/escala-fds/src/configuration/validation"
	workinfo_request_dto "github.com/Lipe-Azevedo/escala-fds/src/controller/workinfo/request"
	"github.com/Lipe-Azevedo/escala-fds/src/model/domain"
	"github.com/Lipe-Azevedo/escala-fds/src/view"
	"github.com/gin-gonic/gin" // Corrigido para gin-gonic/gin
	"go.uber.org/zap"
)

func (wc *workInfoControllerInterface) CreateWorkInfo(c *gin.Context) {
	logger.Info("Init CreateWorkInfo controller",
		zap.String("journey", "createWorkInfo"))

	// Extrair userType e userID do contexto (injetado pelo AuthMiddleware)
	actingUserTypeClaim, exists := c.Get("userType")
	if !exists {
		logger.Error("userType not found in context", nil, zap.String("journey", "createWorkInfo"))
		restErr := rest_err.NewInternalServerError("Could not retrieve user type from context")
		c.JSON(restErr.Code, restErr)
		return
	}
	// O middleware agora armazena domain.UserType diretamente
	actingUserType, ok := actingUserTypeClaim.(domain.UserType)
	if !ok {
		logger.Error("userType in context is not of type domain.UserType", nil, zap.String("journey", "createWorkInfo"))
		restErr := rest_err.NewInternalServerError("Invalid user type format in context")
		c.JSON(restErr.Code, restErr)
		return
	}

	actingUserIDClaim, exists := c.Get("userID")
	if !exists {
		logger.Error("userID not found in context", nil, zap.String("journey", "createWorkInfo"))
		restErr := rest_err.NewInternalServerError("Could not retrieve user ID from context")
		c.JSON(restErr.Code, restErr)
		return
	}
	actingUserID, ok := actingUserIDClaim.(string)
	if !ok {
		logger.Error("userID in context is not of type string", nil, zap.String("journey", "createWorkInfo"))
		restErr := rest_err.NewInternalServerError("Invalid user ID format in context")
		c.JSON(restErr.Code, restErr)
		return
	}

	// Lógica de Permissão: Somente 'master' pode criar WorkInfo.
	if actingUserType != domain.UserTypeMaster {
		logger.Warn("Forbidden attempt to create work info by non-master user",
			zap.String("journey", "createWorkInfo"),
			zap.String("actingUserID", actingUserID),
			zap.String("actingUserType", string(actingUserType)))
		restErr := rest_err.NewForbiddenError("You do not have permission to perform this action.")
		c.JSON(restErr.Code, restErr)
		return
	}
	logger.Info("CreateWorkInfo action performed by master user",
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
		targetUserId, // ID do usuário para o qual o WorkInfo está sendo criado
		domain.Team(workInfoRequest.Team),
		workInfoRequest.Position,
		domain.Shift(workInfoRequest.DefaultShift),
		domain.Weekday(workInfoRequest.WeekdayOff),
		domain.WeekendDayOff(workInfoRequest.WeekendDayOff),
		workInfoRequest.SuperiorID,
	)

	domainResult, serviceErr := wc.service.CreateWorkInfoServices(domainInstance)
	if serviceErr != nil {
		logger.Error("Failed to call workinfo creation service", serviceErr,
			zap.String("journey", "createWorkInfo"),
			zap.String("targetUserId", targetUserId))
		c.JSON(serviceErr.Code, serviceErr)
		return
	}

	logger.Info("WorkInfo created successfully via controller",
		zap.String("targetUserIdForResult", domainResult.GetUserId()), // Nome do campo alterado para clareza
		zap.String("journey", "createWorkInfo"))

	c.JSON(http.StatusCreated, view.ConvertWorkInfoDomainToResponse(domainResult))
}
