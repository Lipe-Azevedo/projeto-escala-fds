package request

import "time"

// CommentRequest é usado para criar um novo comentário.
type CommentRequest struct {
	CollaboratorID string    `json:"collaborator_id" binding:"required"`
	Date           time.Time `json:"date" binding:"required" time_format:"2006-01-02"` // Data no formato YYYY-MM-DD
	Text           string    `json:"text" binding:"required,min=1,max=1000"`
}

// CommentUpdateRequest é usado para atualizar um comentário existente.
type CommentUpdateRequest struct {
	Text string `json:"text" binding:"required,min=1,max=1000"`
}
