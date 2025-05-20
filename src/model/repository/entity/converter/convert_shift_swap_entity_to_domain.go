package converter

import (
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model/repository/entity"
)

func ConvertShiftSwapEntityToDomain(
	entity entity.ShiftSwapEntity,
) model.ShiftSwapDomainInterface {
	return model.NewShiftSwapDomain(
		entity.RequesterID,
		entity.RequestedID,
		model.Shift(entity.CurrentShift),
		model.Shift(entity.NewShift),
		model.Weekday(entity.CurrentDayOff),
		model.Weekday(entity.NewDayOff),
		entity.Reason,
	).(model.ShiftSwapDomainInterface)
}
