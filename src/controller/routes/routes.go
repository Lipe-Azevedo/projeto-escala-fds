package routes

import (
	controller_swap "github.com/Lipe-Azevedo/escala-fds/src/controller/swap"
	controller_user "github.com/Lipe-Azevedo/escala-fds/src/controller/user"
	controller_workinfo "github.com/Lipe-Azevedo/escala-fds/src/controller/workinfo"
	"github.com/gin-gonic/gin"
)

func InitRoutes(
	r *gin.RouterGroup,
	userController controller_user.UserControllerInterface, // Tipo correto
	workInfoController controller_workinfo.WorkInfoControllerInterface,
	swapController controller_swap.SwapControllerInterface,
) {
	// ... o restante do código das rotas ...
	userRoutes := r.Group("/users")
	{
		userRoutes.POST("", userController.CreateUser) // Chamada ao método
		userRoutes.GET("", userController.FindAllUsers)
		userRoutes.GET("/:userId", userController.FindUserByID)
		userRoutes.GET("/email/:userEmail", userController.FindUserByEmail)
		userRoutes.PUT("/:userId", userController.UpdateUser)
		userRoutes.DELETE("/:userId", userController.DeleteUser)
	}
	// ...
}
