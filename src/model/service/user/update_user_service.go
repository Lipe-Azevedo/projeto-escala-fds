package user

import (
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/logger"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/rest_err"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) UpdateUserServices( // Renomeado de UpdateUser para UpdateUserServices
	userId string,
	userDomain model.UserDomainInterface, // Este userDomain é o NewUserUpdateDomain (só com nome/senha)
) *rest_err.RestErr {

	logger.Info(
		"Init UpdateUserServices", // Log atualizado
		zap.String("journey", "updateUser"),
		zap.String("userId", userId))

	// Validação: Verificar se o usuário existe antes de tentar atualizar.
	// existingUser, findErr := ud.FindUserByIDServices(userId)
	// if findErr != nil {
	//  logger.Error("User to update not found by service", findErr,
	//      zap.String("journey", "updateUser"),
	//      zap.String("userId", userId))
	//  return findErr
	// }

	// Se a senha foi fornecida para atualização, criptografá-la.
	// O userDomain que chega aqui é o resultado de model.NewUserUpdateDomain,
	// que pode ter a senha em texto plano se ela foi incluída na request.
	if userDomain.GetPassword() != "" {
		// Precisamos garantir que estamos trabalhando com a interface completa para chamar EncryptPassword.
		// Se userDomain já é um *model.userDomain, isso funciona.
		// No entanto, para ser mais seguro e explícito, poderíamos buscar o usuário
		// e então aplicar as atualizações, ou ter certeza que NewUserUpdateDomain
		// retorna um tipo que implementa EncryptPassword (o que ele faz, implicitamente).

		// Vamos assumir que o userDomain passado pode ter sua senha criptografada diretamente.
		// Se userDomain fosse apenas uma interface com GetName/GetPassword, teríamos um problema.
		// Mas como NewUserUpdateDomain retorna um *userDomain, está OK.
		userDomain.EncryptPassword()
	}

	err := ud.userRepository.UpdateUser(userId, userDomain)
	if err != nil {
		logger.Error(
			"Error trying to call repository to update user", // Mensagem mais específica
			err,
			zap.String("journey", "updateUser"),
			zap.String("userId", userId),
		)
		return err
	}

	logger.Info(
		"UpdateUserServices executed successfully", // Log atualizado
		zap.String("userId", userId),
		zap.String("journey", "updateUser"),
	)
	return nil
}
