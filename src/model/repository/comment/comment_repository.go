package comment

import (
	"time"

	"github.com/Lipe-Azevedo/escala-fds/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/escala-fds/src/model/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	// MONGODB_COMMENTS_COLLECTION_ENV_KEY armazena o nome da variável de ambiente para a coleção de comentários.
	MONGODB_COMMENTS_COLLECTION_ENV_KEY = "MONGODB_COMMENTS_COLLECTION"
)

// CommentRepository define a interface para o repositório de comentários.
type CommentRepository interface {
	CreateComment(commentDomain domain.CommentDomainInterface) (domain.CommentDomainInterface, *rest_err.RestErr)
	FindCommentByID(commentID string) (domain.CommentDomainInterface, *rest_err.RestErr)
	FindCommentsByCollaboratorAndDate(collaboratorID string, date time.Time) ([]domain.CommentDomainInterface, *rest_err.RestErr)
	FindCommentsByCollaboratorAndDateRange(collaboratorID string, startDate time.Time, endDate time.Time) ([]domain.CommentDomainInterface, *rest_err.RestErr)
	UpdateComment(commentID string, commentDomain domain.CommentDomainInterface) (domain.CommentDomainInterface, *rest_err.RestErr)
	DeleteComment(commentID string) *rest_err.RestErr
}

// commentRepository é a implementação da interface CommentRepository.
type commentRepository struct {
	databaseConnection *mongo.Database
}

// NewCommentRepository cria uma nova instância de CommentRepository.
func NewCommentRepository(database *mongo.Database) CommentRepository {
	return &commentRepository{
		databaseConnection: database,
	}
}
