package routes

import (
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(
	r *gin.RouterGroup,
	userController controller.UserControllerInterface,
	workInfoController controller.WorkInfoControllerInterface,
	swapController controller.SwapControllerInterface, // Mantendo original até Fase 3 de Swap
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
	workInfoRoutes := r.Group("/users/:userId/work-info")
	{
		workInfoRoutes.POST("", workInfoController.CreateWorkInfo) // Criação, usa WorkInfoRequest (campos string obrigatórios)
		workInfoRoutes.GET("", workInfoController.FindWorkInfoByUserId)
		workInfoRoutes.PUT("", workInfoController.UpdateWorkInfo) // Atualização (parcial), usa WorkInfoUpdateRequest (campos *string opcionais)
	}

	// Rotas de Swap (mantendo nomenclatura original)
	swapRoutes := r.Group("/shift-swap")
	{
		swapRoutes.POST("", swapController.CreateSwap)
		swapRoutes.GET("/:id", swapController.FindSwapByID)
		swapRoutes.PUT("/:id/status", swapController.UpdateSwapStatus)
	}
}
