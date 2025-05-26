package user

import (
	"net/http"

	"github.com/Lipe-Azevedo/escala-fds/src/configuration/logger"
	"github.com/Lipe-Azevedo/escala-fds/src/configuration/validation"

	// Usando o alias user_request_dto para o pacote de request do usuário
	user_request_dto "github.com/Lipe-Azevedo/escala-fds/src/controller/user/request"
	"github.com/Lipe-Azevedo/escala-fds/src/model/domain"
	"github.com/Lipe-Azevedo/escala-fds/src/view"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller", zap.String("journey", "createUser"))

	var userRequest user_request_dto.UserRequest // <<< Usando o alias aqui
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error validating user info for creation", err, zap.String("journey", "createUser"))
		errRest := validation.ValidateUserError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	userTypeDomain := domain.UserType(userRequest.UserType)

	domain := domain.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userTypeDomain,
	)

	domainResult, serviceErr := uc.service.CreateUserServices(domain)
	if serviceErr != nil {
		logger.Error("Failed to call user creation service", serviceErr, zap.String("journey", "createUser"))
		c.JSON(serviceErr.Code, serviceErr)
		return
	}

	logger.Info("User created successfully via controller",
		zap.String("userId", domainResult.GetID()),
		zap.String("journey", "createUser"))

	c.JSON(http.StatusCreated, view.ConvertUserDomainToResponse(domainResult))
}
