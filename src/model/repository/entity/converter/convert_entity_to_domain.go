package converter

import (
	"github.com/Lipe-Azevedo/escala-fds/src/model/domain" // CORRETO: usa novo módulo e pacote domain
	"github.com/Lipe-Azevedo/escala-fds/src/model/repository/entity"
)

func ConvertEntityToDomain(
	userEntity entity.UserEntity,
) domain.UserDomainInterface { // CORRETO: retorna domain.UserDomainInterface
	domainObj := domain.NewUserDomain( // CORRETO: usa construtor do pacote domain
		userEntity.Email,
		userEntity.Password,
		userEntity.Name,
		domain.UserType(userEntity.UserType), // CORRETO: usa domain.UserType
	)

	domainObj.SetID(userEntity.ID.Hex())
	return domainObj
}
