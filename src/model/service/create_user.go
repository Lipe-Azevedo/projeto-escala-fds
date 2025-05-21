package service

import (
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/logger"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUserServices(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info(
		"Init CreateUserServices", // Corrigido "Init createUser model" para contexto de serviço
		zap.String("journey", "createUser"),
	)

	// Verifica se o e-mail já está registrado
	existingUser, _ := ud.FindUserByEmailServices(userDomain.GetEmail())
	if existingUser != nil {
		logger.Warn("Attempt to create user with already registered email", // Mudado de Error para Warn, pois é um erro de request
			zap.String("email", userDomain.GetEmail()),
			zap.String("journey", "createUser"))
		return nil, rest_err.NewBadRequestError("Email is already registered in another account")
	}

	// Validações específicas do UserType:
	// Conforme a descrição do projeto, WorkInfo é separado.
	// Usuários Master não devem ter WorkInfo. Isso é tratado implicitamente pelo fato de WorkInfo ser uma entidade separada
	// e não vinculada durante a criação do usuário aqui.
	// Colaboradores terão WorkInfo adicionado por um usuário Master através dos endpoints de WorkInfo.

	// Criptografa a senha antes de salvar
	userDomain.EncryptPassword()

	userDomainRepository, err := ud.userRepository.CreateUser(userDomain)
	if err != nil {
		// A camada de repositório já loga o erro específico.
		// Aqui logamos que a chamada ao repositório falhou.
		logger.Error("Error calling repository to create user", err, // err já é *rest_err.RestErr
			zap.String("journey", "createUser"))
		return nil, err
	}

	logger.Info("CreateUserServices executed successfully",
		zap.String("userId", userDomainRepository.GetID()),
		zap.String("journey", "createUser"))

	return userDomainRepository, nil
}
