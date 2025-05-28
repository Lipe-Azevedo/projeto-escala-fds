package workinfo

import (
	"github.com/Lipe-Azevedo/escala-fds/src/model/domain"
	"github.com/Lipe-Azevedo/escala-fds/src/model/repository/entity"
)

func ConvertWorkInfoEntityToDomain(
	entityData entity.WorkInfoEntity,
) domain.WorkInfoDomainInterface {
	return domain.NewWorkInfoDomain(
		entityData.UserID,
		domain.Team(entityData.Team),
		entityData.Position,
		domain.Shift(entityData.DefaultShift),
		domain.Weekday(entityData.WeekdayOff),
		domain.WeekendDayOff(entityData.WeekendDayOff),
		entityData.SuperiorID,
	)
}
