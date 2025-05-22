package controller

import (
	"net/http"

	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/logger"
	// "github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model"          // Descomentar se usar model.UserTypeMaster
	// "github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/rest_err" // Descomentar se usar rest_err.NewForbiddenError
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/view"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (wc *workInfoControllerInterface) FindWorkInfoByUserId(c *gin.Context) {
	logger.Info("Init findWorkInfoByUserId controller",
		zap.String("journey", "findWorkInfoByUserId"))

	targetUserId := c.Param("userId")

	// TODO: Reativar/Confirmar verificação de permissão após implementar JWT/AuthN.
	// A lógica abaixo é um EXEMPLO de como você poderia restringir o acesso.
	// Se JWT não estiver implementado, c.GetString("userType") retornará ""
	// e a condição abaixo (se descomentada) pode não funcionar como esperado.
	// Comente o bloco if para testes sem JWT ou ajuste conforme necessário.

	/* // Exemplo de Lógica de Permissão (permanece comentado):
	actingUserID := c.GetString("userID")
	actingUserType := c.GetString("userType")

	if model.UserType(actingUserType) != model.UserTypeMaster && actingUserID != targetUserId {
		logger.Warn("Forbidden attempt to find work info for another user by non-master user (or auth not implemented)",
			zap.String("journey", "findWorkInfoByUserId"),
			zap.String("actingUserID", actingUserID),
			zap.String("targetUserID", targetUserId))
		restErr := rest_err.NewForbiddenError("You do not have permission to view this information or user type not identified")
		c.JSON(restErr.Code, restErr)
		return
	}
	*/

	workInfoDomain, err := wc.service.FindWorkInfoByUserIdServices(targetUserId)
	if err != nil {
		logger.Error("Error trying to call findWorkInfoByUserId service",
			err,
			zap.String("journey", "findWorkInfoByUserId"),
			zap.String("targetUserIdToFind", targetUserId))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("FindWorkInfoByUserId controller executed successfully",
		zap.String("foundUserId", targetUserId),
		zap.String("journey", "findWorkInfoByUserId"))

	c.JSON(http.StatusOK, view.ConvertWorkInfoDomainToResponse(workInfoDomain))
}
