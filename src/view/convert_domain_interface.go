package view

import (
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/controller/model/response"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model"
)

func ConvertDomainToResponse(
	userDomain model.UserDomainInterface,
) response.UserResponse {
	resp := response.UserResponse{
		ID:       userDomain.GetID(),
		Email:    userDomain.GetEmail(),
		Name:     userDomain.GetName(),
		UserType: string(userDomain.GetUserType()),
	}

	if userDomain.GetWorkInfo() != nil {
		resp.WorkInfo = &response.WorkInfoResponse{
			Team:          userDomain.GetWorkInfo().Team,
			Position:      userDomain.GetWorkInfo().Position,
			DefaultShift:  string(userDomain.GetWorkInfo().DefaultShift),
			WeekdayOff:    string(userDomain.GetWorkInfo().WeekdayOff),
			WeekendDayOff: string(userDomain.GetWorkInfo().WeekendDayOff),
			SuperiorID:    userDomain.GetWorkInfo().SuperiorID,
		}
	}

	return resp
}
