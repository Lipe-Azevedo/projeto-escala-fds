package workinfo

import (
	"github.com/Lipe-Azevedo/escala-fds/src/configuration/rest_err"
	workinfo_request_dto "github.com/Lipe-Azevedo/escala-fds/src/controller/workinfo/request"

	// IMPORT ATUALIZADO: Agora importa do subpacote 'domain'
	"github.com/Lipe-Azevedo/escala-fds/src/model/domain"
	repository_workinfo "github.com/Lipe-Azevedo/escala-fds/src/model/repository/workinfo"
	service_user "github.com/Lipe-Azevedo/escala-fds/src/model/service/user"
)

// WorkInfoDomainService define a interface para os serviços de domínio de WorkInfo.
type WorkInfoDomainService interface {
	CreateWorkInfoServices(
		workInfoReqDomain domain.WorkInfoDomainInterface, // <<< USA domain.WorkInfoDomainInterface
	) (domain.WorkInfoDomainInterface, *rest_err.RestErr) // <<< USA domain.WorkInfoDomainInterface

	FindWorkInfoByUserIdServices(
		userId string,
	) (domain.WorkInfoDomainInterface, *rest_err.RestErr) // <<< USA domain.WorkInfoDomainInterface

	UpdateWorkInfoServices(
		userId string,
		updateRequest workinfo_request_dto.WorkInfoUpdateRequest,
	) (domain.WorkInfoDomainInterface, *rest_err.RestErr) // <<< USA domain.WorkInfoDomainInterface
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
