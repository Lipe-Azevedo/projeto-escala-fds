package converter

import (
	"github.com/Lipe-Azevedo/escala-fds/src/model"
	"github.com/Lipe-Azevedo/escala-fds/src/model/repository/entity"
)

func ConvertWorkInfoEntityToDomain(
	entity entity.WorkInfoEntity,
) model.WorkInfoDomainInterface {
	return model.NewWorkInfoDomain(
		entity.UserID,
		model.Team(entity.Team),
		entity.Position,
		model.Shift(entity.DefaultShift),
		model.Weekday(entity.WeekdayOff),
		model.WeekendDayOff(entity.WeekendDayOff),
		entity.SuperiorID,
	)
}
