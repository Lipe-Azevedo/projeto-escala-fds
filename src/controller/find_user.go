package controller

import (
	"net/http"
	"net/mail"

	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/logger"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/controller/model/response" // Esta importação precisa funcionar

	// Para model.UserTypeMaster no TODO
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/view"
	"github.com/gin-gonic/gin"

	// "go.mongodb.org/mongo-driver/bson/primitive" // Não usado diretamente aqui, mas poderia ser para validação de ID
	"go.uber.org/zap"
)

func (uc *userControllerInterface) FindUserByID(c *gin.Context) {
	logger.Info(
		"Init FindUserByID controller.",
		zap.String("journey", "findUserByID"),
	)

	userId := c.Param("userId")

	// Exemplo de validação de formato de ID (se for ObjectID Hex):
	// if _, err := primitive.ObjectIDFromHex(userId); err != nil {
	// 	logger.Error("Invalid userId format for FindUserByID", err, zap.String("journey", "findUserByID"), zap.String("userId", userId))
	// 	errorMessage := rest_err.NewBadRequestError("UserID is not a valid ID format")
	// 	c.JSON(errorMessage.Code, errorMessage)
	// 	return
	// }

	userDomain, err := uc.service.FindUserByIDServices(userId)
	if err != nil {
		logger.Error(
			"Error trying to call FindUserByIDServices.",
			err,
			zap.String("journey", "findUserByID"),
			zap.String("userId", userId)) // Adicionar ID ao log de erro
		c.JSON(err.Code, err)
		return
	}

	logger.Info(
		"FindUserByID controller executed successfully.",
		zap.String("journey", "findUserByID"),
		zap.String("userId", userDomain.GetID())) // Log com o ID do usuário encontrado
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(
		userDomain,
	))
}

func (uc *userControllerInterface) FindUserByEmail(c *gin.Context) {
	logger.Info(
		"Init FindUserByEmail controller.",
		zap.String("journey", "findUserByEmail"),
	)

	userEmail := c.Param("userEmail")

	if _, err := mail.ParseAddress(userEmail); err != nil {
		logger.Error(
			"Error trying to validate userEmail.",
			err,
			zap.String("journey", "findUserByEmail"),
			zap.String("userEmail", userEmail)) // Adicionar email ao log de erro
		errorMessage := rest_err.NewBadRequestError(
			"UserEmail is not a valid email format")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := uc.service.FindUserByEmailServices(userEmail)
	if err != nil {
		logger.Error(
			"Error trying to call FindUserByEmailServices.",
			err,
			zap.String("journey", "findUserByEmail"),
			zap.String("userEmail", userEmail)) // Adicionar email ao log de erro
		c.JSON(err.Code, err)
		return
	}

	logger.Info(
		"FindUserByEmail controller executed successfully.",
		zap.String("journey", "findUserByEmail"),
		zap.String("userEmail", userDomain.GetEmail())) // Log com o email do usuário encontrado
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(
		userDomain,
	))
}

func (uc *userControllerInterface) FindAllUsers(c *gin.Context) {
	logger.Info("Init FindAllUsers controller.",
		zap.String("journey", "findAllUsers"))

	// TODO: Implementar verificação de permissão JWT.
	// Somente usuários 'master' devem poder listar todos os usuários.
	/*
	   actingUserType := c.GetString("userType")
	   if model.UserType(actingUserType) != model.UserTypeMaster {
	       logger.Warn("Forbidden attempt to list all users by non-master user (or auth not implemented)",
	           zap.String("journey", "findAllUsers"),
	           zap.String("actingUserID", c.GetString("userID")),
	           zap.String("actingUserType", actingUserType))
	       restErr := rest_err.NewForbiddenError("You do not have permission to perform this action or user type not identified")
	       c.JSON(restErr.Code, restErr)
	       return
	   }
	*/

	userDomains, serviceErr := uc.service.FindAllUsersServices()
	if serviceErr != nil {
		logger.Error("Error calling FindAllUsersServices", serviceErr,
			zap.String("journey", "findAllUsers"))
		c.JSON(serviceErr.Code, serviceErr)
		return
	}

	var userResponses []response.UserResponse
	for _, userDomain := range userDomains {
		userResponses = append(userResponses, view.ConvertDomainToResponse(userDomain))
	}

	logger.Info("FindAllUsers controller executed successfully",
		zap.Int("count", len(userResponses)),
		zap.String("journey", "findAllUsers"))

	c.JSON(http.StatusOK, userResponses)
}
