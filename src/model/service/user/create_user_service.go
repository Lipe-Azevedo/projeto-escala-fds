package user

import (
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/logger"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model/domain" // <<< IMPORT ATUALIZADO
	"go.uber.org/zap"
)

// UserDomainService e userRepository vêm de user_service.go e user_repository.go neste mesmo pacote (ou pacotes irmãos)

func (ud *userDomainService) CreateUserServices(
	userDomainReq domain.UserDomainInterface, // <<< domain.UserDomainInterface
) (domain.UserDomainInterface, *rest_err.RestErr) { // <<< domain.UserDomainInterface
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

	// A criptografia será feita aqui antes de salvar, usando o método do domain
	// userDomainReq já é uma instância de domain.UserDomainInterface
	// que tem o método EncryptPassword (que ainda usa MD5, será alterado em seguida).
	userDomainReq.EncryptPassword() // Chamada ao método da interface

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
