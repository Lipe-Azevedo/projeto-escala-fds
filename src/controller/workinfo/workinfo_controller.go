package workinfo

import (
	// Import para o novo pacote do serviço workinfo
	service_workinfo "github.com/Lipe-Azevedo/escala-fds/src/model/service/workinfo"
	"github.com/gin-gonic/gin"
)

// WorkInfoControllerInterface define a interface para os controllers de WorkInfo.
type WorkInfoControllerInterface interface {
	CreateWorkInfo(c *gin.Context)
	FindWorkInfoByUserId(c *gin.Context)
	UpdateWorkInfo(c *gin.Context)
}

// workInfoControllerInterface é a implementação da interface WorkInfoControllerInterface.
type workInfoControllerInterface struct {
	service service_workinfo.WorkInfoDomainService // Tipo ajustado
}

// NewWorkInfoControllerInterface cria uma nova instância de WorkInfoControllerInterface.
func NewWorkInfoControllerInterface(
	serviceInterface service_workinfo.WorkInfoDomainService, // Tipo ajustado
) WorkInfoControllerInterface {
	return &workInfoControllerInterface{
		service: serviceInterface,
	}
}
