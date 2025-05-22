package controller

import (
	"net/http"

	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/logger"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/view"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	// Para controle de acesso mais fino, você poderia importar:
	// "github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model"
	// "github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/rest_err"
)

func (wc *workInfoControllerInterface) FindWorkInfoByUserId(c *gin.Context) {
	logger.Info("Init findWorkInfoByUserId controller",
		zap.String("journey", "findWorkInfoByUserId"))

	targetUserId := c.Param("userId") // ID do usuário cuja WorkInfo está sendo buscada

	// Possível Lógica de Permissão (Exemplo - pode ser adaptada ou removida):
	// actingUserID := c.GetString("userID")      // ID do usuário autenticado
	// actingUserType := c.GetString("userType") // Tipo do usuário autenticado
	//
	// if model.UserType(actingUserType) != model.UserTypeMaster && actingUserID != targetUserId {
	// 	logger.Warn("Forbidden attempt to find work info for another user by non-master user",
	// 		zap.String("journey", "findWorkInfoByUserId"),
	// 		zap.String("actingUserID", actingUserID),
	// 		zap.String("targetUserID", targetUserId))
	// 	restErr := rest_err.NewForbiddenError("You do not have permission to view this information")
	// 	c.JSON(restErr.Code, restErr)
	// 	return
	// }

	workInfoDomain, err := wc.service.FindWorkInfoByUserIdServices(targetUserId)
	if err != nil {
		logger.Error("Error trying to call findWorkInfoByUserId service",
			err, // err já é *rest_err.RestErr
			zap.String("journey", "findWorkInfoByUserId"),
			zap.String("targetUserIdToFind", targetUserId)) // Log específico do ID que falhou
		c.JSON(err.Code, err)
		return
	}

	logger.Info("FindWorkInfoByUserId controller executed successfully",
		zap.String("foundUserId", targetUserId), // Log com o ID buscado
		zap.String("journey", "findWorkInfoByUserId"))

	c.JSON(http.StatusOK, view.ConvertWorkInfoDomainToResponse(workInfoDomain))
}
