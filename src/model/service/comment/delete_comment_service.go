package comment

import (
	// "fmt"
	// "net/http"

	"github.com/Lipe-Azevedo/escala-fds/src/configuration/logger"
	"github.com/Lipe-Azevedo/escala-fds/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/escala-fds/src/model/domain"
	"go.uber.org/zap"
)

func (cs *commentDomainService) DeleteCommentService(
	commentID string,
	requestingUserID string,
	requestingUserType domain.UserType,
) *rest_err.RestErr {
	logger.Info("Init DeleteCommentService",
		zap.String("commentID", commentID),
		zap.String("requestingUserID", requestingUserID))

	if requestingUserType != domain.UserTypeMaster {
		return rest_err.NewForbiddenError("Permission denied to delete comment")
	}

	// Verificar se o comentário existe antes de tentar deletar
	_, err := cs.commentRepository.FindCommentByID(commentID)
	if err != nil {
		return err // Retorna erro se não encontrado ou outro erro do repositório
	}

	// Lógica de permissão adicional (ex: só autor ou master) pode ser adicionada aqui
	// if existingComment.GetAuthorID() != requestingUserID && requestingUserType != domain.UserTypeMaster {
	//  return rest_err.NewForbiddenError("Only the author or a master can delete this comment")
	// }

	return cs.commentRepository.DeleteComment(commentID)
}
