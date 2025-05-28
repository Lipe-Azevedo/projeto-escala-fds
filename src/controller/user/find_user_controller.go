package user

import (
	"net/http"
	"net/mail"

	"github.com/Lipe-Azevedo/escala-fds/src/configuration/logger"
	"github.com/Lipe-Azevedo/escala-fds/src/configuration/rest_err"
	user_response_dto "github.com/Lipe-Azevedo/escala-fds/src/controller/user/response"
	"github.com/Lipe-Azevedo/escala-fds/src/model/domain" // Import domain para domain.UserTypeCollaborator
	"github.com/Lipe-Azevedo/escala-fds/src/view"
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

	userDomainVal, serviceErr := uc.service.FindUserByIDServices(userId)
	if serviceErr != nil {
		logger.Error(
			"Failed to call find user by ID service", serviceErr,
			zap.String("journey", "findUserByID"),
			zap.String("userId", userId))
		c.JSON(serviceErr.Code, serviceErr)
		return
	}

	// Se o usuário for um colaborador, buscar seu WorkInfo
	var workInfoDomainVal domain.WorkInfoDomainInterface // Pode ser nil
	if userDomainVal.GetUserType() == domain.UserTypeCollaborator {
		logger.Info("User is collaborator, attempting to fetch WorkInfo via controller",
			zap.String("userId", userId),
			zap.String("journey", "findUserByID"))
		// uc.workInfoService foi injetado no userControllerInterface
		wi, wiErr := uc.workInfoService.FindWorkInfoByUserIdServices(userId)
		if wiErr != nil {
			if wiErr.Code == http.StatusNotFound { // Comparar com http.StatusNotFound ou o código específico do seu rest_err
				logger.Warn("WorkInfo not found for collaborator, will return user without WorkInfo",
					zap.String("userId", userId),
					zap.String("journey", "findUserByID"))
				// Não é um erro fatal para a busca do usuário, workInfoDomainVal permanecerá nil
			} else {
				// Outro erro ao buscar WorkInfo
				logger.Error("Error fetching WorkInfo for collaborator in controller", wiErr,
					zap.String("userId", userId),
					zap.String("journey", "findUserByID"))
				// Dependendo da política, você pode querer retornar um erro aqui ou apenas logar
				// e continuar sem o WorkInfo. Por enquanto, vamos apenas logar.
			}
		}
		workInfoDomainVal = wi // Atribui mesmo se wi for nil (após um NotFound, por exemplo)
	}

	logger.Info(
		"FindUserByID controller executed successfully",
		zap.String("journey", "findUserByID"),
		zap.String("userId", userDomainVal.GetID()))

	// Passa ambos os domínios para a função de conversão da view
	c.JSON(http.StatusOK, view.ConvertUserDomainToResponse(userDomainVal, workInfoDomainVal))
}

// FindUserByEmail e FindAllUsers não precisam necessariamente carregar WorkInfo por padrão.
// Se for necessário para FindUserByEmail, a lógica seria similar à de FindUserByID.
// Para FindAllUsers, carregar WorkInfo para todos os colaboradores pode ser custoso (N+1).

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

	userDomainVal, serviceErr := uc.service.FindUserByEmailServices(userEmail)
	if serviceErr != nil {
		logger.Error(
			"Failed to call find user by email service", serviceErr,
			zap.String("journey", "findUserByEmail"),
			zap.String("userEmail", userEmail))
		c.JSON(serviceErr.Code, serviceErr)
		return
	}

	var workInfoDomainVal domain.WorkInfoDomainInterface
	if userDomainVal != nil && userDomainVal.GetUserType() == domain.UserTypeCollaborator {
		wi, wiErr := uc.workInfoService.FindWorkInfoByUserIdServices(userDomainVal.GetID())
		if wiErr != nil && wiErr.Code != http.StatusNotFound {
			logger.Error("Error fetching WorkInfo for collaborator in FindUserByEmail", wiErr, zap.String("userId", userDomainVal.GetID()))
		}
		workInfoDomainVal = wi
	}

	logger.Info(
		"FindUserByEmail controller executed successfully",
		zap.String("journey", "findUserByEmail"),
		zap.String("userEmail", userDomainVal.GetEmail()))
	c.JSON(http.StatusOK, view.ConvertUserDomainToResponse(userDomainVal, workInfoDomainVal))
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

	// Para FindAllUsers, geralmente não se anexa WorkInfo para evitar N+1 queries,
	// a menos que haja paginação ou um parâmetro explícito para incluir.
	// O UserResponse DTO já lida com WorkInfo opcional.
	var userResponses []user_response_dto.UserResponse
	for _, userDomainSingle := range userDomains {
		// Não passamos workInfo aqui; ele será nil por padrão no UserResponse.
		// Se quiséssemos incluir, teríamos que buscar para cada colaborador.
		// Para simplicidade agora, não vamos incluir em FindAllUsers.
		// Se necessário, uma nova lógica de busca e anexação seria necessária aqui.
		var singleWorkInfoDomain domain.WorkInfoDomainInterface
		if userDomainSingle.GetUserType() == domain.UserTypeCollaborator {
			// Poderia buscar aqui, mas cuidado com N+1
			// wi, _ := uc.workInfoService.FindWorkInfoByUserIdServices(userDomainSingle.GetID())
			// singleWorkInfoDomain = wi
		}
		userResponses = append(userResponses, view.ConvertUserDomainToResponse(userDomainSingle, singleWorkInfoDomain))
	}

	logger.Info("FindAllUsers controller executed successfully",
		zap.Int("count", len(userResponses)),
		zap.String("journey", "findAllUsers"))

	c.JSON(http.StatusOK, userResponses)
}
