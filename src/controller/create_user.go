package controller

import (
	"net/http"

	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/logger"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/validation"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/controller/model/request"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/view"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller", zap.String("journey", "createUser"))

	var userRequest request.UserRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error validating user info", err, zap.String("journey", "createUser"))
		errRest := validation.ValidateUserError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	domain := model.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		model.UserType(userRequest.UserType),
	)

	// Aqui pode adicionar o WorkInfo para colaboradores
	// if domain.GetUserType() == model.UserTypeCollaborator {
	//     domain.SetWorkInfo(...)
	// }

	domainResult, err := uc.service.CreateUserServices(domain)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User created successfully",
		zap.String("userId", domainResult.GetID()),
		zap.String("journey", "createUser"))

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(domainResult))
}
