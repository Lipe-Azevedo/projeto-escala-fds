package repository

import (
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	// MONGODB_SWAP_COLLECTION_ENV_KEY armazena o nome da variável de ambiente que contém o nome da coleção shift_swap.
	MONGODB_SWAP_COLLECTION_ENV_KEY = "MONGODB_SWAP_COLLECTION"
)

func NewSwapRepository(
	database *mongo.Database,
) SwapRepository {
	return &swapRepository{
		databaseConnection: database, // Corrigido: era 'database' diretamente
	}
}

type swapRepository struct {
	databaseConnection *mongo.Database
}

// FindSwapsByUserID implements SwapRepository.
func (sr *swapRepository) FindSwapsByUserID(userID string) ([]model.SwapDomainInterface, *rest_err.RestErr) {
	panic("unimplemented")
}

// FindSwapsByStatus implements SwapRepository.
// Implementar depois
func (sr *swapRepository) FindSwapsByStatus(status model.SwapStatus) ([]model.SwapDomainInterface, *rest_err.RestErr) {
	panic("unimplemented")
}

type SwapRepository interface {
	CreateSwap(
		swapDomain model.SwapDomainInterface,
	) (model.SwapDomainInterface, *rest_err.RestErr)

	FindSwapByID(
		id string,
	) (model.SwapDomainInterface, *rest_err.RestErr)

	// Métodos adicionados na Fase 3 para consultar trocas de turno
	FindSwapsByUserID(
		userID string, // Pode ser solicitante ou solicitado
	) ([]model.SwapDomainInterface, *rest_err.RestErr)

	FindSwapsByStatus(
		status model.SwapStatus,
	) ([]model.SwapDomainInterface, *rest_err.RestErr)
	// Fim dos novos métodos

	UpdateSwap(
		id string,
		swapDomain model.SwapDomainInterface,
	) *rest_err.RestErr
}
