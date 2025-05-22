package service

import (
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/logger"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) FindUserByIDServices(
	id string,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info(
		"Init FindUserByIDServices model.", // "Init findUserByIDServices model."
		zap.String("journey", "findUserByID"))

	return ud.userRepository.FindUserByID(id)
}

func (ud *userDomainService) FindUserByEmailServices(
	email string,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info(
		"Init FindUserByEmailServices.", // "Init findUserByEmail services."
		zap.String("journey", "findUserByEmail"))

	return ud.userRepository.FindUserByEmail(email)
}

// Novo método para buscar todos os usuários
func (ud *userDomainService) FindAllUsersServices() ([]model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info(
		"Init FindAllUsersServices.",
		zap.String("journey", "findAllUsers"))

	// Aqui poderiam entrar lógicas de negócio, como paginação,
	// ou filtros adicionais, se necessário no futuro.
	// Por ora, apenas repassa para o repositório.

	userDomains, err := ud.userRepository.FindAllUsers()
	if err != nil {
		logger.Error("Error calling repository for FindAllUsers", err,
			zap.String("journey", "findAllUsers"))
		return nil, err
	}

	logger.Info("FindAllUsersServices executed successfully",
		zap.Int("count", len(userDomains)),
		zap.String("journey", "findAllUsers"))
	return userDomains, nil
}
