package comment

import (
	"fmt"

	"github.com/Lipe-Azevedo/escala-fds/src/configuration/logger"
	"github.com/Lipe-Azevedo/escala-fds/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/escala-fds/src/controller/comment/request"
	"github.com/Lipe-Azevedo/escala-fds/src/model/domain"
	"go.uber.org/zap"
)

func (cs *commentDomainService) CreateCommentService(
	commentReq request.CommentRequest,
	authorID string,
) (domain.CommentDomainInterface, *rest_err.RestErr) {
	logger.Info("Init CreateCommentService",
		zap.String("journey", "createComment"),
		zap.String("collaboratorId", commentReq.CollaboratorID),
		zap.String("authorId", authorID))

	// Validar se o colaborador existe
	collaborator, err := cs.userService.FindUserByIDServices(commentReq.CollaboratorID)
	if err != nil {
		logger.Error("Collaborator not found for comment creation", err, zap.String("collaboratorId", commentReq.CollaboratorID))
		return nil, rest_err.NewBadRequestError(fmt.Sprintf("Collaborator with ID %s not found", commentReq.CollaboratorID))
	}
	if collaborator.GetUserType() != domain.UserTypeCollaborator {
		return nil, rest_err.NewBadRequestError(fmt.Sprintf("User with ID %s is not a collaborator", commentReq.CollaboratorID))
	}

	// Validar se o autor existe (embora ele venha do token, uma verificação extra não faz mal)
	_, err = cs.userService.FindUserByIDServices(authorID)
	if err != nil {
		logger.Error("Author (master) not found for comment creation", err, zap.String("authorId", authorID))
		// Este seria um erro estranho, pois o autor vem de um token válido
		return nil, rest_err.NewInternalServerError(fmt.Sprintf("Author with ID %s not found", authorID))
	}

	// Normalizar a data para garantir que apenas ano, mês e dia sejam relevantes, zerando horas/min/etc.
	// O construtor do domínio já faz isso.
	commentDomain := domain.NewCommentDomain(
		commentReq.CollaboratorID,
		authorID,
		commentReq.Date,
		commentReq.Text,
	)

	createdComment, repoErr := cs.commentRepository.CreateComment(commentDomain)
	if repoErr != nil {
		logger.Error("Error calling repository to create comment", repoErr, zap.String("journey", "createComment"))
		return nil, repoErr
	}

	logger.Info("CreateCommentService executed successfully",
		zap.String("commentId", createdComment.GetID()),
		zap.String("journey", "createComment"))

	return createdComment, nil
}
