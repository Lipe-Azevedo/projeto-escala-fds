package request

import "time"

type CommentRequest struct {
	CollaboratorID string    `json:"collaborator_id" binding:"required"`
	Date           time.Time `json:"date" binding:"required"`
	Text           string    `json:"text" binding:"required,min=1,max=1000"`
}

type CommentUpdateRequest struct {
	Text string `json:"text" binding:"required,min=1,max=1000"`
}
