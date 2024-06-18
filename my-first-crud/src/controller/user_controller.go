package controller

import (
	"github.com/arturbaccarin/go-my-first-crud/src/model/service"
	"github.com/gin-gonic/gin"
)

func NewUserControllerInterface(service service.UserDomainService) UserControllerInterface {
	return &userControllerInterface{
		service: service,
	}
}

type UserControllerInterface interface {
	FindUserById(c *gin.Context)
	FindUserByEmail(c *gin.Context)

	DeleteUserById(c *gin.Context)
	CreateUser(c *gin.Context)
	UpdateUserById(c *gin.Context)
	LoginUser(c *gin.Context)
}

type userControllerInterface struct {
	service service.UserDomainService
}
