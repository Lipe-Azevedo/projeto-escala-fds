package comment

import (
	"net/http"

	"github.com/Lipe-Azevedo/escala-fds/src/configuration/logger"
	"github.com/Lipe-Azevedo/escala-fds/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/escala-fds/src/configuration/validation"
	"github.com/Lipe-Azevedo/escala-fds/src/controller/comment/request"
	"github.com/Lipe-Azevedo/escala-fds/src/model/domain"
	"github.com/Lipe-Azevedo/escala-fds/src/view"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (cc *commentControllerInterface) UpdateComment(c *gin.Context) {
	logger.Info("Init UpdateComment controller", zap.String("journey", "updateComment"))
	commentID := c.Param("commentId")

	if _, err := primitive.ObjectIDFromHex(commentID); err != nil {
		logger.Error("Invalid comment ID format for update", err, zap.String("commentID", commentID), zap.String("journey", "updateComment"))
		c.JSON(http.StatusBadRequest, rest_err.NewBadRequestError("Invalid comment ID format"))
		return
	}

	var commentUpdateReq request.CommentUpdateRequest
	if err := c.ShouldBindJSON(&commentUpdateReq); err != nil {
		logger.Error("Error validating comment update request data", err, zap.String("journey", "updateComment"))
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	// CORREÇÃO: Obter valores do contexto corretamente
	requestingUserIDClaim, exists := c.Get("userID")
	if !exists {
		restErr := rest_err.NewInternalServerError("User ID not found in context")
		logger.Error(restErr.Message, nil, zap.String("journey", "updateComment"))
		c.JSON(restErr.Code, restErr)
		return
	}
	requestingUserID := requestingUserIDClaim.(string)

	requestingUserTypeClaim, exists := c.Get("userType")
	if !exists {
		restErr := rest_err.NewInternalServerError("User type not found in context")
		logger.Error(restErr.Message, nil, zap.String("journey", "updateComment"))
		c.JSON(restErr.Code, restErr)
		return
	}
	requestingUserType := requestingUserTypeClaim.(domain.UserType)

	// Permissão: Apenas Master pode atualizar (ou adicionar lógica para autor original)
	if requestingUserType != domain.UserTypeMaster {
		logger.Warn("User does not have permission to update this comment",
			zap.String("journey", "updateComment"),
			zap.String("requestingUserID", requestingUserID),
			zap.String("requestingUserType", string(requestingUserType)))
		c.JSON(http.StatusForbidden, rest_err.NewForbiddenError("You do not have permission to update this comment."))
		return
	}

	updatedCommentDomain, serviceErr := cc.service.UpdateCommentService(commentID, commentUpdateReq, requestingUserID, requestingUserType)
	if serviceErr != nil {
		c.JSON(serviceErr.Code, serviceErr)
		return
	}

	logger.Info("UpdateComment controller executed successfully",
		zap.String("commentId", updatedCommentDomain.GetID()),
		zap.String("journey", "updateComment"))
	c.JSON(http.StatusOK, view.ConvertCommentDomainToResponse(updatedCommentDomain))
}
