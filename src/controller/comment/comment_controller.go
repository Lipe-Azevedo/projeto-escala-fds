package comment

import (
	service_comment "github.com/Lipe-Azevedo/escala-fds/src/model/service/comment"
	"github.com/gin-gonic/gin"
)

// CommentControllerInterface define a interface para os controllers de Comentário.
type CommentControllerInterface interface {
	CreateComment(c *gin.Context)
	FindCommentByID(c *gin.Context)
	FindCommentsByCollaboratorAndDate(c *gin.Context)      // GET /api/comments/collaborator/:collaboratorId/date/:dateString
	FindCommentsByCollaboratorForDateRange(c *gin.Context) // GET /api/comments/collaborator/:collaboratorId/range?startDate=...&endDate=...
	UpdateComment(c *gin.Context)
	DeleteComment(c *gin.Context)
}

// commentControllerInterface é a implementação da CommentControllerInterface.
type commentControllerInterface struct {
	service service_comment.CommentDomainService
}

// NewCommentControllerInterface cria uma nova instância de CommentControllerInterface.
func NewCommentControllerInterface(
	serviceInterface service_comment.CommentDomainService,
) CommentControllerInterface {
	return &commentControllerInterface{
		service: serviceInterface,
	}
}
