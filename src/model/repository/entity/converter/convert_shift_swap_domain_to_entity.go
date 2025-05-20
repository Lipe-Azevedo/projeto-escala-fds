package converter

import (
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model/repository/entity"
)

func ConvertShiftSwapDomainToEntity(
	domain model.ShiftSwapDomainInterface,
) *entity.ShiftSwapEntity {
	return &entity.ShiftSwapEntity{
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
