package converter

import (
	"github.com/Lipe-Azevedo/escala-fds/src/model"
	"github.com/Lipe-Azevedo/escala-fds/src/model/repository/entity"
)

func ConvertDomainToEntity(
	domain model.UserDomainInterface,
) *entity.UserEntity {
	return &entity.UserEntity{
		Email:    domain.GetEmail(),
		Password: domain.GetPassword(),
		Name:     domain.GetName(),
		UserType: string(domain.GetUserType()),
	}
}
