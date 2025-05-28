package main

import (
	controller_comment "github.com/Lipe-Azevedo/escala-fds/src/controller/comment" // NOVO IMPORT
	controller_swap "github.com/Lipe-Azevedo/escala-fds/src/controller/swap"
	controller_user "github.com/Lipe-Azevedo/escala-fds/src/controller/user"
	controller_workinfo "github.com/Lipe-Azevedo/escala-fds/src/controller/workinfo"

	repo_comment "github.com/Lipe-Azevedo/escala-fds/src/model/repository/comment" // NOVO IMPORT
	repo_swap "github.com/Lipe-Azevedo/escala-fds/src/model/repository/swap"
	repo_user "github.com/Lipe-Azevedo/escala-fds/src/model/repository/user"
	repo_workinfo "github.com/Lipe-Azevedo/escala-fds/src/model/repository/workinfo"

	service_comment "github.com/Lipe-Azevedo/escala-fds/src/model/service/comment" // NOVO IMPORT
	service_swap "github.com/Lipe-Azevedo/escala-fds/src/model/service/swap"
	service_user "github.com/Lipe-Azevedo/escala-fds/src/model/service/user"
	service_workinfo "github.com/Lipe-Azevedo/escala-fds/src/model/service/workinfo"

	"go.mongodb.org/mongo-driver/mongo"
)

// initDependencies agora também retorna CommentControllerInterface
func initDependencies(
	database *mongo.Database,
) (
	controller_user.UserControllerInterface,
	controller_workinfo.WorkInfoControllerInterface,
	controller_swap.SwapControllerInterface,
	controller_comment.CommentControllerInterface, // NOVO RETORNO
) {
	// Repositories
	userRepo := repo_user.NewUserRepository(database)
	workInfoRepo := repo_workinfo.NewWorkInfoRepository(database)
	swapRepo := repo_swap.NewSwapRepository(database)
	commentRepo := repo_comment.NewCommentRepository(database) // NOVO REPOSITÓRIO

	// Services
	// UserService não depende de outros serviços nesta configuração
	userService := service_user.NewUserDomainService(userRepo)
	// WorkInfoService depende de UserService (para validar a existência do usuário, por exemplo)
	workInfoService := service_workinfo.NewWorkInfoDomainService(workInfoRepo, userService)
	// SwapService não depende de outros serviços nesta configuração
	swapService := service_swap.NewSwapDomainService(swapRepo)
	// CommentService depende de UserService (para validar colaborador e autor)
	commentService := service_comment.NewCommentDomainService(commentRepo, userService) // NOVO SERVIÇO

	// Controllers
	userController := controller_user.NewUserControllerInterface(userService, workInfoService)
	workInfoController := controller_workinfo.NewWorkInfoControllerInterface(workInfoService)
	swapController := controller_swap.NewSwapControllerInterface(swapService)
	commentController := controller_comment.NewCommentControllerInterface(commentService) // NOVO CONTROLLER

	return userController, workInfoController, swapController, commentController // RETORNO ATUALIZADO
}
