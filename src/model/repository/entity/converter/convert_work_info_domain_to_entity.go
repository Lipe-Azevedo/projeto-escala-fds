package converter

import (
	"github.com/Lipe-Azevedo/escala-fds/src/model/domain"
	"github.com/Lipe-Azevedo/escala-fds/src/model/repository/entity"
)

func ConvertWorkInfoDomainToEntity(
	domain domain.WorkInfoDomainInterface,
) *entity.WorkInfoEntity {
	return &entity.WorkInfoEntity{
		UserID:        domain.GetUserId(),
		Team:          string(domain.GetTeam()),
		Position:      domain.GetPosition(),
		DefaultShift:  string(domain.GetDefaultShift()),
		WeekdayOff:    string(domain.GetWeekdayOff()),
		WeekendDayOff: string(domain.GetWeekendDayOff()),
		SuperiorID:    domain.GetSuperiorID(),
	}
}
