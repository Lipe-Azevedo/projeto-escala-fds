package user

import (
	"github.com/Lipe-Azevedo/escala-fds/src/configuration/logger"
	"github.com/Lipe-Azevedo/escala-fds/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/escala-fds/src/model/domain"
	"go.uber.org/zap"
)

// userDomainService struct e NewUserDomainService devem estar em user_service.go neste pacote.
// userRepository (interface) deve ser importada ou definida de forma que ud.userRepository seja válido.

func (ud *userDomainService) CreateUserServices(
	userDomainReq domain.UserDomainInterface,
) (domain.UserDomainInterface, *rest_err.RestErr) {
	logger.Info(
		"Init CreateUserServices",
		zap.String("journey", "createUser"),
		zap.String("email", userDomainReq.GetEmail()))

	existingUser, _ := ud.FindUserByEmailServices(userDomainReq.GetEmail())
	if existingUser != nil {
		logger.Warn("Attempt to create user with already registered email",
			zap.String("email", userDomainReq.GetEmail()),
			zap.String("journey", "createUser"))
		return nil, rest_err.NewConflictError("Email is already registered in another account.")
	}

	// A criptografia será feita aqui antes de salvar.
	// O método EncryptPassword() agora retorna um erro.
	if err := userDomainReq.EncryptPassword(); err != nil { // <<< MODIFICADO AQUI para tratar o erro
		logger.Error("Error encrypting password during user creation", err, // Adicionado 'err' ao log
			zap.String("journey", "createUser"),
			zap.String("email", userDomainReq.GetEmail()))
		// Retorna um erro interno do servidor, pois a falha na criptografia é um problema sério.
		return nil, rest_err.NewInternalServerError("Error processing user credentials")
	}

	userDomainRepository, repoErr := ud.userRepository.CreateUser(userDomainReq)
	if repoErr != nil {
		logger.Error("Error calling repository to create user", repoErr,
			zap.String("journey", "createUser"))
		return nil, repoErr
	}

	logger.Info("CreateUserServices executed successfully",
		zap.String("userId", userDomainRepository.GetID()),
		zap.String("journey", "createUser"))

	return userDomainRepository, nil
}
