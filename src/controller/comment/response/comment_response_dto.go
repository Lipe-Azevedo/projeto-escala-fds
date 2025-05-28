package response

import "time"

// CommentResponse representa a estrutura de um comentário para respostas JSON.
type CommentResponse struct {
	ID             string     `json:"id"`
	CollaboratorID string     `json:"collaborator_id"`
	AuthorID       string     `json:"author_id"`
	Date           string     `json:"date"` // Data formatada como YYYY-MM-DD
	Text           string     `json:"text"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      *time.Time `json:"updated_at,omitempty"`
}
