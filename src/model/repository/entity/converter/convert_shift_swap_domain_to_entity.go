package converter

import (
	"github.com/Lipe-Azevedo/escala-fds/src/model"
	"github.com/Lipe-Azevedo/escala-fds/src/model/repository/entity"
)

func ConvertSwapDomainToEntity(
	domain model.SwapDomainInterface,
) *entity.SwapEntity {
	return &entity.SwapEntity{
		ID:            domain.GetID(),
		RequesterID:   domain.GetRequesterID(),
		RequestedID:   domain.GetRequestedID(),
		CurrentShift:  string(domain.GetCurrentShift()),
		NewShift:      string(domain.GetNewShift()),
		CurrentDayOff: string(domain.GetCurrentDayOff()),
		NewDayOff:     string(domain.GetNewDayOff()),
		Status:        string(domain.GetStatus()),
		Reason:        domain.GetReason(),
		CreatedAt:     domain.GetCreatedAt(),
		ApprovedAt:    domain.GetApprovedAt(),
		ApprovedBy:    domain.GetApprovedBy(),
	}
}
