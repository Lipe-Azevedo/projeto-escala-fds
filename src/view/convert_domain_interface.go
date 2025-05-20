package view

import (
	"time"

	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/controller/model/response"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model"
)

func ConvertDomainToResponse(
	userDomain model.UserDomainInterface,
) response.UserResponse {
	return response.UserResponse{
		ID:    userDomain.GetID(),
		Email: userDomain.GetEmail(),
		Name:  userDomain.GetName(),
	}
}

func ConvertWorkInfoDomainToResponse(
	workInfoDomain model.WorkInfoDomainInterface,
) response.WorkInfoResponse {
	return response.WorkInfoResponse{
		Team:          string(workInfoDomain.GetTeam()),
		Position:      workInfoDomain.GetPosition(),
		DefaultShift:  string(workInfoDomain.GetDefaultShift()),
		WeekdayOff:    string(workInfoDomain.GetWeekdayOff()),
		WeekendDayOff: string(workInfoDomain.GetWeekendDayOff()),
		SuperiorID:    workInfoDomain.GetSuperiorID(),
	}
}

func ConvertShiftSwapDomainToResponse(
	domain model.ShiftSwapDomainInterface,
) response.ShiftSwapResponse {
	approvedAt := formatTimePointer(domain.GetApprovedAt())
	approvedBy := domain.GetApprovedBy()

	return response.ShiftSwapResponse{
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
