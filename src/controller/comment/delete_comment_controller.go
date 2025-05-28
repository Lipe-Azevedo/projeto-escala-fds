package comment

import (
	"net/http"

	"github.com/Lipe-Azevedo/escala-fds/src/configuration/logger"
	"github.com/Lipe-Azevedo/escala-fds/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/escala-fds/src/model/domain"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (cc *commentControllerInterface) DeleteComment(c *gin.Context) {
	logger.Info("Init DeleteComment controller", zap.String("journey", "deleteComment"))
	commentID := c.Param("commentId")

	if _, err := primitive.ObjectIDFromHex(commentID); err != nil {
		logger.Error("Invalid comment ID format for delete", err, zap.String("commentID", commentID), zap.String("journey", "deleteComment"))
		c.JSON(http.StatusBadRequest, rest_err.NewBadRequestError("Invalid comment ID format"))
		return
	}

	// CORREÇÃO: Obter valores do contexto corretamente
	requestingUserIDClaim, exists := c.Get("userID")
	if !exists {
		restErr := rest_err.NewInternalServerError("User ID not found in context")
		logger.Error(restErr.Message, nil, zap.String("journey", "deleteComment"))
		c.JSON(restErr.Code, restErr)
		return
	}
	requestingUserID := requestingUserIDClaim.(string)

	requestingUserTypeClaim, exists := c.Get("userType")
	if !exists {
		restErr := rest_err.NewInternalServerError("User type not found in context")
		logger.Error(restErr.Message, nil, zap.String("journey", "deleteComment"))
		c.JSON(restErr.Code, restErr)
		return
	}
	requestingUserType := requestingUserTypeClaim.(domain.UserType)

	// Permissão: Apenas Master pode deletar
	if requestingUserType != domain.UserTypeMaster {
		logger.Warn("User does not have permission to delete this comment",
			zap.String("journey", "deleteComment"),
			zap.String("requestingUserID", requestingUserID),
			zap.String("requestingUserType", string(requestingUserType)))
		c.JSON(http.StatusForbidden, rest_err.NewForbiddenError("You do not have permission to delete this comment."))
		return
	}

	serviceErr := cc.service.DeleteCommentService(commentID, requestingUserID, requestingUserType)
	if serviceErr != nil {
		c.JSON(serviceErr.Code, serviceErr)
		return
	}

	logger.Info("DeleteComment controller executed successfully",
		zap.String("commentId", commentID),
		zap.String("journey", "deleteComment"))
	c.JSON(http.StatusOK, gin.H{"message": "Comment deleted successfully"})
}
