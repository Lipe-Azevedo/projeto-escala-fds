package swap

import (
	// Import para o novo pacote do serviço swap
	service_swap "github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model/service/swap"
	"github.com/gin-gonic/gin"
)

// SwapControllerInterface define a interface para os controllers de Swap.
type SwapControllerInterface interface {
	CreateSwap(c *gin.Context)
	FindSwapByID(c *gin.Context)
	UpdateSwapStatus(c *gin.Context)
	// Adicionar outros métodos se necessário, ex: FindSwapsByUser, FindAllPendingSwaps etc.
}

// swapControllerInterface é a implementação da interface SwapControllerInterface.
type swapControllerInterface struct {
	service service_swap.SwapDomainService // Tipo ajustado
}

// NewSwapControllerInterface cria uma nova instância de SwapControllerInterface.
func NewSwapControllerInterface(
	serviceInterface service_swap.SwapDomainService, // Tipo ajustado
) SwapControllerInterface {
	return &swapControllerInterface{
		service: serviceInterface,
	}
}
