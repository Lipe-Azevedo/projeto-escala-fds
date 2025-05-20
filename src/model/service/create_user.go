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
		"Init createUser model",
		zap.String("journey", "createUser"),
	)

	user, _ := ud.FindUserByEmailServices(userDomain.GetEmail())
	if user != nil {
		return nil, rest_err.NewBadRequestError("Email is already regidtred in another account")
	}

	// // Validação para usuários master
	// if userDomain.GetUserType() == model.UserTypeMaster && userDomain.GetWorkInfo() != nil {
	// 	return nil, rest_err.NewBadRequestError("Master users cannot have work info")
	// }

	// // Validação para colaboradores
	// if userDomain.GetUserType() == model.UserTypeCollaborator {
	// 	if userDomain.GetWorkInfo() == nil {
	// 		return nil, rest_err.NewBadRequestError("Collaborators must have work info")
	// 	}

	// 	// Validações adicionais do WorkInfo podem ser adicionadas aqui
	// 	if userDomain.GetWorkInfo().SuperiorID == "" {
	// 		return nil, rest_err.NewBadRequestError("Collaborators must have a superior")
	// 	}
	// }

	userDomain.EncryptPassword()

	userDomainRepository, err := ud.userRepository.CreateUser(userDomain)
	if err != nil {
		logger.Error("Error trying to call repository", err, zap.String("journey", "createUser"))
		return nil, err
	}

	logger.Info("CreateUser service executed successfully",
		zap.String("userId", userDomainRepository.GetID()),
		zap.String("journey", "createUser"))

	return userDomainRepository, nil
}
