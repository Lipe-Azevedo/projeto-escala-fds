package comment

import (
	"net/http"
	"time"

	"github.com/Lipe-Azevedo/escala-fds/src/configuration/logger"
	"github.com/Lipe-Azevedo/escala-fds/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/escala-fds/src/controller/comment/response" // Usar o DTO de response
	"github.com/Lipe-Azevedo/escala-fds/src/model/domain"
	"github.com/Lipe-Azevedo/escala-fds/src/view"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (cc *commentControllerInterface) FindCommentByID(c *gin.Context) {
	logger.Info("Init FindCommentByID controller", zap.String("journey", "findCommentByID"))
	commentID := c.Param("commentId")

	if _, err := primitive.ObjectIDFromHex(commentID); err != nil {
		logger.Error("Invalid comment ID format", err, zap.String("commentID", commentID), zap.String("journey", "findCommentByID"))
		c.JSON(http.StatusBadRequest, rest_err.NewBadRequestError("Invalid comment ID format"))
		return
	}

	// CORREÇÃO: Obter valores do contexto corretamente
	requestingUserIDClaim, exists := c.Get("userID")
	if !exists {
		restErr := rest_err.NewInternalServerError("User ID not found in context")
		logger.Error(restErr.Message, nil, zap.String("journey", "findCommentByID"))
		c.JSON(restErr.Code, restErr)
		return
	}
	requestingUserID := requestingUserIDClaim.(string)

	requestingUserTypeClaim, exists := c.Get("userType")
	if !exists {
		restErr := rest_err.NewInternalServerError("User type not found in context")
		logger.Error(restErr.Message, nil, zap.String("journey", "findCommentByID"))
		c.JSON(restErr.Code, restErr)
		return
	}
	requestingUserType := requestingUserTypeClaim.(domain.UserType)

	// Lógica de permissão: Por enquanto, apenas 'master' pode ver qualquer comentário.
	// Adapte se colaboradores puderem ver comentários sobre eles mesmos.
	if requestingUserType != domain.UserTypeMaster {
		logger.Warn("User does not have permission to view this comment",
			zap.String("journey", "findCommentByID"),
			zap.String("requestingUserID", requestingUserID),
			zap.String("requestingUserType", string(requestingUserType)))
		c.JSON(http.StatusForbidden, rest_err.NewForbiddenError("You do not have permission to view this comment."))
		return
	}

	commentDomain, serviceErr := cc.service.FindCommentByIDService(commentID)
	if serviceErr != nil {
		c.JSON(serviceErr.Code, serviceErr)
		return
	}

	logger.Info("FindCommentByID controller executed successfully",
		zap.String("commentId", commentDomain.GetID()),
		zap.String("journey", "findCommentByID"))
	c.JSON(http.StatusOK, view.ConvertCommentDomainToResponse(commentDomain))
}

func (cc *commentControllerInterface) FindCommentsByCollaboratorAndDate(c *gin.Context) {
	logger.Info("Init FindCommentsByCollaboratorAndDate controller", zap.String("journey", "findCommentsByCollaboratorAndDate"))

	collaboratorID := c.Param("collaboratorId")
	dateString := c.Param("dateString") // Formato YYYY-MM-DD

	if collaboratorID == "" {
		c.JSON(http.StatusBadRequest, rest_err.NewBadRequestError("Collaborator ID is required in path"))
		return
	}

	date, err := time.Parse("2006-01-02", dateString)
	if err != nil {
		logger.Error("Invalid date format for FindCommentsByCollaboratorAndDate", err, zap.String("dateString", dateString), zap.String("journey", "findCommentsByCollaboratorAndDate"))
		c.JSON(http.StatusBadRequest, rest_err.NewBadRequestError("Invalid date format, use YYYY-MM-DD"))
		return
	}

	// CORREÇÃO: Obter valores do contexto corretamente
	requestingUserIDClaim, exists := c.Get("userID")
	if !exists {
		restErr := rest_err.NewInternalServerError("User ID not found in context")
		logger.Error(restErr.Message, nil, zap.String("journey", "findCommentsByCollaboratorAndDate"))
		c.JSON(restErr.Code, restErr)
		return
	}
	requestingUserID := requestingUserIDClaim.(string)

	requestingUserTypeClaim, exists := c.Get("userType")
	if !exists {
		restErr := rest_err.NewInternalServerError("User type not found in context")
		logger.Error(restErr.Message, nil, zap.String("journey", "findCommentsByCollaboratorAndDate"))
		c.JSON(restErr.Code, restErr)
		return
	}
	requestingUserType := requestingUserTypeClaim.(domain.UserType)

	if requestingUserType != domain.UserTypeMaster {
		logger.Warn("User does not have permission to view these comments",
			zap.String("journey", "findCommentsByCollaboratorAndDate"),
			zap.String("requestingUserID", requestingUserID),
			zap.String("requestingUserType", string(requestingUserType)))
		c.JSON(http.StatusForbidden, rest_err.NewForbiddenError("You do not have permission to view these comments."))
		return
	}

	commentDomains, serviceErr := cc.service.FindCommentsByCollaboratorAndDateService(collaboratorID, date, requestingUserID, requestingUserType)
	if serviceErr != nil {
		c.JSON(serviceErr.Code, serviceErr)
		return
	}

	var commentResponses []response.CommentResponse // Usar o DTO específico
	for _, cd := range commentDomains {
		commentResponses = append(commentResponses, view.ConvertCommentDomainToResponse(cd))
	}
	logger.Info("FindCommentsByCollaboratorAndDate controller executed successfully",
		zap.String("collaboratorId", collaboratorID),
		zap.String("date", dateString),
		zap.Int("count", len(commentResponses)),
		zap.String("journey", "findCommentsByCollaboratorAndDate"))
	c.JSON(http.StatusOK, commentResponses)
}

func (cc *commentControllerInterface) FindCommentsByCollaboratorForDateRange(c *gin.Context) {
	logger.Info("Init FindCommentsByCollaboratorForDateRange controller", zap.String("journey", "findCommentsByCollaboratorForDateRange"))

	collaboratorID := c.Param("collaboratorId")
	startDateStr := c.Query("startDate") // YYYY-MM-DD
	endDateStr := c.Query("endDate")     // YYYY-MM-DD

	if collaboratorID == "" {
		c.JSON(http.StatusBadRequest, rest_err.NewBadRequestError("Collaborator ID is required in path"))
		return
	}
	if startDateStr == "" || endDateStr == "" {
		c.JSON(http.StatusBadRequest, rest_err.NewBadRequestError("Both startDate and endDate are required query parameters (YYYY-MM-DD)"))
		return
	}

	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		logger.Error("Invalid startDate format", err, zap.String("startDate", startDateStr), zap.String("journey", "findCommentsByCollaboratorForDateRange"))
		c.JSON(http.StatusBadRequest, rest_err.NewBadRequestError("Invalid startDate format, use YYYY-MM-DD"))
		return
	}
	endDate, err := time.Parse("2006-01-02", endDateStr)
	if err != nil {
		logger.Error("Invalid endDate format", err, zap.String("endDate", endDateStr), zap.String("journey", "findCommentsByCollaboratorForDateRange"))
		c.JSON(http.StatusBadRequest, rest_err.NewBadRequestError("Invalid endDate format, use YYYY-MM-DD"))
		return
	}

	// CORREÇÃO: Obter valores do contexto corretamente
	requestingUserIDClaim, exists := c.Get("userID")
	if !exists {
		restErr := rest_err.NewInternalServerError("User ID not found in context")
		logger.Error(restErr.Message, nil, zap.String("journey", "findCommentsByCollaboratorForDateRange"))
		c.JSON(restErr.Code, restErr)
		return
	}
	requestingUserID := requestingUserIDClaim.(string)

	requestingUserTypeClaim, exists := c.Get("userType")
	if !exists {
		restErr := rest_err.NewInternalServerError("User type not found in context")
		logger.Error(restErr.Message, nil, zap.String("journey", "findCommentsByCollaboratorForDateRange"))
		c.JSON(restErr.Code, restErr)
		return
	}
	requestingUserType := requestingUserTypeClaim.(domain.UserType)

	if requestingUserType != domain.UserTypeMaster {
		logger.Warn("User does not have permission to view these comments",
			zap.String("journey", "findCommentsByCollaboratorForDateRange"),
			zap.String("requestingUserID", requestingUserID),
			zap.String("requestingUserType", string(requestingUserType)))
		c.JSON(http.StatusForbidden, rest_err.NewForbiddenError("You do not have permission to view these comments."))
		return
	}

	commentDomains, serviceErr := cc.service.FindCommentsByCollaboratorForDateRangeService(collaboratorID, startDate, endDate, requestingUserID, requestingUserType)
	if serviceErr != nil {
		c.JSON(serviceErr.Code, serviceErr)
		return
	}
	var commentResponses []response.CommentResponse // Usar o DTO específico
	for _, cd := range commentDomains {
		commentResponses = append(commentResponses, view.ConvertCommentDomainToResponse(cd))
	}
	logger.Info("FindCommentsByCollaboratorForDateRange controller executed successfully",
		zap.String("collaboratorId", collaboratorID),
		zap.String("startDate", startDateStr),
		zap.String("endDate", endDateStr),
		zap.Int("count", len(commentResponses)),
		zap.String("journey", "findCommentsByCollaboratorForDateRange"))
	c.JSON(http.StatusOK, commentResponses)
}
