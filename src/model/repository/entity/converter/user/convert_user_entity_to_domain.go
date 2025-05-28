package user

import (
	"github.com/Lipe-Azevedo/escala-fds/src/model/domain"
	"github.com/Lipe-Azevedo/escala-fds/src/model/repository/entity"
)

func ConvertEntityToDomain(
	userEntity entity.UserEntity,
) domain.UserDomainInterface {
	domainObj := domain.NewUserDomain(
		userEntity.Email,
		userEntity.Password, // Senha já hasheada do BD
		userEntity.Name,
		domain.UserType(userEntity.UserType),
	)
	domainObj.SetID(userEntity.ID.Hex())
	return domainObj
}
