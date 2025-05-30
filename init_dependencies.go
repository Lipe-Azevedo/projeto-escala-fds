package main

import (
	"os" // NOVO IMPORT para os.Getenv

	controller_comment "github.com/Lipe-Azevedo/escala-fds/src/controller/comment"
	controller_swap "github.com/Lipe-Azevedo/escala-fds/src/controller/swap"
	controller_user "github.com/Lipe-Azevedo/escala-fds/src/controller/user"
	controller_workinfo "github.com/Lipe-Azevedo/escala-fds/src/controller/workinfo"

	repo_comment "github.com/Lipe-Azevedo/escala-fds/src/model/repository/comment"
	repo_swap "github.com/Lipe-Azevedo/escala-fds/src/model/repository/swap"
	repo_user "github.com/Lipe-Azevedo/escala-fds/src/model/repository/user"
	repo_workinfo "github.com/Lipe-Azevedo/escala-fds/src/model/repository/workinfo"

	service_comment "github.com/Lipe-Azevedo/escala-fds/src/model/service/comment"
	service_swap "github.com/Lipe-Azevedo/escala-fds/src/model/service/swap"
	service_user "github.com/Lipe-Azevedo/escala-fds/src/model/service/user"
	service_workinfo "github.com/Lipe-Azevedo/escala-fds/src/model/service/workinfo"

	"go.mongodb.org/mongo-driver/mongo"
)

func initDependencies(
	database *mongo.Database,
) (
	controller_user.UserControllerInterface,
	controller_workinfo.WorkInfoControllerInterface,
	controller_swap.SwapControllerInterface,
	controller_comment.CommentControllerInterface,
) {
	// Obter a chave secreta do ambiente.
	// main.go já verificou se não está vazia.
	jwtSecret := os.Getenv("JWT_SECRET_KEY")

	// Repositories
	userRepo := repo_user.NewUserRepository(database)
	workInfoRepo := repo_workinfo.NewWorkInfoRepository(database)
	swapRepo := repo_swap.NewSwapRepository(database)
	commentRepo := repo_comment.NewCommentRepository(database)

	// Services
	userService := service_user.NewUserDomainService(userRepo, jwtSecret) // Passa jwtSecret
	workInfoService := service_workinfo.NewWorkInfoDomainService(workInfoRepo, userService)
	swapService := service_swap.NewSwapDomainService(swapRepo)
	commentService := service_comment.NewCommentDomainService(commentRepo, userService)

	// Controllers
	userController := controller_user.NewUserControllerInterface(userService, workInfoService)
	workInfoController := controller_workinfo.NewWorkInfoControllerInterface(workInfoService)
	swapController := controller_swap.NewSwapControllerInterface(swapService)
	commentController := controller_comment.NewCommentControllerInterface(commentService)

	return userController, workInfoController, swapController, commentController
}
