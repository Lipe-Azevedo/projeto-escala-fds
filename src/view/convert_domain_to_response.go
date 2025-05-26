package view

import (
	"time"

	// DTOs de User (já reorganizado)
	user_response_dto "github.com/Lipe-Azevedo/meu-primeio-crud-go/src/controller/user/response"
	// DTOs de WorkInfo (NOVO LOCAL)
	workinfo_response_dto "github.com/Lipe-Azevedo/meu-primeio-crud-go/src/controller/workinfo/response"

	// Alias para o pacote de response global/antigo (APENAS PARA SWAP AGORA)
	global_response "github.com/Lipe-Azevedo/meu-primeio-crud-go/src/controller/model/response"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model"
)

// ConvertUserDomainToResponse converte um UserDomainInterface para user_response_dto.UserResponse.
func ConvertUserDomainToResponse(
	userDomain model.UserDomainInterface,
) user_response_dto.UserResponse {
	// NOTA: O campo WorkInfo em user_response_dto.UserResponse ainda está comentado.
	// Popula-lo aqui exigiria que userDomain contivesse WorkInfo, ou que este conversor
	// tivesse acesso ao WorkInfoService. Isso será abordado na fase de refatoração
	// da lógica de User e WorkInfo, se desejado que UserResponse inclua WorkInfo.
	return user_response_dto.UserResponse{
		ID:       userDomain.GetID(),
		Email:    userDomain.GetEmail(),
		Name:     userDomain.GetName(),
		UserType: string(userDomain.GetUserType()),
		// WorkInfo: // Descomentar e popular quando a lógica de buscar User com WorkInfo for implementada.
	}
}

// ConvertWorkInfoDomainToResponse agora usa workinfo_response_dto.WorkInfoResponse.
func ConvertWorkInfoDomainToResponse(
	workInfoDomain model.WorkInfoDomainInterface,
) workinfo_response_dto.WorkInfoResponse { // <<< Tipo de retorno ajustado
	return workinfo_response_dto.WorkInfoResponse{ // <<< Usando o DTO específico de workinfo
		UserID:        workInfoDomain.GetUserId(),
		Team:          string(workInfoDomain.GetTeam()),
		Position:      workInfoDomain.GetPosition(),
		DefaultShift:  string(workInfoDomain.GetDefaultShift()),
		WeekdayOff:    string(workInfoDomain.GetWeekdayOff()),
		WeekendDayOff: string(workInfoDomain.GetWeekendDayOff()),
		SuperiorID:    workInfoDomain.GetSuperiorID(),
	}
}

// ConvertSwapDomainToResponse usa global_response.SwapResponse (temporariamente).
func ConvertSwapDomainToResponse(
	domain model.SwapDomainInterface,
) global_response.SwapResponse { // <<< Ainda usa o DTO global para Swap
	approvedAt := formatTimePointer(domain.GetApprovedAt())
	approvedBy := domain.GetApprovedBy()

	return global_response.SwapResponse{
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
