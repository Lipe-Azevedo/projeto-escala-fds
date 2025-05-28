package routes

import (
	"github.com/Lipe-Azevedo/escala-fds/src/controller/middleware" // NOVO IMPORT
	controller_swap "github.com/Lipe-Azevedo/escala-fds/src/controller/swap"
	controller_user "github.com/Lipe-Azevedo/escala-fds/src/controller/user"
	controller_workinfo "github.com/Lipe-Azevedo/escala-fds/src/controller/workinfo"
	"github.com/gin-gonic/gin"
)

func InitRoutes(
	r *gin.RouterGroup,
	userController controller_user.UserControllerInterface,
	workInfoController controller_workinfo.WorkInfoControllerInterface,
	swapController controller_swap.SwapControllerInterface,
) {
	// Rota de Login (não protegida por JWT inicialmente)
	r.POST("/login", userController.LoginUser)

	// Rota pública para criar usuário (se aplicável, senão mover para grupo protegido)
	// Se a criação de usuário for apenas para masters logados, esta rota deve estar dentro de '/api'
	r.POST("/users", userController.CreateUser) // Assumindo que pode ser pública por enquanto

	// Grupo de rotas protegidas por JWT
	api := r.Group("/api")
	api.Use(middleware.AuthMiddleware()) // Aplicar middleware de autenticação JWT a este grupo
	{
		userRoutes := api.Group("/users")
		{
			// Removido POST "" daqui, pois foi definido como público acima. Se for protegido, adicione aqui.
			userRoutes.GET("", userController.FindAllUsers)
			userRoutes.GET("/:userId", userController.FindUserByID)
			userRoutes.GET("/email/:userEmail", userController.FindUserByEmail)
			userRoutes.PUT("/:userId", userController.UpdateUser)
			userRoutes.DELETE("/:userId", userController.DeleteUser)
		}

		workInfoRoutes := api.Group("/workinfo")
		{
			workInfoRoutes.POST("/:userId", workInfoController.CreateWorkInfo)
			workInfoRoutes.GET("/:userId", workInfoController.FindWorkInfoByUserId)
			workInfoRoutes.PUT("/:userId", workInfoController.UpdateWorkInfo)
		}

		swapRoutes := api.Group("/swaps")
		{
			swapRoutes.POST("", swapController.CreateSwap)
			swapRoutes.GET("/:id", swapController.FindSwapByID)
			swapRoutes.PUT("/:id/status", swapController.UpdateSwapStatus)
			// Adicionar outras rotas de swap conforme necessário
		}
	}
}
