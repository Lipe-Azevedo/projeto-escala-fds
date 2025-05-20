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
	shiftSwapController controller.ShiftSwapControllerInterface,
) {
	// Rotas de User
	r.GET("/getUserByID/:userId", userController.FindUserByID)
	r.GET("/getUserByEmail/:userEmail", userController.FindUserByEmail)
	r.POST("/createUser", userController.CreateUser)
	r.PUT("/updateUser/:userId", userController.UpdateUser)
	r.DELETE("/deleteUser/:userId", userController.DeleteUser)

	// Rotas de WorkInfo
	r.GET("/workInfo/:userId", workInfoController.FindWorkInfoByUserId)
	r.POST("/workInfo/:userId", workInfoController.CreateWorkInfo)
	r.PUT("/workInfo/:userId", workInfoController.UpdateWorkInfo)

	// Rotas de ShiftSwap
	r.POST("/shift-swap", shiftSwapController.CreateShiftSwap)
	r.GET("/shift-swap/:id", shiftSwapController.FindShiftSwapByID)
	r.PUT("/shift-swap/:id/status", shiftSwapController.UpdateShiftSwapStatus)
}
