package comment

import (
	"github.com/Lipe-Azevedo/escala-fds/src/configuration/logger"
	"github.com/Lipe-Azevedo/escala-fds/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/escala-fds/src/controller/comment/request"
	"github.com/Lipe-Azevedo/escala-fds/src/model/domain"
	"go.uber.org/zap"
)

func (cs *commentDomainService) UpdateCommentService(
	commentID string,
	commentUpdateReq request.CommentUpdateRequest,
	requestingUserID string,
	requestingUserType domain.UserType,
) (domain.CommentDomainInterface, *rest_err.RestErr) {
	logger.Info("Init UpdateCommentService",
		zap.String("commentID", commentID),
		zap.String("requestingUserID", requestingUserID))

	if requestingUserType != domain.UserTypeMaster {
		return nil, rest_err.NewForbiddenError("Permission denied to update comment")
	}

	existingComment, err := cs.commentRepository.FindCommentByID(commentID)
	if err != nil {
		return nil, err // Retorna o erro do repositório (ex: NotFound)
	}

	// Lógica de permissão adicional: o autor original pode editar? Ou qualquer master?
	// Por enquanto, qualquer master pode (já verificado acima).
	// Se fosse apenas o autor:
	// if existingComment.GetAuthorID() != requestingUserID && requestingUserType != domain.UserTypeMaster {
	//  return nil, rest_err.NewForbiddenError("Only the author or a master can update this comment")
	// }

	existingComment.SetText(commentUpdateReq.Text)
	// UpdatedAt é setado no repositório

	return cs.commentRepository.UpdateComment(commentID, existingComment)
}
