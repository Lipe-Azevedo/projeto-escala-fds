package main

import (
	// User (já reorganizado)
	controller_user "github.com/Lipe-Azevedo/meu-primeio-crud-go/src/controller/user"
	repository_user "github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model/repository/user"
	service_user "github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model/service/user"

	// WorkInfo (já reorganizado)
	controller_workinfo "github.com/Lipe-Azevedo/meu-primeio-crud-go/src/controller/workinfo"
	repository_workinfo "github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model/repository/workinfo"
	service_workinfo "github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model/service/workinfo"

	// Swap (já reorganizado)
	controller_swap "github.com/Lipe-Azevedo/meu-primeio-crud-go/src/controller/swap"
	repository_swap "github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model/repository/swap"
	service_swap "github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model/service/swap"

	"go.mongodb.org/mongo-driver/mongo"
)

func initDependencies(
	database *mongo.Database,
) (
	controller_user.UserControllerInterface,
	controller_workinfo.WorkInfoControllerInterface,
	controller_swap.SwapControllerInterface,
) {
	// User
	userRepo := repository_user.NewUserRepository(database)
	userService := service_user.NewUserDomainService(userRepo)
	userController := controller_user.NewUserControllerInterface(userService)

	// WorkInfo
	workInfoRepo := repository_workinfo.NewWorkInfoRepository(database)
	workInfoService := service_workinfo.NewWorkInfoDomainService(workInfoRepo, userService)
	workInfoController := controller_workinfo.NewWorkInfoControllerInterface(workInfoService)

	// Swap
	swapRepo := repository_swap.NewSwapRepository(database)
	swapService := service_swap.NewSwapDomainService(swapRepo)
	swapController := controller_swap.NewSwapControllerInterface(swapService)

	return userController, workInfoController, swapController
}
