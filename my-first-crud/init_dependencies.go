package main

import (
	"github.com/arturbaccarin/go-my-first-crud/src/controller"
	"github.com/arturbaccarin/go-my-first-crud/src/model/repository"
	"github.com/arturbaccarin/go-my-first-crud/src/model/service"
	"go.mongodb.org/mongo-driver/mongo"
)

func initDependencies(database *mongo.Database) controller.UserControllerInterface {
	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)

	return controller.NewUserControllerInterface(service)
}
