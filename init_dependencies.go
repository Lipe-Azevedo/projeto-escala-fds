package main

import (
	// User (já reorganizado)
	controller_user "github.com/Lipe-Azevedo/meu-primeio-crud-go/src/controller/user"
	repository_user "github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model/repository/user"
	service_user "github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model/service/user"

	// WorkInfo (NOVOS IMPORTS)
	controller_workinfo "github.com/Lipe-Azevedo/meu-primeio-crud-go/src/controller/workinfo"
	repository_workinfo "github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model/repository/workinfo"
	service_workinfo "github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model/service/workinfo"

	// Swap (AINDA USA OS ANTIGOS CAMINHOS - SERÃO ATUALIZADOS DEPOIS)
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/controller"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model/repository"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model/service"

	"go.mongodb.org/mongo-driver/mongo"
)

func initDependencies(
	database *mongo.Database,
) (
	controller_user.UserControllerInterface,
	controller_workinfo.WorkInfoControllerInterface, // <<< Tipo ajustado
	controller.SwapControllerInterface, // Mantido por enquanto
) {
	// User
	userRepo := repository_user.NewUserRepository(database)
	userService := service_user.NewUserDomainService(userRepo) // userService é do tipo service_user.UserDomainService
	userController := controller_user.NewUserControllerInterface(userService)

	// WorkInfo
	workInfoRepo := repository_workinfo.NewWorkInfoRepository(database)
	// NewWorkInfoDomainService espera um service_user.UserDomainService, que é o tipo de userService.
	workInfoService := service_workinfo.NewWorkInfoDomainService(workInfoRepo, userService)
	workInfoController := controller_workinfo.NewWorkInfoControllerInterface(workInfoService)

	// Swap (Mantendo a inicialização antiga por enquanto)
	swapRepo := repository.NewSwapRepository(database)
	swapService := service.NewSwapDomainService(swapRepo)
	swapController := controller.NewSwapControllerInterface(swapService)

	return userController, workInfoController, swapController
}
