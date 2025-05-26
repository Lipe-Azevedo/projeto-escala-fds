package user

import (
	"net/http"
	"net/mail"

	"github.com/Lipe-Azevedo/escala-fds/src/configuration/logger"
	"github.com/Lipe-Azevedo/escala-fds/src/configuration/rest_err"
	controller_response "github.com/Lipe-Azevedo/escala-fds/src/controller/user/response"
	"github.com/Lipe-Azevedo/escala-fds/src/view" // Import do pacote view
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	// "github.com/Lipe-Azevedo/escala-fds/src/model" // Removido se não usado diretamente para model.UserTypeMaster
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

	c.JSON(http.StatusOK, view.ConvertUserDomainToResponse(userDomain)) // <<< CHAMADA ATUALIZADA
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
	c.JSON(http.StatusOK, view.ConvertUserDomainToResponse(userDomain)) // <<< CHAMADA ATUALIZADA
}

func (uc *userControllerInterface) FindAllUsers(c *gin.Context) {
	logger.Info("Init FindAllUsers controller",
		zap.String("journey", "findAllUsers"))

	userDomains, serviceErr := uc.service.FindAllUsersServices()
	if serviceErr != nil {
		logger.Error("Failed to call find all users service", serviceErr,
			zap.String("journey", "findAllUsers"))
		c.JSON(serviceErr.Code, serviceErr)
		return
	}

	var userResponses []controller_response.UserResponse
	for _, userDomain := range userDomains {
		userResponses = append(userResponses, view.ConvertUserDomainToResponse(userDomain)) // <<< CHAMADA ATUALIZADA
	}

	logger.Info("FindAllUsers controller executed successfully",
		zap.Int("count", len(userResponses)),
		zap.String("journey", "findAllUsers"))

	c.JSON(http.StatusOK, userResponses)
}
