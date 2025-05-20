package repository

import (
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	MONGODB_SHIFT_SWAP_DB = "MONGODB_SHIFT_SWAP_DB"
)

func NewShiftSwapRepository(
	database *mongo.Database,
) ShiftSwapRepository {
	return &shiftSwapRepository{
		database,
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

	UpdateShiftSwap(
		id string,
		shiftSwapDomain model.ShiftSwapDomainInterface,
	) *rest_err.RestErr
}
