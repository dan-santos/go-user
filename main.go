package main

import (
	"log/slog"

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
	
	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup)

	if err := router.Run(":8080"); err != nil {
		slog.Error("Error on statup of API")
	}
}