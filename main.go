package main

import (
	"log"
	"toggler/configs"
	"toggler/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	// run DB
	configs.ConnectDB()

	engine := gin.Default()
	apiGroup := engine.Group("/api")

	// Routes
	routes.UserRoute(apiGroup)
	routes.FeatureRoutes(apiGroup)

	serverPort := configs.EnvPort()
	err := engine.Run("localhost:" + serverPort)
	if err != nil {
		log.Fatal(err.Error())
	}
}
