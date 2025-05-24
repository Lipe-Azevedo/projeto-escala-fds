package routes

import (
	// Import para a nova interface do controller de User
	controller_user "github.com/Lipe-Azevedo/meu-primeio-crud-go/src/controller/user"

	// Imports para as interfaces antigas de WorkInfo e Swap (serão ajustados depois)
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(
	r *gin.RouterGroup,
	userController controller_user.UserControllerInterface, // Tipo ajustado
	workInfoController controller.WorkInfoControllerInterface, // Mantido
	swapController controller.SwapControllerInterface, // Mantido
) {
	// Rotas de User
	userRoutes := r.Group("/users")
	{
		userRoutes.POST("", userController.CreateUser)
		userRoutes.GET("", userController.FindAllUsers) // Este handler está em controller/user/find_user_controller.go
		userRoutes.GET("/:userId", userController.FindUserByID)
		userRoutes.GET("/email/:userEmail", userController.FindUserByEmail)
		userRoutes.PUT("/:userId", userController.UpdateUser)
		userRoutes.DELETE("/:userId", userController.DeleteUser)
		// Futura rota de login: userRoutes.POST("/login", userController.LoginUser)
	}

	// Rotas de WorkInfo (permanecem como estão por enquanto)
	workInfoRoutes := r.Group("/users/:userId/work-info")
	{
		workInfoRoutes.POST("", workInfoController.CreateWorkInfo)
		workInfoRoutes.GET("", workInfoController.FindWorkInfoByUserId)
		workInfoRoutes.PUT("", workInfoController.UpdateWorkInfo)
	}

	// Rotas de Swap (permanecem como estão por enquanto)
	swapRoutes := r.Group("/shift-swap") // Mantido o nome "shift-swap" conforme original
	{
		swapRoutes.POST("", swapController.CreateSwap)
		swapRoutes.GET("/:id", swapController.FindSwapByID)
		swapRoutes.PUT("/:id/status", swapController.UpdateSwapStatus)
	}
}
