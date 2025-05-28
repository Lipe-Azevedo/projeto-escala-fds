package domain

import (
	"time"
)

// CommentDomainInterface define a interface para o domínio de Comentário.
type CommentDomainInterface interface {
	GetID() string
	GetCollaboratorID() string
	GetAuthorID() string
	GetDate() time.Time // Data específica a qual o comentário se refere
	GetText() string
	GetCreatedAt() time.Time
	GetUpdatedAt() *time.Time // Pode ser nil se nunca atualizado

	SetID(string)
	SetText(string)         // Permitir atualização do texto
	SetUpdatedAt(time.Time) // Para registrar quando foi atualizado
}

// commentDomain é a struct que representa o domínio de Comentário.
type commentDomain struct {
	id             string
	collaboratorID string
	authorID       string
	date           time.Time // Data do evento/dia comentado
	text           string
	createdAt      time.Time
	updatedAt      *time.Time
}

// NewCommentDomain construtor para CommentDomainInterface.
// A data (date) é o dia específico para o qual o comentário é feito.
func NewCommentDomain(
	collaboratorID string,
	authorID string,
	date time.Time, // Data do comentário (referente ao dia)
	text string,
) CommentDomainInterface {
	// Normaliza a data para garantir que apenas ano, mês e dia sejam relevantes, zerando horas/min/etc.
	normalizedDate := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
	return &commentDomain{
		collaboratorID: collaboratorID,
		authorID:       authorID,
		date:           normalizedDate,
		text:           text,
		createdAt:      time.Now(),
	}
}

// Getters
func (cd *commentDomain) GetID() string             { return cd.id }
func (cd *commentDomain) GetCollaboratorID() string { return cd.collaboratorID }
func (cd *commentDomain) GetAuthorID() string       { return cd.authorID }
func (cd *commentDomain) GetDate() time.Time        { return cd.date }
func (cd *commentDomain) GetText() string           { return cd.text }
func (cd *commentDomain) GetCreatedAt() time.Time   { return cd.createdAt }
func (cd *commentDomain) GetUpdatedAt() *time.Time  { return cd.updatedAt }

// Setters
func (cd *commentDomain) SetID(id string)                  { cd.id = id }
func (cd *commentDomain) SetText(text string)              { cd.text = text }
func (cd *commentDomain) SetUpdatedAt(updatedAt time.Time) { cd.updatedAt = &updatedAt }
