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
	"go.uber.org/zap"
)

// CreateComment cria um novo comentário.
// Requer autenticação JWT e que o usuário seja 'master'.
func (cc *commentControllerInterface) CreateComment(c *gin.Context) {
	logger.Info("Init CreateComment controller", zap.String("journey", "createComment"))

	// Extrair AuthorID e UserType do token JWT
	authorIDClaim, exists := c.Get("userID")
	if !exists {
		errMsg := "Failed to get userID from token"
		logger.Error(errMsg, nil, zap.String("journey", "createComment"))
		c.JSON(http.StatusInternalServerError, rest_err.NewInternalServerError(errMsg))
		return
	}
	authorID := authorIDClaim.(string)

	authorTypeClaim, exists := c.Get("userType")
	if !exists {
		errMsg := "Failed to get userType from token"
		logger.Error(errMsg, nil, zap.String("journey", "createComment"))
		c.JSON(http.StatusInternalServerError, rest_err.NewInternalServerError(errMsg))
		return
	}
	authorType := authorTypeClaim.(domain.UserType)

	// Verificação de Permissão: Somente 'master' pode criar comentários.
	if authorType != domain.UserTypeMaster {
		errMsg := "User does not have permission to create comments"
		logger.Warn(errMsg,
			zap.String("journey", "createComment"),
			zap.String("userID", authorID),
			zap.String("userType", string(authorType)))
		c.JSON(http.StatusForbidden, rest_err.NewForbiddenError(errMsg))
		return
	}

	var commentRequest request.CommentRequest
	if err := c.ShouldBindJSON(&commentRequest); err != nil {
		logger.Error("Error validating comment request data for creation", err,
			zap.String("journey", "createComment"))
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	// Chamar o serviço para criar o comentário
	commentDomain, serviceErr := cc.service.CreateCommentService(commentRequest, authorID)
	if serviceErr != nil {
		// O serviço já deve ter logado o erro.
		c.JSON(serviceErr.Code, serviceErr)
		return
	}

	logger.Info("Comment created successfully via controller",
		zap.String("commentId", commentDomain.GetID()),
		zap.String("journey", "createComment"))

	c.JSON(http.StatusCreated, view.ConvertCommentDomainToResponse(commentDomain))
}
