// src/controller/routes/routes.go
package routes

import (
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(
	r *gin.RouterGroup,
	userController controller.UserControllerInterface,
	workInfoController controller.WorkInfoControllerInterface,
	swapController controller.SwapControllerInterface, // Usando o nome original da interface
) {
	// Rotas de User
	userRoutes := r.Group("/users")
	{
		userRoutes.POST("", userController.CreateUser)
		userRoutes.GET("", userController.FindAllUsers) // Rota para listar todos os usuários
		userRoutes.GET("/:userId", userController.FindUserByID)
		userRoutes.GET("/email/:userEmail", userController.FindUserByEmail)
		userRoutes.PUT("/:userId", userController.UpdateUser)
		userRoutes.DELETE("/:userId", userController.DeleteUser)
	}

	// Rotas de WorkInfo
	workInfoRoutes := r.Group("/users/:userId/work-info")
	{
		workInfoRoutes.POST("", workInfoController.CreateWorkInfo)
		workInfoRoutes.GET("", workInfoController.FindWorkInfoByUserId)
		workInfoRoutes.PUT("", workInfoController.UpdateWorkInfo)
	}

	// Rotas de Swap (usando a nomenclatura original, pois a Fase 3 não foi implementada)
	swapRoutes := r.Group("/shift-swap") // Usando o prefixo original
	{
		// Usando os nomes de handler originais do swapController
		swapRoutes.POST("", swapController.CreateSwap)
		swapRoutes.GET("/:id", swapController.FindSwapByID)
		swapRoutes.PUT("/:id/status", swapController.UpdateSwapStatus)
		// As rotas para FindPendingSwaps e FindSwapsByUser, que usariam a nomenclatura "Swap",
		// só seriam adicionadas e funcionariam corretamente após a Fase 3.
	}
}
