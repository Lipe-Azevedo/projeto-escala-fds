package swap

import (
	"github.com/Lipe-Azevedo/escala-fds/src/model/domain"
	"github.com/Lipe-Azevedo/escala-fds/src/model/repository/entity"
)

func ConvertSwapEntityToDomain(
	entityData entity.SwapEntity,
) domain.SwapDomainInterface {
	swapDomainInstance := domain.NewSwapDomain(
		entityData.RequesterID,
		entityData.RequestedID,
		domain.Shift(entityData.CurrentShift),
		domain.Shift(entityData.NewShift),
		domain.Weekday(entityData.CurrentDayOff),
		domain.Weekday(entityData.NewDayOff),
		entityData.Reason,
	)

	swapDomainInstance.SetID(entityData.ID)
	swapDomainInstance.SetStatus(domain.SwapStatus(entityData.Status))

	if entityData.ApprovedAt != nil {
		swapDomainInstance.SetApprovedAt(*entityData.ApprovedAt)
	}
	if entityData.ApprovedBy != nil {
		swapDomainInstance.SetApprovedBy(*entityData.ApprovedBy)
	}

	// **NOTA CRÍTICA SOBRE CreatedAt:** (Mesma nota da resposta anterior)
	// O campo `entityData.CreatedAt` não é transferido corretamente. `GetCreatedAt()`
	// retornará o `time.Now()` de `NewSwapDomain()`. Isso requer uma alteração
	// em `src/model/domain/swap_domain.go` para ser corrigido (novo construtor ou setter).

	return swapDomainInstance
}
