package user

import (
	"net/http"

	"github.com/Lipe-Azevedo/escala-fds/src/configuration/logger"
	"github.com/Lipe-Azevedo/escala-fds/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/escala-fds/src/configuration/validation"

	// Usando o alias user_request_dto para o pacote de request do usuário
	user_request_dto "github.com/Lipe-Azevedo/escala-fds/src/controller/user/request"
	"github.com/Lipe-Azevedo/escala-fds/src/model/domain"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) UpdateUser(c *gin.Context) {
	logger.Info(
		"Init UpdateUser controller",
		zap.String("journey", "updateUser"),
	)
	var userRequest user_request_dto.UserUpdateRequest // <<< Usando o alias aqui
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
		return
	}

	domain := domain.NewUserUpdateDomain(
		userRequest.Name,
		userRequest.Password,
	)

	serviceErr := uc.service.UpdateUserServices(userId, domain)
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
