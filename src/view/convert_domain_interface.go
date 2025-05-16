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
		Age:   userDomain.GetAge(),
	}
}
