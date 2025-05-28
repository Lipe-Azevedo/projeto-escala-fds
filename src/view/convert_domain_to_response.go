package view

import (
	"time"

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
		// Converte o workInfoDomain para workinfo_response_dto.WorkInfoResponse
		// e atribui ao campo WorkInfo de userResp.
		// A função ConvertWorkInfoDomainToResponse já existe e faz essa conversão.
		wiResp := ConvertWorkInfoDomainToResponse(workInfoDomain)
		userResp.WorkInfo = &wiResp // Atribui o endereço da WorkInfoResponse convertida
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
	// Corrigido: GetApprovedBy retorna *string, então a conversão direta é melhor
	// approvedBy := formatTimePointerToString(swapDomainVal.GetApprovedBy())
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

func formatTimePointer(t *time.Time) *string {
	if t == nil {
		return nil
	}
	formatted := t.Format(time.RFC3339)
	return &formatted
}

// Removido formatTimePointerToString pois approvedBy já é *string
// func formatTimePointerToString(s *string) *string {
//     if s == nil {
//         return nil
//     }
//     val := *s
//     return &val
// }
