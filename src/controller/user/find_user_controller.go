package user

import (
	"net/http"
	"net/mail" // Para FindUserByEmail

	"github.com/Lipe-Azevedo/escala-fds/src/configuration/logger"
	"github.com/Lipe-Azevedo/escala-fds/src/configuration/rest_err"

	// DTO de response do usuário, usado para declarar o tipo da slice em FindAllUsers
	user_response_dto "github.com/Lipe-Azevedo/escala-fds/src/controller/user/response"
	"github.com/Lipe-Azevedo/escala-fds/src/model/domain" // Para domain.UserTypeCollaborator
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

	userDomainVal, serviceErr := uc.service.FindUserByIDServices(userId) // uc.service é o UserDomainService
	if serviceErr != nil {
		logger.Error(
			"Failed to call find user by ID service", serviceErr,
			zap.String("journey", "findUserByID"),
			zap.String("userId", userId))
		c.JSON(serviceErr.Code, serviceErr)
		return
	}

	var workInfoDomainVal domain.WorkInfoDomainInterface // Pode ser nil
	if userDomainVal.GetUserType() == domain.UserTypeCollaborator {
		logger.Info("User is collaborator, attempting to fetch WorkInfo via controller",
			zap.String("userId", userId),
			zap.String("journey", "findUserByID"))
		// uc.workInfoService foi injetado no userControllerInterface
		wi, wiErr := uc.workInfoService.FindWorkInfoByUserIdServices(userId) // userId é o ID do usuário
		if wiErr != nil {
			// Se WorkInfo não for encontrado (404), não é um erro fatal para a busca do usuário.
			// A view.ConvertUserDomainToResponse vai lidar com workInfoDomainVal sendo nil.
			if wiErr.Code != http.StatusNotFound {
				// Para outros erros ao buscar WorkInfo, logamos mas não necessariamente bloqueamos a resposta do usuário.
				// Poderíamos optar por retornar um erro aqui se WorkInfo for crítico.
				logger.Error("Error fetching WorkInfo for collaborator in controller, proceeding without it", wiErr,
					zap.String("userId", userId),
					zap.String("journey", "findUserByID"))
			} else {
				logger.Warn("WorkInfo not found for collaborator, proceeding without it",
					zap.String("userId", userId),
					zap.String("journey", "findUserByID"))
			}
		}
		workInfoDomainVal = wi // Atribui wi (pode ser nil se não encontrado ou se houve erro não fatal)
	}

	logger.Info(
		"FindUserByID controller executed successfully",
		zap.String("journey", "findUserByID"),
		zap.String("userId", userDomainVal.GetID()))

	// Passa ambos os domínios para a função de conversão da view
	c.JSON(http.StatusOK, view.ConvertUserDomainToResponse(userDomainVal, workInfoDomainVal))
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

	userDomainVal, serviceErr := uc.service.FindUserByEmailServices(userEmail)
	if serviceErr != nil {
		logger.Error(
			"Failed to call find user by email service", serviceErr,
			zap.String("journey", "findUserByEmail"),
			zap.String("userEmail", userEmail))
		c.JSON(serviceErr.Code, serviceErr)
		return
	}

	var workInfoDomainVal domain.WorkInfoDomainInterface                                    // Pode ser nil
	if userDomainVal != nil && userDomainVal.GetUserType() == domain.UserTypeCollaborator { // Adicionado nil check para userDomainVal
		logger.Info("User is collaborator, attempting to fetch WorkInfo via controller (FindUserByEmail)",
			zap.String("userId", userDomainVal.GetID()), // Usar o ID do usuário encontrado
			zap.String("journey", "findUserByEmail"))
		wi, wiErr := uc.workInfoService.FindWorkInfoByUserIdServices(userDomainVal.GetID())
		if wiErr != nil {
			if wiErr.Code != http.StatusNotFound {
				logger.Error("Error fetching WorkInfo for collaborator in controller (FindUserByEmail), proceeding without it", wiErr,
					zap.String("userId", userDomainVal.GetID()),
					zap.String("journey", "findUserByEmail"))
			} else {
				logger.Warn("WorkInfo not found for collaborator (FindUserByEmail), proceeding without it",
					zap.String("userId", userDomainVal.GetID()),
					zap.String("journey", "findUserByEmail"))
			}
		}
		workInfoDomainVal = wi
	}

	logger.Info(
		"FindUserByEmail controller executed successfully",
		zap.String("journey", "findUserByEmail"),
		zap.String("userEmail", userDomainVal.GetEmail())) // Assumindo userDomainVal não é nil aqui
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

	var userResponses []user_response_dto.UserResponse
	for _, userDomainSingle := range userDomains {
		// Para FindAllUsers, geralmente não anexamos WorkInfo para cada usuário
		// para evitar N+1 queries, a menos que explicitamente solicitado.
		// Se precisarmos, a lógica seria similar à de FindUserByID, dentro deste loop.
		var singleWorkInfoDomain domain.WorkInfoDomainInterface
		if userDomainSingle.GetUserType() == domain.UserTypeCollaborator {
			// Exemplo de como poderia ser (CUIDADO COM N+1 QUERIES):
			// wi, _ := uc.workInfoService.FindWorkInfoByUserIdServices(userDomainSingle.GetID())
			// singleWorkInfoDomain = wi
			// Por ora, deixamos nil.
		}
		userResponses = append(userResponses, view.ConvertUserDomainToResponse(userDomainSingle, singleWorkInfoDomain))
	}

	logger.Info("FindAllUsers controller executed successfully",
		zap.Int("count", len(userResponses)),
		zap.String("journey", "findAllUsers"))

	c.JSON(http.StatusOK, userResponses)
}
