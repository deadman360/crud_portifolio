package main

import (
	"github.com/deadman360/crud_portifolio/src/controller"
	"github.com/deadman360/crud_portifolio/src/controller/repository"
	"github.com/deadman360/crud_portifolio/src/model/service"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitDependecies(database *mongo.Database) controller.UserControllerInterface {
	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	userController := controller.NewUserControllerInterface(service)
	return userController
}
