package controller

import (
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model/service"
	"github.com/gin-gonic/gin"
)

type ShiftSwapControllerInterface interface {
	CreateShiftSwap(c *gin.Context)
	FindShiftSwapByID(c *gin.Context)
	UpdateShiftSwapStatus(c *gin.Context)
}

type shiftSwapControllerInterface struct {
	service service.ShiftSwapDomainService
}

func NewShiftSwapControllerInterface(
	serviceInterface service.ShiftSwapDomainService,
) ShiftSwapControllerInterface {
	return &shiftSwapControllerInterface{
		service: serviceInterface,
	}
}
