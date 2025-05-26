package workinfo

import (
	"github.com/Lipe-Azevedo/escala-fds/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/escala-fds/src/model/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	// MONGODB_WORKINFO_COLLECTION_ENV_KEY armazena o nome da variável de ambiente que contém o nome da coleção work_info.
	MONGODB_WORKINFO_COLLECTION_ENV_KEY = "MONGODB_WORKINFO_COLLECTION"
)

// WorkInfoRepository define a interface para o repositório de WorkInfo.
type WorkInfoRepository interface {
	CreateWorkInfo(
		workInfoDomain domain.WorkInfoDomainInterface,
	) (domain.WorkInfoDomainInterface, *rest_err.RestErr)

	FindWorkInfoByUserId(
		userId string,
	) (domain.WorkInfoDomainInterface, *rest_err.RestErr)

	UpdateWorkInfo(
		userId string,
		workInfoDomain domain.WorkInfoDomainInterface,
	) *rest_err.RestErr // Modificado para refletir a lógica de atualização que pode não retornar o domínio diretamente do repo
}

// workInfoRepository é a implementação da interface WorkInfoRepository.
type workInfoRepository struct {
	databaseConnection *mongo.Database // Nome do campo padronizado
}

// NewWorkInfoRepository cria uma nova instância de WorkInfoRepository.
func NewWorkInfoRepository(
	database *mongo.Database,
) WorkInfoRepository {
	return &workInfoRepository{
		databaseConnection: database, // Nome do campo padronizado
	}
}
