package user

import (
	"net/http"

	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/logger"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/validation"

	// Ajustar o import para o request específico do usuário
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/controller/user/request"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/view" // View ainda é global
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller", zap.String("journey", "createUser"))

	var userRequest request.UserRequest // Usando request do pacote user/request
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error validating user info for creation", err, zap.String("journey", "createUser"))
		errRest := validation.ValidateUserError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	// Validação do UserType no controller antes de passar para o serviço
	userTypeDomain := model.UserType(userRequest.UserType)
	if userTypeDomain != model.UserTypeCollaborator && userTypeDomain != model.UserTypeMaster {
		logger.Error("Invalid user_type provided", nil,
			zap.String("journey", "createUser"),
			zap.String("userTypeReceived", userRequest.UserType))
		// O validador "oneof" já deve pegar isso, mas uma checagem explícita não faz mal.
		// Se o oneof já garante, esta validação pode ser redundante.
		// No entanto, a conversão para model.UserType é boa.
	}

	domain := model.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userTypeDomain,
	)

	// O serviço cuidará da lógica de criptografia e verificação de email duplicado.
	domainResult, serviceErr := uc.service.CreateUserServices(domain)
	if serviceErr != nil {
		// O serviço já deve logar os detalhes do erro. O controller loga a falha na chamada do serviço.
		logger.Error("Failed to call user creation service", serviceErr, zap.String("journey", "createUser"))
		c.JSON(serviceErr.Code, serviceErr)
		return
	}

	logger.Info("User created successfully via controller",
		zap.String("userId", domainResult.GetID()),
		zap.String("journey", "createUser"))

	// O view.ConvertDomainToResponse usará o UserResponse de src/controller/user/response/
	// após o ajuste dos imports no pacote view.
	c.JSON(http.StatusCreated, view.ConvertDomainToResponse(domainResult)) // Alterado para StatusCreated
}
