package converter

import (
	"github.com/Lipe-Azevedo/escala-fds/src/model/domain"
	"github.com/Lipe-Azevedo/escala-fds/src/model/repository/entity"
)

func ConvertSwapEntityToDomain(
	entity entity.SwapEntity,
) domain.SwapDomainInterface {
	return domain.NewSwapDomain(
		entity.RequesterID,
		entity.RequestedID,
		domain.Shift(entity.CurrentShift),
		domain.Shift(entity.NewShift),
		domain.Weekday(entity.CurrentDayOff),
		domain.Weekday(entity.NewDayOff),
		entity.Reason,
	).(domain.SwapDomainInterface)
}
