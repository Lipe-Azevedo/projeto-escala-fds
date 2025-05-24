package user

import (
	"net/http"
	"net/mail"

	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/logger"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/rest_err"

	// Ajustar o import para o response específico do usuário se necessário, ou view cuidará disso.
	// O view.ConvertDomainToResponse usará o UserResponse de src/controller/user/response/
	// após o ajuste dos imports no pacote view.
	controller_response "github.com/Lipe-Azevedo/meu-primeio-crud-go/src/controller/user/response" // Para model.UserTypeMaster (se usado em permissões)
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/view"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) FindUserByID(c *gin.Context) {
	logger.Info(
		"Init FindUserByID controller",
		zap.String("journey", "findUserByID"),
	)

	userId := c.Param("userId")
	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		logger.Error("Invalid userId format for FindUserByID in controller", err,
			zap.String("journey", "findUserByID"),
			zap.String("userId", userId))
		errorMessage := rest_err.NewBadRequestError("UserID is not a valid ID format")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, serviceErr := uc.service.FindUserByIDServices(userId)
	if serviceErr != nil {
		logger.Error(
			"Failed to call find user by ID service", serviceErr,
			zap.String("journey", "findUserByID"),
			zap.String("userId", userId))
		c.JSON(serviceErr.Code, serviceErr)
		return
	}

	logger.Info(
		"FindUserByID controller executed successfully",
		zap.String("journey", "findUserByID"),
		zap.String("userId", userDomain.GetID()))

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))
}

func (uc *userControllerInterface) FindUserByEmail(c *gin.Context) {
	logger.Info(
		"Init FindUserByEmail controller",
		zap.String("journey", "findUserByEmail"),
	)

	userEmail := c.Param("userEmail")
	if _, err := mail.ParseAddress(userEmail); err != nil {
		logger.Error(
			"Invalid userEmail format for FindUserByEmail in controller", err,
			zap.String("journey", "findUserByEmail"),
			zap.String("userEmail", userEmail))
		errorMessage := rest_err.NewBadRequestError("UserEmail is not a valid email format")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, serviceErr := uc.service.FindUserByEmailServices(userEmail)
	if serviceErr != nil {
		logger.Error(
			"Failed to call find user by email service", serviceErr,
			zap.String("journey", "findUserByEmail"),
			zap.String("userEmail", userEmail))
		c.JSON(serviceErr.Code, serviceErr)
		return
	}

	logger.Info(
		"FindUserByEmail controller executed successfully",
		zap.String("journey", "findUserByEmail"),
		zap.String("userEmail", userDomain.GetEmail()))
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))
}

func (uc *userControllerInterface) FindAllUsers(c *gin.Context) {
	logger.Info("Init FindAllUsers controller",
		zap.String("journey", "findAllUsers"))

	// TODO: (Pós-JWT) Verificar permissão. Somente 'master'.
	// actingUserType := c.GetString("userType") // Virá do token JWT
	// if model.UserType(actingUserType) != model.UserTypeMaster {
	//     logger.Warn("Forbidden attempt to list all users by non-master user", ...)
	//     restErr := rest_err.NewForbiddenError("You do not have permission to perform this action.")
	//     c.JSON(restErr.Code, restErr)
	//     return
	// }

	userDomains, serviceErr := uc.service.FindAllUsersServices()
	if serviceErr != nil {
		logger.Error("Failed to call find all users service", serviceErr,
			zap.String("journey", "findAllUsers"))
		c.JSON(serviceErr.Code, serviceErr)
		return
	}

	// Convertendo cada domain para response.
	// O view.ConvertDomainToResponse será usado para cada item.
	// Se o UserResponse estivesse no pacote global, o import aqui seria diferente.
	// Como está em controller/user/response, o pacote view precisará conhecê-lo
	// ou teremos que fazer a conversão aqui.
	// Por ora, assumimos que view.ConvertDomainToResponse lida com UserResponse de alguma forma
	// ou que o pacote view será atualizado para importar de controller_response.UserResponse.
	// Vamos manter a conversão usando o view por enquanto.
	var userResponses []controller_response.UserResponse // Usando o response específico
	for _, userDomain := range userDomains {
		userResponses = append(userResponses, view.ConvertDomainToResponse(userDomain))
	}

	logger.Info("FindAllUsers controller executed successfully",
		zap.Int("count", len(userResponses)),
		zap.String("journey", "findAllUsers"))

	c.JSON(http.StatusOK, userResponses)
}
