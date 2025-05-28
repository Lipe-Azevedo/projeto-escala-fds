package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CommentEntity representa a estrutura de um comentário no MongoDB.
type CommentEntity struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	CollaboratorID string             `bson:"collaborator_id"` // ID do usuário colaborador
	AuthorID       string             `bson:"author_id"`       // ID do usuário que criou o comentário (master)
	Date           time.Time          `bson:"date"`            // Data específica a qual o comentário se refere (armazenar como UTC meia-noite)
	Text           string             `bson:"text"`
	CreatedAt      time.Time          `bson:"created_at"`
	UpdatedAt      *time.Time         `bson:"updated_at,omitempty"` // Opcional, apenas se atualizado
}
