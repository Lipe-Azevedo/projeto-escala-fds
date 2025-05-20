package view

import (
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/controller/model/response"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model"
)

func ConvertDomainToResponse(
	userDomain model.UserDomainInterface,
) response.UserResponse {
	return response.UserResponse{
		ID:    userDomain.GetID(),
		Email: userDomain.GetEmail(),
		Name:  userDomain.GetName(),
	}
}

func ConvertWorkInfoDomainToResponse(
	workInfoDomain model.WorkInfoDomainInterface,
) response.WorkInfoResponse {
	return response.WorkInfoResponse{
		Team:          string(workInfoDomain.GetTeam()),
		Position:      workInfoDomain.GetPosition(),
		DefaultShift:  string(workInfoDomain.GetDefaultShift()),
		WeekdayOff:    string(workInfoDomain.GetWeekdayOff()),
		WeekendDayOff: string(workInfoDomain.GetWeekendDayOff()),
		SuperiorID:    workInfoDomain.GetSuperiorID(),
	}
}
