package view

import (
	"time"

	// Import para o novo local do UserResponse
	user_response_dto "github.com/Lipe-Azevedo/meu-primeio-crud-go/src/controller/user/response"
	// Import para os responses de WorkInfo e Swap (ainda dos locais antigos, serão ajustados)
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/controller/model/response" // Mantido para WorkInfoResponse e SwapResponse por enquanto
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model"
)

// ConvertUserDomainToResponse converte um UserDomainInterface para UserResponse.
// Nome da função alterado para maior clareza.
func ConvertUserDomainToResponse( // <<< RENOMEADO AQUI
	userDomain model.UserDomainInterface,
) user_response_dto.UserResponse {
	return user_response_dto.UserResponse{
		ID:       userDomain.GetID(),
		Email:    userDomain.GetEmail(),
		Name:     userDomain.GetName(),
		UserType: string(userDomain.GetUserType()),
		// WorkInfo: // Será descomentado e populado quando WorkInfoResponse for reorganizado
	}
}

// ConvertWorkInfoDomainToResponse permanece o mesmo por enquanto, usando o response global.
// Será ajustado quando WorkInfo for reorganizado.
func ConvertWorkInfoDomainToResponse(
	workInfoDomain model.WorkInfoDomainInterface,
) response.WorkInfoResponse {
	return response.WorkInfoResponse{
		UserID:        workInfoDomain.GetUserId(),
		Team:          string(workInfoDomain.GetTeam()),
		Position:      workInfoDomain.GetPosition(),
		DefaultShift:  string(workInfoDomain.GetDefaultShift()),
		WeekdayOff:    string(workInfoDomain.GetWeekdayOff()),
		WeekendDayOff: string(workInfoDomain.GetWeekendDayOff()),
		SuperiorID:    workInfoDomain.GetSuperiorID(),
	}
}

// ConvertSwapDomainToResponse permanece o mesmo por enquanto, usando o response global.
// Será ajustado quando Swap for reorganizado.
func ConvertSwapDomainToResponse(
	domain model.SwapDomainInterface,
) response.SwapResponse {
	approvedAt := formatTimePointer(domain.GetApprovedAt())
	approvedBy := domain.GetApprovedBy()

	return response.SwapResponse{
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
