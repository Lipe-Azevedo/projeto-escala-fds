package view

import (
	"time"

	// DTOs de User (já reorganizado)
	user_response_dto "github.com/Lipe-Azevedo/meu-primeio-crud-go/src/controller/user/response"
	// DTOs de WorkInfo (já reorganizado)
	workinfo_response_dto "github.com/Lipe-Azevedo/meu-primeio-crud-go/src/controller/workinfo/response"
	// DTOs de Swap (já reorganizado)
	swap_response_dto "github.com/Lipe-Azevedo/meu-primeio-crud-go/src/controller/swap/response"

	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model"
)

// ConvertUserDomainToResponse converte um UserDomainInterface para user_response_dto.UserResponse.
func ConvertUserDomainToResponse(
	userDomain model.UserDomainInterface,
) user_response_dto.UserResponse {
	return user_response_dto.UserResponse{
		ID:       userDomain.GetID(),
		Email:    userDomain.GetEmail(),
		Name:     userDomain.GetName(),
		UserType: string(userDomain.GetUserType()),
		// WorkInfo:
	}
}

// ConvertWorkInfoDomainToResponse agora usa workinfo_response_dto.WorkInfoResponse.
func ConvertWorkInfoDomainToResponse(
	workInfoDomain model.WorkInfoDomainInterface,
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

// ConvertSwapDomainToResponse agora usa swap_response_dto.SwapResponse.
func ConvertSwapDomainToResponse(
	domain model.SwapDomainInterface,
) swap_response_dto.SwapResponse {
	approvedAt := formatTimePointer(domain.GetApprovedAt())
	approvedBy := domain.GetApprovedBy()

	return swap_response_dto.SwapResponse{
		ID:            domain.GetID(),
		RequesterID:   domain.GetRequesterID(),
		RequestedID:   domain.GetRequestedID(),
		CurrentShift:  string(domain.GetCurrentShift()),
		NewShift:      string(domain.GetNewShift()),
		CurrentDayOff: string(domain.GetCurrentDayOff()),
		NewDayOff:     string(domain.GetNewDayOff()),
		Status:        string(domain.GetStatus()),
		Reason:        domain.GetReason(),
		CreatedAt:     domain.GetCreatedAt().Format(time.RFC3339),
		ApprovedAt:    approvedAt,
		ApprovedBy:    approvedBy,
	}
}

func formatTimePointer(t *time.Time) *string {
	if t == nil {
		return nil
	}
	formatted := t.Format(time.RFC3339)
	return &formatted
}
