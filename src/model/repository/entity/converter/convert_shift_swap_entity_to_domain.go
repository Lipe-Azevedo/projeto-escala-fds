package converter

import (
	"github.com/Lipe-Azevedo/escala-fds/src/model"
	"github.com/Lipe-Azevedo/escala-fds/src/model/repository/entity"
)

func ConvertSwapEntityToDomain(
	entity entity.SwapEntity,
) model.SwapDomainInterface {
	return model.NewSwapDomain(
		entity.RequesterID,
		entity.RequestedID,
		model.Shift(entity.CurrentShift),
		model.Shift(entity.NewShift),
		model.Weekday(entity.CurrentDayOff),
		model.Weekday(entity.NewDayOff),
		entity.Reason,
	).(model.SwapDomainInterface)
}
