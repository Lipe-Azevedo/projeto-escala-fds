package repository

import (
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	// MONGODB_SHIFTSWAP_COLLECTION_ENV_KEY armazena o nome da variável de ambiente que contém o nome da coleção shift_swap.
	MONGODB_SHIFTSWAP_COLLECTION_ENV_KEY = "MONGODB_SHIFTSWAP_COLLECTION"
)

func NewShiftSwapRepository(
	database *mongo.Database,
) ShiftSwapRepository {
	return &shiftSwapRepository{
		databaseConnection: database, // Corrigido: era 'database' diretamente
	}
}

type shiftSwapRepository struct {
	databaseConnection *mongo.Database
}

type ShiftSwapRepository interface {
	CreateShiftSwap(
		shiftSwapDomain model.ShiftSwapDomainInterface,
	) (model.ShiftSwapDomainInterface, *rest_err.RestErr)

	FindShiftSwapByID(
		id string,
	) (model.ShiftSwapDomainInterface, *rest_err.RestErr)

	// Métodos adicionados na Fase 3 para consultar trocas de turno
	FindShiftSwapsByUserID(
		userID string, // Pode ser solicitante ou solicitado
	) ([]model.ShiftSwapDomainInterface, *rest_err.RestErr)

	FindShiftSwapsByStatus(
		status model.ShiftSwapStatus,
	) ([]model.ShiftSwapDomainInterface, *rest_err.RestErr)
	// Fim dos novos métodos

	UpdateShiftSwap(
		id string,
		shiftSwapDomain model.ShiftSwapDomainInterface,
	) *rest_err.RestErr
}
