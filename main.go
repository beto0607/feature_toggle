package main

import (
	"log"
	"toggler/configs"
	"toggler/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	engine := gin.Default()
	apiGroup := engine.Group("/api")

	// Routes
	routes.UsersRoute(apiGroup)
	routes.FeaturesRoutes(apiGroup)
	routes.AccountsRoutes(apiGroup)

	serverPort := configs.EnvPort()
	err := engine.Run("localhost:" + serverPort)
	log.Println("Running on localhost:" + serverPort)

	if err != nil {
		log.Fatal(err.Error())
	}
}
