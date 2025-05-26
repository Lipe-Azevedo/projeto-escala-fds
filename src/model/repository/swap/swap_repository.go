package swap

import (
	"github.com/Lipe-Azevedo/escala-fds/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/escala-fds/src/model/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	// MONGODB_SWAPS_COLLECTION_ENV_KEY armazena o nome da variável de ambiente que contém o nome da coleção de trocas.
	MONGODB_SWAPS_COLLECTION_ENV_KEY = "MONGODB_SWAPS_COLLECTION" // No seu .env está MONGODB_SWAP_COLLECTION, ajustei para plural aqui para consistência com Users/WorkInfos. Confirme qual você prefere. Vou usar SWAPS por enquanto.
)

// SwapRepository define a interface para o repositório de trocas.
type SwapRepository interface {
	CreateSwap(
		swapDomain domain.SwapDomainInterface,
	) (domain.SwapDomainInterface, *rest_err.RestErr)

	FindSwapByID(
		id string,
	) (domain.SwapDomainInterface, *rest_err.RestErr)

	FindSwapsByUserID(
		userID string,
	) ([]domain.SwapDomainInterface, *rest_err.RestErr)

	FindSwapsByStatus(
		status domain.SwapStatus,
	) ([]domain.SwapDomainInterface, *rest_err.RestErr)

	UpdateSwap(
		id string,
		swapDomain domain.SwapDomainInterface,
	) *rest_err.RestErr
}

// swapRepository é a implementação da interface SwapRepository.
type swapRepository struct {
	databaseConnection *mongo.Database
}

// NewSwapRepository cria uma nova instância de SwapRepository.
func NewSwapRepository(
	database *mongo.Database,
) SwapRepository {
	return &swapRepository{
		databaseConnection: database,
	}
}
