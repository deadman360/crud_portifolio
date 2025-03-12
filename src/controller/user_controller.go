package controller

import (
	"github.com/deadman360/crud_portifolio/src/model/service"
	"github.com/gin-gonic/gin"
)

func NewUserControllerInterface(service service.UserDomainService) UserControllerInterface {
	return &userControllerInterface{
		service: service,
	}
}

type UserControllerInterface interface {
	CreateUser(*gin.Context)
	DeleteUserById(*gin.Context)
	FindUserById(*gin.Context)
	FindUserByEmail(*gin.Context)
	UpdateUserById(*gin.Context)
}
type userControllerInterface struct {
	service service.UserDomainService
}
