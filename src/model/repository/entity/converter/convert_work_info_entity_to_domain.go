package converter

import (
	"github.com/Lipe-Azevedo/escala-fds/src/model/domain"
	"github.com/Lipe-Azevedo/escala-fds/src/model/repository/entity"
)

func ConvertWorkInfoEntityToDomain(
	entity entity.WorkInfoEntity,
) domain.WorkInfoDomainInterface {
	return domain.NewWorkInfoDomain(
		entity.UserID,
		domain.Team(entity.Team),
		entity.Position,
		domain.Shift(entity.DefaultShift),
		domain.Weekday(entity.WeekdayOff),
		domain.WeekendDayOff(entity.WeekendDayOff),
		entity.SuperiorID,
	)
}
