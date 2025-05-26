package view

import (
	"time"

	// DTOs de User (já reorganizado)
	user_response_dto "github.com/Lipe-Azevedo/escala-fds/src/controller/user/response"
	// DTOs de WorkInfo (já reorganizado)
	workinfo_response_dto "github.com/Lipe-Azevedo/escala-fds/src/controller/workinfo/response"
	// DTOs de Swap (já reorganizado)
	swap_response_dto "github.com/Lipe-Azevedo/escala-fds/src/controller/swap/response"

	// IMPORT ATUALIZADO: Agora importa do subpacote 'domain'
	"github.com/Lipe-Azevedo/escala-fds/src/model/domain"
)

// ConvertUserDomainToResponse converte um domain.UserDomainInterface para user_response_dto.UserResponse.
func ConvertUserDomainToResponse(
	userDomain domain.UserDomainInterface, // <<< USA domain.UserDomainInterface
) user_response_dto.UserResponse {
	return user_response_dto.UserResponse{
		ID:       userDomain.GetID(),
		Email:    userDomain.GetEmail(),
		Name:     userDomain.GetName(),
		UserType: string(userDomain.GetUserType()), // GetUserType() retorna domain.UserType
		// WorkInfo:
	}
}

// ConvertWorkInfoDomainToResponse agora usa domain.WorkInfoDomainInterface e workinfo_response_dto.WorkInfoResponse.
func ConvertWorkInfoDomainToResponse(
	workInfoDomain domain.WorkInfoDomainInterface, // <<< USA domain.WorkInfoDomainInterface
) workinfo_response_dto.WorkInfoResponse {
	return workinfo_response_dto.WorkInfoResponse{
		UserID:        workInfoDomain.GetUserId(),
		Team:          string(workInfoDomain.GetTeam()), // GetTeam() retorna domain.Team
		Position:      workInfoDomain.GetPosition(),
		DefaultShift:  string(workInfoDomain.GetDefaultShift()),  // GetDefaultShift() retorna domain.Shift
		WeekdayOff:    string(workInfoDomain.GetWeekdayOff()),    // GetWeekdayOff() retorna domain.Weekday
		WeekendDayOff: string(workInfoDomain.GetWeekendDayOff()), // GetWeekendDayOff() retorna domain.WeekendDayOff
		SuperiorID:    workInfoDomain.GetSuperiorID(),
	}
}

// ConvertSwapDomainToResponse agora usa domain.SwapDomainInterface e swap_response_dto.SwapResponse.
func ConvertSwapDomainToResponse(
	swapDomainVal domain.SwapDomainInterface, // <<< USA domain.SwapDomainInterface (renomeei param para evitar conflito com package)
) swap_response_dto.SwapResponse {
	approvedAt := formatTimePointer(swapDomainVal.GetApprovedAt())
	approvedBy := formatTimePointerToString(swapDomainVal.GetApprovedBy()) // Ajustado para lidar com *string

	return swap_response_dto.SwapResponse{
		ID:            swapDomainVal.GetID(),
		RequesterID:   swapDomainVal.GetRequesterID(),
		RequestedID:   swapDomainVal.GetRequestedID(),
		CurrentShift:  string(swapDomainVal.GetCurrentShift()),  // GetCurrentShift() retorna domain.Shift
		NewShift:      string(swapDomainVal.GetNewShift()),      // GetNewShift() retorna domain.Shift
		CurrentDayOff: string(swapDomainVal.GetCurrentDayOff()), // GetCurrentDayOff() retorna domain.Weekday
		NewDayOff:     string(swapDomainVal.GetNewDayOff()),     // GetNewDayOff() retorna domain.Weekday
		Status:        string(swapDomainVal.GetStatus()),        // GetStatus() retorna domain.SwapStatus
		Reason:        swapDomainVal.GetReason(),
		CreatedAt:     swapDomainVal.GetCreatedAt().Format(time.RFC3339),
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

// Nova função auxiliar para converter *string para *string (tratando nil)
// Ou simplesmente usar o valor diretamente se o DTO já espera *string
func formatTimePointerToString(s *string) *string {
	if s == nil {
		return nil
	}
	val := *s
	return &val
}
