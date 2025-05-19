package converter

import (
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model/repository/entity"
)

func ConvertWorkInfoDomainToEntity(
	domain model.WorkInfoDomainInterface,
) *entity.WorkInfoEntity {
	return &entity.WorkInfoEntity{
		Team:          string(domain.GetTeam()),
		Position:      domain.GetPosition(),
		DefaultShift:  string(domain.GetDefaultShift()),
		WeekdayOff:    string(domain.GetWeekdayOff()),
		WeekendDayOff: string(domain.GetWeekendDayOff()),
		SuperiorID:    domain.GetSuperiorID(),
	}
}
