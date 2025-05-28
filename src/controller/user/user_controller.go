package user // << GARANTA QUE O PACOTE É 'user'

import (
	// Imports para os serviços que o controller de usuário utiliza
	service_user "github.com/Lipe-Azevedo/escala-fds/src/model/service/user"
	service_workinfo "github.com/Lipe-Azevedo/escala-fds/src/model/service/workinfo" // Adicionado para buscar WorkInfo
	"github.com/gin-gonic/gin"
)

// UserControllerInterface define a interface para os controllers de usuário.
// Todos os métodos que são chamados em routes.go devem estar aqui.
type UserControllerInterface interface {
	CreateUser(c *gin.Context)
	FindUserByID(c *gin.Context)
	FindUserByEmail(c *gin.Context)
	FindAllUsers(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	// LoginUser(c *gin.Context) // Para o futuro endpoint de login
}

// userControllerInterface é a implementação da interface UserControllerInterface.
// Os métodos (handlers) estão em arquivos separados neste mesmo pacote 'user'.
type userControllerInterface struct {
	service         service_user.UserDomainService
	workInfoService service_workinfo.WorkInfoDomainService // Adicionado para buscar WorkInfo
}

// NewUserControllerInterface cria uma nova instância de UserControllerInterface.
func NewUserControllerInterface(
	userService service_user.UserDomainService,
	workInfoService service_workinfo.WorkInfoDomainService, // Adicionado para buscar WorkInfo
) UserControllerInterface {
	return &userControllerInterface{
		service:         userService,
		workInfoService: workInfoService,
	}
}
