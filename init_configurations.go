package main

import (
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/controller"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model/repository"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model/service"
	"go.mongodb.org/mongo-driver/mongo"
)

func initDependencies(
	database *mongo.Database,
) (controller.UserControllerInterface, controller.WorkInfoControllerInterface, controller.ShiftSwapControllerInterface) {
	// User
	userRepo := repository.NewUserRepository(database)
	userService := service.NewUserDomainService(userRepo)
	userController := controller.NewUserControllerInterface(userService)

	// WorkInfo
	workInfoRepo := repository.NewWorkInfoRepository(database)
	workInfoService := service.NewWorkInfiDomainService(workInfoRepo)
	workInfoController := controller.NewWorkInfoControllerInterface(workInfoService)

	// ShiftSwap
	shiftSwapRepo := repository.NewShiftSwapRepository(database)
	shiftSwapService := service.NewShiftSwapDomainService(shiftSwapRepo)
	shiftSwapController := controller.NewShiftSwapControllerInterface(shiftSwapService)

	return userController, workInfoController, shiftSwapController
}
