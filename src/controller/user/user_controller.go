package user

import (
	// ATENÇÃO: O import para o serviço de usuário precisará ser ajustado
	// para "github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model/service/user"
	// após a reorganização da camada de serviço.
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model/service" // ESTE IMPORT SERÁ AJUSTADO
	"github.com/gin-gonic/gin"
)

// UserControllerInterface define a interface para os controllers de usuário.
// As implementações dos handlers (métodos) estarão em arquivos separados neste pacote.
type UserControllerInterface interface {
	CreateUser(c *gin.Context)
	FindUserByID(c *gin.Context)
	FindUserByEmail(c *gin.Context)
	FindAllUsers(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	// LoginUser(c *gin.Context) // Adicionaremos quando implementarmos JWT
}

// userControllerInterface é a implementação da interface UserControllerInterface.
type userControllerInterface struct {
	service service.UserDomainService // O tipo aqui será service_user.UserDomainService
}

// NewUserControllerInterface cria uma nova instância de UserControllerInterface.
func NewUserControllerInterface(
	serviceInterface service.UserDomainService, // O tipo aqui será service_user.UserDomainService
) UserControllerInterface {
	return &userControllerInterface{
		service: serviceInterface,
	}
}
