package main

import (
	"github.com/arturbaccarin/huncoding-go-videos/hexagonal-api/adapter/input/routes"
	"github.com/arturbaccarin/huncoding-go-videos/hexagonal-api/configuration/logger"
	"github.com/gin-gonic/gin"
)

func main() {
	logger.Info("about to start application")
	router := gin.Default()

	routes.InitRoutes(router)

	err := router.Run(":8080")
	if err != nil {
		logger.Error("error trying to start server", err)
		return
	}
}
