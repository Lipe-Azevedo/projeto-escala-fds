package user

import (
	"net/http"

	"github.com/Lipe-Azevedo/escala-fds/src/configuration/logger" // Import para rest_err.RestErr
	"github.com/Lipe-Azevedo/escala-fds/src/configuration/validation"
	user_response_dto "github.com/Lipe-Azevedo/escala-fds/src/controller/user/response"
	"github.com/Lipe-Azevedo/escala-fds/src/model/domain"
	"github.com/Lipe-Azevedo/escala-fds/src/view"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// LoginRequest DTO específico para login
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// LoginResponse DTO para resposta do login, incluindo o token e dados do usuário
type LoginResponse struct {
	Token string                         `json:"token"`
	User  user_response_dto.UserResponse `json:"user"`
}

func (uc *userControllerInterface) LoginUser(c *gin.Context) {
	logger.Info("Init LoginUser controller", zap.String("journey", "loginUser"))

	var loginRequest LoginRequest
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		logger.Error("Error validating login request", err, zap.String("journey", "loginUser"))
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	token, userDomain, serviceErr := uc.service.LoginUserServices(loginRequest.Email, loginRequest.Password)
	if serviceErr != nil {
		// O serviço de login já loga os detalhes do erro.
		// Apenas retornamos o erro para o cliente.
		c.JSON(serviceErr.Code, serviceErr)
		return
	}

	logger.Info("User logged in successfully via controller",
		zap.String("userId", userDomain.GetID()),
		zap.String("journey", "loginUser"))

	var workInfoDomainVal domain.WorkInfoDomainInterface
	if userDomain.GetUserType() == domain.UserTypeCollaborator {
		wi, wiErr := uc.workInfoService.FindWorkInfoByUserIdServices(userDomain.GetID())
		// A verificação wiErr != nil é crucial.
		// Se wiErr NÃO for nil E o código do erro NÃO for StatusNotFound, então logamos o aviso.
		if wiErr != nil && wiErr.Code != http.StatusNotFound {
			// CORREÇÃO APLICADA AQUI:
			logger.Warn("Error fetching WorkInfo for logged in collaborator, proceeding without it",
				zap.Error(wiErr), // Usar zap.Error() para encapsular o erro para o logger
				zap.String("userId", userDomain.GetID()),
				zap.String("journey", "loginUser"))
		}
		workInfoDomainVal = wi // wi pode ser nil se não encontrado ou se ocorreu um erro não fatal (como NotFound)
	}

	userResponseData := view.ConvertUserDomainToResponse(userDomain, workInfoDomainVal)

	c.JSON(http.StatusOK, LoginResponse{
		Token: token,
		User:  userResponseData,
	})
}
