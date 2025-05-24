package user

import (
	// Import ajustado para o novo pacote do serviço de usuário
	service_user "github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model/service/user"
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
	service service_user.UserDomainService // Tipo ajustado para service_user.UserDomainService
}

// NewUserControllerInterface cria uma nova instância de UserControllerInterface.
func NewUserControllerInterface(
	serviceInterface service_user.UserDomainService, // Tipo ajustado para service_user.UserDomainService
) UserControllerInterface {
	return &userControllerInterface{
		service: serviceInterface,
	}
}
