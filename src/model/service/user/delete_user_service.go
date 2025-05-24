package user

import (
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/logger"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ud *userDomainService) DeleteUserServices( // Renomeado de DeleteUser para DeleteUserServices
	userId string) *rest_err.RestErr {
	logger.Info(
		"Init DeleteUserServices", // Log atualizado
		zap.String("journey", "deleteUser"),
		zap.String("userId", userId))

	// Validação: Verificar se o usuário existe antes de tentar deletar.
	// Embora o repositório já retorne NotFound, adicionar a verificação aqui pode ser útil para lógicas de serviço futuras.
	// Por ora, vamos manter a chamada direta ao repositório.
	// _, findErr := ud.FindUserByIDServices(userId)
	// if findErr != nil {
	//  logger.Error("User to delete not found by service", findErr,
	//      zap.String("journey", "deleteUser"),
	//      zap.String("userId", userId))
	//  return findErr // Retorna o erro de "não encontrado" do FindUserByIDServices
	// }

	err := ud.userRepository.DeleteUser(userId)
	if err != nil {
		logger.Error(
			"Error trying to call repository to delete user", // Mensagem de log mais específica
			err,
			zap.String("journey", "deleteUser"),
			zap.String("userId", userId),
		)
		return err
	}

	logger.Info(
		"DeleteUserServices executed successfully", // Log atualizado
		zap.String("userId", userId),
		zap.String("journey", "deleteUser"),
	)
	return nil
}
