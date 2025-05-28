package comment

import (
	"github.com/Lipe-Azevedo/escala-fds/src/model/domain"
	"github.com/Lipe-Azevedo/escala-fds/src/model/repository/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ConvertCommentDomainToEntity converte um CommentDomain para CommentEntity.
func ConvertCommentDomainToEntity(
	commentDomain domain.CommentDomainInterface,
) *entity.CommentEntity {
	entity := &entity.CommentEntity{
		CollaboratorID: commentDomain.GetCollaboratorID(),
		AuthorID:       commentDomain.GetAuthorID(),
		Date:           commentDomain.GetDate(), // A data já deve estar normalizada no domínio
		Text:           commentDomain.GetText(),
		CreatedAt:      commentDomain.GetCreatedAt(),
		UpdatedAt:      commentDomain.GetUpdatedAt(),
	}

	// Se o ID do domínio já existe (ex: para atualização), converte para ObjectID
	if commentDomain.GetID() != "" {
		objID, _ := primitive.ObjectIDFromHex(commentDomain.GetID())
		entity.ID = objID
	}

	return entity
}