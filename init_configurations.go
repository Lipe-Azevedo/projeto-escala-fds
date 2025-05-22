package main

import (
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/controller"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model/repository"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model/service"
	"go.mongodb.org/mongo-driver/mongo"
)

func initDependencies(
	database *mongo.Database,
) (controller.UserControllerInterface, controller.WorkInfoControllerInterface, controller.SwapControllerInterface) {
	// User
	userRepo := repository.NewUserRepository(database)
	userService := service.NewUserDomainService(userRepo) // userService já é criado
	userController := controller.NewUserControllerInterface(userService)

	// WorkInfo
	workInfoRepo := repository.NewWorkInfoRepository(database)
	// Passar userService para NewWorkInfoDomainService
	workInfoService := service.NewWorkInfoDomainService(workInfoRepo, userService) // Corrigido: NewWorkInfiDomainService e adicionado userService
	workInfoController := controller.NewWorkInfoControllerInterface(workInfoService)

	swapRepo := repository.NewSwapRepository(database)
	// Se SwapService precisar de userService, ele também seria passado aqui.
	swapService := service.NewSwapDomainService(swapRepo)
	swapController := controller.NewSwapControllerInterface(swapService)

	return userController, workInfoController, swapController
}
