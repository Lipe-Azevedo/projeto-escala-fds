package comment

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Lipe-Azevedo/escala-fds/src/configuration/logger"
	"github.com/Lipe-Azevedo/escala-fds/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/escala-fds/src/model/domain"
	"go.uber.org/zap"
)

func (cs *commentDomainService) FindCommentByIDService(commentID string) (domain.CommentDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindCommentByIDService", zap.String("commentID", commentID))

	// Adicionar validação do commentID se necessário (formato, etc.)
	return cs.commentRepository.FindCommentByID(commentID)
}

func (cs *commentDomainService) FindCommentsByCollaboratorAndDateService(
	collaboratorID string,
	date time.Time,
	requestingUserID string,
	requestingUserType domain.UserType,
) ([]domain.CommentDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindCommentsByCollaboratorAndDateService",
		zap.String("collaboratorID", collaboratorID),
		zap.Time("date", date),
		zap.String("requestingUserID", requestingUserID))

	// Validação de permissão (ex: só master pode ver)
	if requestingUserType != domain.UserTypeMaster {
		return nil, rest_err.NewForbiddenError("Permission denied to view comments")
	}

	// Validar se o colaborador existe
	_, err := cs.userService.FindUserByIDServices(collaboratorID)
	if err != nil {
		if err.Code == http.StatusNotFound {
			return nil, rest_err.NewNotFoundError(fmt.Sprintf("Collaborator with ID %s not found", collaboratorID))
		}
		return nil, err
	}

	return cs.commentRepository.FindCommentsByCollaboratorAndDate(collaboratorID, date)
}

func (cs *commentDomainService) FindCommentsByCollaboratorForDateRangeService(
	collaboratorID string,
	startDate time.Time,
	endDate time.Time,
	requestingUserID string,
	requestingUserType domain.UserType,
) ([]domain.CommentDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindCommentsByCollaboratorForDateRangeService",
		zap.String("collaboratorID", collaboratorID),
		zap.Time("startDate", startDate),
		zap.Time("endDate", endDate))

	if requestingUserType != domain.UserTypeMaster {
		return nil, rest_err.NewForbiddenError("Permission denied to view comments")
	}
	_, err := cs.userService.FindUserByIDServices(collaboratorID)
	if err != nil {
		if err.Code == http.StatusNotFound {
			return nil, rest_err.NewNotFoundError(fmt.Sprintf("Collaborator with ID %s not found", collaboratorID))
		}
		return nil, err
	}

	if startDate.After(endDate) {
		return nil, rest_err.NewBadRequestError("Start date cannot be after end date")
	}

	return cs.commentRepository.FindCommentsByCollaboratorAndDateRange(collaboratorID, startDate, endDate)
}
