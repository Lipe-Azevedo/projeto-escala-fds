package routes

import (
	// User (já reorganizado)
	controller_user "github.com/Lipe-Azevedo/meu-primeio-crud-go/src/controller/user"
	// WorkInfo (NOVO IMPORT)
	controller_workinfo "github.com/Lipe-Azevedo/meu-primeio-crud-go/src/controller/workinfo"

	// Swap (AINDA USA O ANTIGO CAMINHO)
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(
	r *gin.RouterGroup,
	userController controller_user.UserControllerInterface,
	workInfoController controller_workinfo.WorkInfoControllerInterface, // <<< Tipo ajustado
	swapController controller.SwapControllerInterface, // Mantido
) {
	// Rotas de User
	userRoutes := r.Group("/users")
	{
		userRoutes.POST("", userController.CreateUser)
		userRoutes.GET("", userController.FindAllUsers)
		userRoutes.GET("/:userId", userController.FindUserByID)
		userRoutes.GET("/email/:userEmail", userController.FindUserByEmail)
		userRoutes.PUT("/:userId", userController.UpdateUser)
		userRoutes.DELETE("/:userId", userController.DeleteUser)
	}

	// Rotas de WorkInfo
	// As rotas em si não mudam, apenas o tipo do controller que as manipula.
	workInfoRoutes := r.Group("/users/:userId/work-info")
	{
		workInfoRoutes.POST("", workInfoController.CreateWorkInfo)
		workInfoRoutes.GET("", workInfoController.FindWorkInfoByUserId)
		workInfoRoutes.PUT("", workInfoController.UpdateWorkInfo)
	}

	// Rotas de Swap (mantendo nomenclatura original)
	swapRoutes := r.Group("/shift-swap")
	{
		swapRoutes.POST("", swapController.CreateSwap)
		swapRoutes.GET("/:id", swapController.FindSwapByID)
		swapRoutes.PUT("/:id/status", swapController.UpdateSwapStatus)
	}
}
