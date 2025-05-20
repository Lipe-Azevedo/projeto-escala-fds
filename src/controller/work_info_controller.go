package controller

import (
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/model/service"
	"github.com/gin-gonic/gin"
)

func NewWorkInfoControllerInterface(
	serviceInterface service.WorkInfoDomainService,
) WorkInfoControllerInterface {
	return &workInfoControllerInterface{
		service: serviceInterface,
	}
}

type WorkInfoControllerInterface interface {
	FindWorkInfoByUserId(c *gin.Context)
	CreateWorkInfo(c *gin.Context)
	UpdateWorkInfo(c *gin.Context)
}

type workInfoControllerInterface struct {
	service service.WorkInfoDomainService
}
