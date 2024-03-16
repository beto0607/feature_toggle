package main

import (
	"log"
	"toggler/configs"
	"toggler/routes"

	"net/http"
)

func main() {
	prepareDB()
	prepareServer()
}

func prepareServer() {
	serverPort := configs.EnvPort()
	hostname := configs.EnvHostname()

	routes.DoApiRouting()

	serverAddress := hostname + ":" + serverPort

	server := &http.Server{
		Addr:         serverAddress,
		ReadTimeout:  configs.DefaultReadTimeout(),
		WriteTimeout: configs.DefaultWriteTimeout(),
	}

	log.Println("Starting server on " + server.Addr)
	log.Fatal(server.ListenAndServe())
}

func prepareDB() {
	if configs.ShouldLoadDB() {
		configs.ConnectToDB()
	}
}
