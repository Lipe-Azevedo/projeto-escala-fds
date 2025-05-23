package repository

import (
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	// MONGODB_SWAPS_COLLECTION_ENV_KEY armazena o nome da variável de ambiente que contém o nome da coleção de trocas.
	MONGODB_SWAPS_COLLECTION_ENV_KEY = "MONGODB_SWAPS_COLLECTION"
)

// NewSwapRepository cria uma nova instância de SwapRepository.
func NewSwapRepository(
	database *mongo.Database,
) SwapRepository {
	return &swapRepository{
		databaseConnection: database,
	}
}

// swapRepository é a implementação da interface SwapRepository.
// Seus métodos são definidos em arquivos separados (ex: create_swap_repository.go, find_swap_repository.go, etc.).
type swapRepository struct {
	databaseConnection *mongo.Database
}

// SwapRepository define a interface para o repositório de trocas.
type SwapRepository interface {
	CreateSwap( // Implementação, por exemplo, em create_swap_repository.go
		swapDomain model.SwapDomainInterface,
	) (model.SwapDomainInterface, *rest_err.RestErr)

	FindSwapByID( // Implementação, por exemplo, em find_swap_repository.go
		id string,
	) (model.SwapDomainInterface, *rest_err.RestErr)

	FindSwapsByUserID( // Implementação, por exemplo, em find_swap_repository.go
		userID string,
	) ([]model.SwapDomainInterface, *rest_err.RestErr)

	FindSwapsByStatus( // Implementação, por exemplo, em find_swap_repository.go
		status model.SwapStatus,
	) ([]model.SwapDomainInterface, *rest_err.RestErr)

	UpdateSwap( // Implementação, por exemplo, em update_swap_repository.go
		id string,
		swapDomain model.SwapDomainInterface,
	) *rest_err.RestErr
}
