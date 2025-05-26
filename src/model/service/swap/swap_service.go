package swap

import (
	"github.com/Lipe-Azevedo/escala-fds/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/escala-fds/src/model/domain"

	repository_swap "github.com/Lipe-Azevedo/escala-fds/src/model/repository/swap"
)

type SwapDomainService interface {
	CreateSwapServices(
		swapDomain domain.SwapDomainInterface,
	) (domain.SwapDomainInterface, *rest_err.RestErr)

	FindSwapByIDServices(
		id string,
	) (domain.SwapDomainInterface, *rest_err.RestErr)

	UpdateSwapServices(
		id string,
		swapDomain domain.SwapDomainInterface,
	) *rest_err.RestErr
}

// swapDomainService é a implementação da interface SwapDomainService.
type swapDomainService struct {
	repository repository_swap.SwapRepository // Tipo ajustado
}

// NewSwapDomainService cria uma nova instância de SwapDomainService.
func NewSwapDomainService(
	repository repository_swap.SwapRepository, // Tipo ajustado
) SwapDomainService {
	return &swapDomainService{
		repository: repository,
	}
}
