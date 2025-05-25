package workinfo

import (
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/controller/model/request" // Mantido para WorkInfoUpdateRequest na interface
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model"

	// Import para o novo pacote do repositório workinfo
	repository_workinfo "github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model/repository/workinfo"
	// Import para o novo pacote do serviço user (já reorganizado)
	service_user "github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model/service/user"
)

// WorkInfoDomainService define a interface para os serviços de domínio de WorkInfo.
type WorkInfoDomainService interface {
	CreateWorkInfoServices(
		workInfoDomain model.WorkInfoDomainInterface,
	) (model.WorkInfoDomainInterface, *rest_err.RestErr)

	FindWorkInfoByUserIdServices(
		userId string,
	) (model.WorkInfoDomainInterface, *rest_err.RestErr)

	// UpdateWorkInfoServices agora lida com atualizações parciais
	// O tipo request.WorkInfoUpdateRequest ainda vem do caminho antigo.
	// Ele será movido para src/controller/workinfo/request/ quando reorganizarmos o controller.
	// Por agora, manteremos o import original para ele.
	UpdateWorkInfoServices(
		userId string,
		updateRequest request.WorkInfoUpdateRequest,
	) (model.WorkInfoDomainInterface, *rest_err.RestErr)
}

// workInfoDomainService é a implementação da interface WorkInfoDomainService.
type workInfoDomainService struct {
	workInfoRepository repository_workinfo.WorkInfoRepository // Tipo ajustado
	userDomainService  service_user.UserDomainService         // Tipo ajustado
}

// NewWorkInfoDomainService cria uma nova instância de WorkInfoDomainService.
func NewWorkInfoDomainService(
	workInfoRepository repository_workinfo.WorkInfoRepository, // Tipo ajustado
	userDomainService service_user.UserDomainService, // Tipo ajustado
) WorkInfoDomainService {
	return &workInfoDomainService{
		workInfoRepository: workInfoRepository,
		userDomainService:  userDomainService,
	}
}
