package main

import (
	// Imports para os novos pacotes de controller
	controller_user "github.com/Lipe-Azevedo/meu-primeio-crud-go/src/controller/user"
	// controller_workinfo "github.com/Lipe-Azevedo/meu-primeio-crud-go/src/controller/workinfo" // Será adicionado quando WorkInfo for reorganizado
	// controller_swap "github.com/Lipe-Azevedo/meu-primeio-crud-go/src/controller/swap" // Será adicionado quando Swap for reorganizado

	// Imports para os novos pacotes de repositório
	repository_user "github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model/repository/user"
	// repository_workinfo "github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model/repository/workinfo" // Será adicionado
	// repository_swap "github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model/repository/swap" // Será adicionado

	// Imports para os novos pacotes de serviço
	service_user "github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model/service/user"
	// service_workinfo "github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model/service/workinfo" // Será adicionado
	// service_swap "github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model/service/swap" // Será adicionado

	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/controller"       // Para interfaces de WorkInfo e Swap (temporário)
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model/repository" // Para interfaces de WorkInfo e Swap (temporário)
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model/service"    // Para interfaces de WorkInfo e Swap (temporário)

	"go.mongodb.org/mongo-driver/mongo"
)

// initDependencies agora retorna as interfaces específicas dos novos pacotes para User.
// As interfaces de WorkInfo e Swap ainda virão dos pacotes antigos até serem reorganizadas.
func initDependencies(
	database *mongo.Database,
) (
	controller_user.UserControllerInterface, // Ajustado
	controller.WorkInfoControllerInterface, // Mantido por enquanto
	controller.SwapControllerInterface, // Mantido por enquanto
) {
	// User
	userRepo := repository_user.NewUserRepository(database)
	userService := service_user.NewUserDomainService(userRepo)
	userController := controller_user.NewUserControllerInterface(userService)

	// WorkInfo (Mantendo a inicialização antiga por enquanto)
	// Quando WorkInfo for reorganizado, estas linhas serão atualizadas:
	workInfoRepo := repository.NewWorkInfoRepository(database) // Usando o construtor antigo
	// O WorkInfoService depende do UserService. O userService já é o novo tipo.
	// O NewWorkInfoDomainService precisará aceitar o novo tipo de userService.
	// Vamos assumir que a interface UserDomainService não mudou sua assinatura de métodos,
	// então o userService (novo) ainda deve ser compatível onde o antigo era esperado.
	workInfoService := service.NewWorkInfoDomainService(workInfoRepo, userService)   // Usando o construtor antigo
	workInfoController := controller.NewWorkInfoControllerInterface(workInfoService) // Usando o construtor antigo

	// Swap (Mantendo a inicialização antiga por enquanto)
	// Quando Swap for reorganizado, estas linhas serão atualizadas:
	swapRepo := repository.NewSwapRepository(database)                   // Usando o construtor antigo
	swapService := service.NewSwapDomainService(swapRepo)                // Usando o construtor antigo
	swapController := controller.NewSwapControllerInterface(swapService) // Usando o construtor antigo

	return userController, workInfoController, swapController
}
