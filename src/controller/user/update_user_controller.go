package user

import (
	"net/http"

	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/logger"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/validation"

	// Ajustar o import para o request específico do usuário
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/controller/user/request"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) UpdateUser(c *gin.Context) {
	logger.Info(
		"Init UpdateUser controller",
		zap.String("journey", "updateUser"),
	)
	var userRequest request.UserUpdateRequest // Usando request do pacote user/request
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error(
			"Error trying to validate user info for update",
			err,
			zap.String("journey", "updateUser"))
		errRest := validation.ValidateUserError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	userId := c.Param("userId")
	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		logger.Error("Invalid userId format for update in controller", err,
			zap.String("journey", "updateUser"),
			zap.String("userId", userId))
		restErr := rest_err.NewBadRequestError("Invalid userId format, must be a hex value")
		c.JSON(restErr.Code, restErr)
		return // Adicionado return
	}

	// TODO: (Pós-JWT) Lógica de Permissão:
	// - Master pode atualizar qualquer campo (exceto UserType, Email).
	// - Colaborador só pode atualizar seu próprio Name e Password.
	// actingUserID := c.GetString("userID")
	// actingUserType := c.GetString("userType")
	// if model.UserType(actingUserType) != model.UserTypeMaster && actingUserID != userId {
	//     restErr := rest_err.NewForbiddenError("You cannot update another user's information.")
	//     c.JSON(restErr.Code, restErr)
	//     return
	// }
	// Se for colaborador atualizando a si mesmo, ok.
	// Se for master, ok.

	// O UserUpdateRequest só tem Name e Password.
	domain := model.NewUserUpdateDomain(
		userRequest.Name,     // Será string vazia se não fornecido (omitempty)
		userRequest.Password, // Será string vazia se não fornecido (omitempty)
	)

	serviceErr := uc.service.UpdateUserServices(userId, domain) // Chamando UpdateUserServices
	if serviceErr != nil {
		logger.Error(
			"Failed to call user update service",
			serviceErr,
			zap.String("journey", "updateUser"),
			zap.String("userId", userId))
		c.JSON(serviceErr.Code, serviceErr)
		return
	}

	logger.Info(
		"UpdateUser controller executed successfully",
		zap.String("userId", userId),
		zap.String("journey", "updateUser"))

	c.Status(http.StatusOK)
}
