package comment

import (
	"time"

	"github.com/Lipe-Azevedo/escala-fds/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/escala-fds/src/controller/comment/request" // Para DTOs de request
	"github.com/Lipe-Azevedo/escala-fds/src/model/domain"
	repo_comment "github.com/Lipe-Azevedo/escala-fds/src/model/repository/comment"
	service_user "github.com/Lipe-Azevedo/escala-fds/src/model/service/user" // Para validar existência de colaborador e autor
)

// CommentDomainService define a interface para os serviços de domínio de Comentário.
type CommentDomainService interface {
	CreateCommentService(
		commentReq request.CommentRequest, // Usando DTO de request
		authorID string, // Obtido do JWT
	) (domain.CommentDomainInterface, *rest_err.RestErr)

	FindCommentByIDService(commentID string) (domain.CommentDomainInterface, *rest_err.RestErr)

	FindCommentsByCollaboratorAndDateService(
		collaboratorID string,
		date time.Time,
		requestingUserID string, // Para verificações de permissão (quem está pedindo)
		requestingUserType domain.UserType,
	) ([]domain.CommentDomainInterface, *rest_err.RestErr)

	FindCommentsByCollaboratorForDateRangeService( // Nome ligeiramente alterado para clareza
		collaboratorID string,
		startDate time.Time,
		endDate time.Time,
		requestingUserID string,
		requestingUserType domain.UserType,
	) ([]domain.CommentDomainInterface, *rest_err.RestErr)

	UpdateCommentService(
		commentID string,
		commentUpdateReq request.CommentUpdateRequest, // Usando DTO de update
		requestingUserID string, // Quem está tentando atualizar
		requestingUserType domain.UserType,
	) (domain.CommentDomainInterface, *rest_err.RestErr)

	DeleteCommentService(
		commentID string,
		requestingUserID string,
		requestingUserType domain.UserType,
	) *rest_err.RestErr
}

// commentDomainService é a implementação da CommentDomainService.
type commentDomainService struct {
	commentRepository repo_comment.CommentRepository
	userService       service_user.UserDomainService // Para validar se colaborador e autor existem
}

// NewCommentDomainService cria uma nova instância de CommentDomainService.
func NewCommentDomainService(
	commentRepo repo_comment.CommentRepository,
	userService service_user.UserDomainService,
) CommentDomainService {
	return &commentDomainService{
		commentRepository: commentRepo,
		userService:       userService,
	}
}
