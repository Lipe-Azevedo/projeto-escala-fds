package user

import (
	"net/http"

	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/logger"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/rest_err"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) DeleteUser(c *gin.Context) {
	logger.Info(
		"Init DeleteUser controller",
		zap.String("journey", "deleteUser"))

	userId := c.Param("userId")
	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		logger.Error("Invalid userId format for deletion in controller", err,
			zap.String("journey", "deleteUser"),
			zap.String("userId", userId))
		restErr := rest_err.NewBadRequestError("Invalid userId format, must be a hex value")
		c.JSON(restErr.Code, restErr)
		return // Adicionado return
	}

	serviceErr := uc.service.DeleteUserServices(userId) // Chamando DeleteUserServices
	if serviceErr != nil {
		logger.Error(
			"Failed to call user deletion service",
			serviceErr,
			zap.String("journey", "deleteUser"),
			zap.String("userId", userId))
		c.JSON(serviceErr.Code, serviceErr)
		return
	}

	logger.Info(
		"DeleteUser controller executed successfully",
		zap.String("userId", userId),
		zap.String("journey", "deleteUser"))

	c.Status(http.StatusOK)
}
