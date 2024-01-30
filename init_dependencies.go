package main

import (
	"github.com/dan-santos/go-user/src/controller"
	"github.com/dan-santos/go-user/src/model/repository"
	"github.com/dan-santos/go-user/src/model/service"
	"go.mongodb.org/mongo-driver/mongo"
)

func initDependencies(database *mongo.Database)(
	controller.UserControllerInterface,
) {
	// init dependencies
	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	controller := controller.NewUserControllerInterface(service)
	
	return controller
}