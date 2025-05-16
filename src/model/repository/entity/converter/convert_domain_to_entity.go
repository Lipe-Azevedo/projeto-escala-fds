package converter

import (
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model/repository/entity"
)

func ConvertDomainToEntity(domain model.UserDomainInterface) *entity.UserEntity {
	entity := &entity.UserEntity{
		Email:    domain.GetEmail(),
		Password: domain.GetPassword(),
		Name:     domain.GetName(),
		UserType: string(domain.GetUserType()),
	}

	if domain.GetWorkInfo() != nil {
		entity.WorkInfo = &entity.WorkInfoEntity{
			Team:          domain.GetWorkInfo().Team,
			Position:      domain.GetWorkInfo().Position,
			DefaultShift:  string(domain.GetWorkInfo().DefaultShift),
			WeekdayOff:    string(domain.GetWorkInfo().WeekdayOff),
			WeekendDayOff: string(domain.GetWorkInfo().WeekendDayOff),
			SuperiorID:    domain.GetWorkInfo().SuperiorID,
		}
	}

	return entity
}
