package routes

import (
	controller_comment "github.com/Lipe-Azevedo/escala-fds/src/controller/comment" // NOVO IMPORT
	"github.com/Lipe-Azevedo/escala-fds/src/controller/middleware"
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
	commentController controller_comment.CommentControllerInterface, // NOVO PARÂMETRO
) {
	r.POST("/login", userController.LoginUser)
	r.POST("/users", userController.CreateUser)

	api := r.Group("/api")
	api.Use(middleware.AuthMiddleware())
	{
		userProtectedRoutes := api.Group("/users")
		{
			userProtectedRoutes.GET("", userController.FindAllUsers)
			userProtectedRoutes.GET("/:userId", userController.FindUserByID)
			userProtectedRoutes.GET("/email/:userEmail", userController.FindUserByEmail)
			userProtectedRoutes.PUT("/:userId", userController.UpdateUser)
			userProtectedRoutes.DELETE("/:userId", userController.DeleteUser)
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
		}

		// NOVAS ROTAS PARA COMENTÁRIOS
		commentRoutes := api.Group("/comments")
		{
			commentRoutes.POST("", commentController.CreateComment) // POST /api/comments (collabID no body)
			commentRoutes.GET("/:commentId", commentController.FindCommentByID)
			// GET /api/comments/collaborator/:collaboratorId/date/:dateString (YYYY-MM-DD)
			commentRoutes.GET("/collaborator/:collaboratorId/date/:dateString", commentController.FindCommentsByCollaboratorAndDate)
			// GET /api/comments/collaborator/:collaboratorId/range?startDate=YYYY-MM-DD&endDate=YYYY-MM-DD
			commentRoutes.GET("/collaborator/:collaboratorId/range", commentController.FindCommentsByCollaboratorForDateRange)
			commentRoutes.PUT("/:commentId", commentController.UpdateComment)
			commentRoutes.DELETE("/:commentId", commentController.DeleteComment)
		}
	}
}
