package controller

import (
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model/service"
	"github.com/gin-gonic/gin"
)

type SwapControllerInterface interface {
	CreateSwap(c *gin.Context)
	FindSwapByID(c *gin.Context)
	UpdateSwapStatus(c *gin.Context)
}

type swapControllerInterface struct {
	service service.SwapDomainService
}

func NewSwapControllerInterface(
	serviceInterface service.SwapDomainService,
) SwapControllerInterface {
	return &swapControllerInterface{
		service: serviceInterface,
	}
}
