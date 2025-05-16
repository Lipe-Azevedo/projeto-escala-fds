package converter

import (
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model/repository/entity"
)

func ConvertEntityToDomain(entity entity.UserEntity) model.UserDomainInterface {
	domain := model.NewUserDomain(
		entity.Email,
		entity.Password,
		entity.Name,
		model.UserType(entity.UserType),
	)

	domain.SetID(entity.ID.Hex())

	if entity.WorkInfo != nil {
		domain.SetWorkInfo(&model.WorkInfo{
			Team:          entity.WorkInfo.Team,
			Position:      entity.WorkInfo.Position,
			DefaultShift:  model.Shift(entity.WorkInfo.DefaultShift),
			WeekdayOff:    model.Weekday(entity.WorkInfo.WeekdayOff),
			WeekendDayOff: model.WeekendDayOff(entity.WorkInfo.WeekendDayOff),
			SuperiorID:    entity.WorkInfo.SuperiorID,
		})
	}

	return domain
}
