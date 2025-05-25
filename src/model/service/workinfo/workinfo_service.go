package workinfo

import (
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/rest_err"
	// Import ajustado para o novo local do WorkInfoUpdateRequest, usando alias
	workinfo_request_dto "github.com/Lipe-Azevedo/meu-primeio-crud-go/src/controller/workinfo/request"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model"
	repository_workinfo "github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model/repository/workinfo"
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

	UpdateWorkInfoServices(
		userId string,
		updateRequest workinfo_request_dto.WorkInfoUpdateRequest, // <<< Tipo ajustado com alias
	) (model.WorkInfoDomainInterface, *rest_err.RestErr)
}

// workInfoDomainService é a implementação da interface WorkInfoDomainService.
type workInfoDomainService struct {
	workInfoRepository repository_workinfo.WorkInfoRepository
	userDomainService  service_user.UserDomainService
}

// NewWorkInfoDomainService cria uma nova instância de WorkInfoDomainService.
func NewWorkInfoDomainService(
	workInfoRepository repository_workinfo.WorkInfoRepository,
	userDomainService service_user.UserDomainService,
) WorkInfoDomainService {
	return &workInfoDomainService{
		workInfoRepository: workInfoRepository,
		userDomainService:  userDomainService,
	}
}
