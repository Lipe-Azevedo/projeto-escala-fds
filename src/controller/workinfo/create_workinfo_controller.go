package workinfo

import (
	"net/http"

	"github.com/Lipe-Azevedo/escala-fds/src/configuration/logger"
	"github.com/Lipe-Azevedo/escala-fds/src/configuration/rest_err" // Adicionado para NewForbiddenError
	"github.com/Lipe-Azevedo/escala-fds/src/configuration/validation"

	// Import para o DTO de request de workinfo, usando alias
	workinfo_request_dto "github.com/Lipe-Azevedo/escala-fds/src/controller/workinfo/request"
	"github.com/Lipe-Azevedo/escala-fds/src/model" // Para model.UserTypeMaster e model.Team, etc.
	"github.com/Lipe-Azevedo/escala-fds/src/view"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (wc *workInfoControllerInterface) CreateWorkInfo(c *gin.Context) {
	logger.Info("Init CreateWorkInfo controller",
		zap.String("journey", "createWorkInfo"))

	// TODO: (Pós-JWT) A verificação de permissão será feita pelo middleware JWT.
	// O middleware deverá injetar 'userID' e 'userType' no contexto do Gin (c.GetString).
	// Exemplo de como seria a lógica de permissão (DESCOMENTE E ADAPTE QUANDO JWT ESTIVER PRONTO):
	/*
		actingUserType := c.GetString("userType") // Vem do token JWT
		actingUserID := c.GetString("userID")     // Vem do token JWT

		if model.UserType(actingUserType) != model.UserTypeMaster {
			logger.Warn("Forbidden attempt to create work info by non-master user",
				zap.String("journey", "createWorkInfo"),
				zap.String("actingUserID", actingUserID),
				zap.String("actingUserType", actingUserType))
			restErr := rest_err.NewForbiddenError("You do not have permission to perform this action.")
			c.JSON(restErr.Code, restErr)
			return
		}
		logger.Info("CreateWorkInfo action performed by master user", zap.String("actingUserID", actingUserID))
	*/

	var workInfoRequest workinfo_request_dto.WorkInfoRequest // Usando DTO específico
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

	domain := model.NewWorkInfoDomain(
		targetUserId, // ID do usuário para o qual o WorkInfo está sendo criado
		model.Team(workInfoRequest.Team),
		workInfoRequest.Position,
		model.Shift(workInfoRequest.DefaultShift),
		model.Weekday(workInfoRequest.WeekdayOff),
		model.WeekendDayOff(workInfoRequest.WeekendDayOff),
		workInfoRequest.SuperiorID,
	)

	domainResult, serviceErr := wc.service.CreateWorkInfoServices(domain)
	if serviceErr != nil {
		logger.Error("Failed to call workinfo creation service", serviceErr,
			zap.String("journey", "createWorkInfo"),
			zap.String("targetUserId", targetUserId))
		c.JSON(serviceErr.Code, serviceErr)
		return
	}

	logger.Info("WorkInfo created successfully via controller",
		zap.String("targetUserId", domainResult.GetUserId()),
		zap.String("journey", "createWorkInfo"))

	// view.ConvertWorkInfoDomainToResponse será ajustado para usar workinfo_response_dto
	c.JSON(http.StatusCreated, view.ConvertWorkInfoDomainToResponse(domainResult))
}
