package main

import (
	"context"
	"log/slog"

	"github.com/dan-santos/go-user/src/configs/database/mongodb"
	"github.com/dan-santos/go-user/src/configs/logger"
	"github.com/dan-santos/go-user/src/controller/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	logger.Info("Initiating User API statup")
	err := godotenv.Load()
	if err != nil {
		slog.Error("Error loading .env file")
	}

	database, err := mongodb.NewMongoDBConnection(context.Background())
	if err != nil {
		logger.Error("Error trying to connect to database", err)
		return
	}
	controller := initDependencies(database)
	
	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, controller)

	if err := router.Run(":8080"); err != nil {
		slog.Error("Error on statup of API")
	}
}