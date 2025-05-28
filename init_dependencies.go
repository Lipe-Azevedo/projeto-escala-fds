package main

// ... (imports como antes) ...
import (
	controller_user "github.com/Lipe-Azevedo/escala-fds/src/controller/user"
	repository_user "github.com/Lipe-Azevedo/escala-fds/src/model/repository/user"
	service_user "github.com/Lipe-Azevedo/escala-fds/src/model/service/user"

	controller_workinfo "github.com/Lipe-Azevedo/escala-fds/src/controller/workinfo"
	repository_workinfo "github.com/Lipe-Azevedo/escala-fds/src/model/repository/workinfo"
	service_workinfo "github.com/Lipe-Azevedo/escala-fds/src/model/service/workinfo"

	controller_swap "github.com/Lipe-Azevedo/escala-fds/src/controller/swap"
	repository_swap "github.com/Lipe-Azevedo/escala-fds/src/model/repository/swap"
	service_swap "github.com/Lipe-Azevedo/escala-fds/src/model/service/swap"

	"go.mongodb.org/mongo-driver/mongo"
)

func initDependencies(
	database *mongo.Database,
) (
	controller_user.UserControllerInterface,
	controller_workinfo.WorkInfoControllerInterface,
	controller_swap.SwapControllerInterface,
) {
	// Repositories
	userRepo := repository_user.NewUserRepository(database)
	workInfoRepo := repository_workinfo.NewWorkInfoRepository(database)
	swapRepo := repository_swap.NewSwapRepository(database)

	// Services
	userService := service_user.NewUserDomainService(userRepo)
	workInfoService := service_workinfo.NewWorkInfoDomainService(workInfoRepo, userService)
	swapService := service_swap.NewSwapDomainService(swapRepo)

	// Controllers
	// UserController agora recebe workInfoService também
	userController := controller_user.NewUserControllerInterface(userService, workInfoService)
	workInfoController := controller_workinfo.NewWorkInfoControllerInterface(workInfoService)
	swapController := controller_swap.NewSwapControllerInterface(swapService)

	return userController, workInfoController, swapController
}
