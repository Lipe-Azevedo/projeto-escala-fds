package swap

import (
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model"

	// Import para o novo pacote do repositório swap
	repository_swap "github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model/repository/swap"
)

// SwapDomainService define a interface para os serviços de domínio de Swap.
type SwapDomainService interface {
	CreateSwapServices(
		swapDomain model.SwapDomainInterface,
	) (model.SwapDomainInterface, *rest_err.RestErr)

	FindSwapByIDServices(
		id string,
	) (model.SwapDomainInterface, *rest_err.RestErr)

	// Se precisar expor outros finders do repositório (ex: FindSwapsByUserID), adicione-os aqui.

	UpdateSwapServices(
		id string,
		swapDomain model.SwapDomainInterface, // O domain aqui pode ser um NewSwapUpdateDomain
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
