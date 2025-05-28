package view

import (
	"time"

	comment_response_dto "github.com/Lipe-Azevedo/escala-fds/src/controller/comment/response" // NOVO IMPORT
	swap_response_dto "github.com/Lipe-Azevedo/escala-fds/src/controller/swap/response"
	user_response_dto "github.com/Lipe-Azevedo/escala-fds/src/controller/user/response"
	workinfo_response_dto "github.com/Lipe-Azevedo/escala-fds/src/controller/workinfo/response"

	"github.com/Lipe-Azevedo/escala-fds/src/model/domain"
)

// ConvertUserDomainToResponse agora aceita um workInfoDomain opcional.
func ConvertUserDomainToResponse(
	userDomain domain.UserDomainInterface,
	workInfoDomain domain.WorkInfoDomainInterface, // Pode ser nil
) user_response_dto.UserResponse {
	userResp := user_response_dto.UserResponse{
		ID:       userDomain.GetID(),
		Email:    userDomain.GetEmail(),
		Name:     userDomain.GetName(),
		UserType: string(userDomain.GetUserType()),
	}

	if workInfoDomain != nil && userDomain.GetUserType() == domain.UserTypeCollaborator {
		wiResp := ConvertWorkInfoDomainToResponse(workInfoDomain)
		userResp.WorkInfo = &wiResp
	}

	return userResp
}

func ConvertWorkInfoDomainToResponse(
	workInfoDomain domain.WorkInfoDomainInterface,
) workinfo_response_dto.WorkInfoResponse {
	return workinfo_response_dto.WorkInfoResponse{
		UserID:        workInfoDomain.GetUserId(),
		Team:          string(workInfoDomain.GetTeam()),
		Position:      workInfoDomain.GetPosition(),
		DefaultShift:  string(workInfoDomain.GetDefaultShift()),
		WeekdayOff:    string(workInfoDomain.GetWeekdayOff()),
		WeekendDayOff: string(workInfoDomain.GetWeekendDayOff()),
		SuperiorID:    workInfoDomain.GetSuperiorID(),
	}
}

func ConvertSwapDomainToResponse(
	swapDomainVal domain.SwapDomainInterface,
) swap_response_dto.SwapResponse {
	approvedAt := formatTimePointer(swapDomainVal.GetApprovedAt())
	approvedBy := swapDomainVal.GetApprovedBy()

	return swap_response_dto.SwapResponse{
		ID:            swapDomainVal.GetID(),
		RequesterID:   swapDomainVal.GetRequesterID(),
		RequestedID:   swapDomainVal.GetRequestedID(),
		CurrentShift:  string(swapDomainVal.GetCurrentShift()),
		NewShift:      string(swapDomainVal.GetNewShift()),
		CurrentDayOff: string(swapDomainVal.GetCurrentDayOff()),
		NewDayOff:     string(swapDomainVal.GetNewDayOff()),
		Status:        string(swapDomainVal.GetStatus()),
		Reason:        swapDomainVal.GetReason(),
		CreatedAt:     swapDomainVal.GetCreatedAt().Format(time.RFC3339),
		ApprovedAt:    approvedAt,
		ApprovedBy:    approvedBy,
	}
}

// ConvertCommentDomainToResponse converte um CommentDomain para CommentResponse DTO.
func ConvertCommentDomainToResponse( // <-- NOVA FUNÇÃO
	commentDomain domain.CommentDomainInterface,
) comment_response_dto.CommentResponse {
	return comment_response_dto.CommentResponse{
		ID:             commentDomain.GetID(),
		CollaboratorID: commentDomain.GetCollaboratorID(),
		AuthorID:       commentDomain.GetAuthorID(),
		Date:           commentDomain.GetDate().Format("2006-01-02"), // Formata a data
		Text:           commentDomain.GetText(),
		CreatedAt:      commentDomain.GetCreatedAt(),
		UpdatedAt:      commentDomain.GetUpdatedAt(),
	}
}

// formatTimePointer é uma função auxiliar para formatar *time.Time para *string.
func formatTimePointer(t *time.Time) *string {
	if t == nil {
		return nil
	}
	formatted := t.Format(time.RFC3339)
	return &formatted
}
